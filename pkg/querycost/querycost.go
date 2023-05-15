package querycost

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"os"
)

func NewNewQuery(_profile string, _start string, _end string, _granularity string, _groupby string, _filter string, _metrics string, _output string, _filterType string, _formatter string) {
	var output Output
	switch _output {
	case "":
		output = StandardOutput{}
	default:
		output = NewCSVOutput(_output)
	}
	NewQuery(_profile, _start, _end, _granularity, _groupby, _filter, _metrics, output, _filterType, _formatter)
}
func NewQuery(_profile string, _start string, _end string, _granularity string, _groupby string, _filter string, _metrics string, output Output, _filterType string, _formatter string) {

	profiles := arrayFromParameter(_profile)
	filter := arrayFromParameter(_filter)
	metrics := arrayFromParameter(_metrics)
	startDate, endDate := startDateEndDate(_start, _end)
	filterType := _filterType

	if len(profiles) == 0 {
		envProfile := os.Getenv("AWS_PROFILE")
		if envProfile == "" {
			panic("No profile found. Use either -a or set the AWS_PROFILE environment variable.")
		}
		profiles = []string{envProfile}
	}
	var formatter Formatter
	switch _formatter {
	case "alias":
		formatter = ReplaceAccountAliasFormatter{}
	case "custom":
		formatter = CustomFormatter{}
	default:
		formatter = SimpleFormatter{}
	}

	for _, __profile := range profiles {
		query := Query{Profile: __profile, StartDate: startDate, EndDate: endDate, Granularity: types.Granularity(_granularity), Dimension: _groupby, Filter: filter, Metrics: metrics, FilterType: filterType, Formatter: formatter}
		QueryCostWithQuery(query, output)
	}
}

func QueryCostWithQuery(query Query, out Output) {
	awsInput := prepareAWSInput(query)
	awsOutput := executeQueryWithAWSInput(query, awsInput)
	resultsCosts := Result{awsOutput, query}
	formatted := query.Formatter.Format(resultsCosts)
	out.DisplayResult(formatted)
}

func executeQueryWithAWSInput(query Query, input *costexplorer.GetCostAndUsageInput) *costexplorer.GetCostAndUsageOutput {
	cfg, err := GetConfigForProfile(query.Profile)

	svc := costexplorer.NewFromConfig(cfg)

	output, err := svc.GetCostAndUsage(context.Background(), input)
	if err != nil {
		panic(fmt.Sprintf("ERROR [%s] while running query with input: +%v", err, input))
	}

	var token = output.NextPageToken

	for token != nil {
		// fmt.Printf("Next Token: %v\n", *token)
		input.NextPageToken = token
		outputBis, _ := svc.GetCostAndUsage(context.Background(), input)
		output.ResultsByTime = append(output.ResultsByTime, outputBis.ResultsByTime...)
		token = outputBis.NextPageToken
	}
	return output
}

func prepareAWSInput(query Query) *costexplorer.GetCostAndUsageInput {
	input := &costexplorer.GetCostAndUsageInput{
		Filter:      getFilterFromQuery(query),
		Granularity: query.Granularity,
		TimePeriod: &types.DateInterval{
			Start: aws.String(query.StartDate),
			End:   aws.String(query.EndDate),
		},
		Metrics: query.Metrics,
		GroupBy: []types.GroupDefinition{
			{
				Type: types.GroupDefinitionTypeDimension,
				Key:  aws.String(query.Dimension),
			},
		},
	}
	return input
}

func getFilterFromQuery(query Query) *types.Expression {
	filter := query.Filter
	filterType := query.FilterType

	var _filter *types.Expression
	if len(filter) != 0 {
		_filter = &types.Expression{
			//CostCategories: &types.CostCategoryValues{
			//
			//},
			Dimensions: &types.DimensionValues{
				Key:    types.Dimension(filterType),
				Values: filter,
			},
		}
	}
	return _filter
}

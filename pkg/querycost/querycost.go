package querycost

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"os"
)

type Query struct {
	Profile     string
	StartDate   string
	EndDate     string
	Granularity types.Granularity
	Dimension   string
	Filter      []string
	Metrics     []string
	Output      Output
	FilterType  string
}

type Result struct {
	Output *costexplorer.GetCostAndUsageOutput
	Query  Query
}

func QueryCost(_profile string, _start string, _end string, _granularity string, _groupby string, _filter string, _metrics string, _output string, _filterType string) {

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

	for _, __profile := range profiles {
		query := Query{Profile: __profile, StartDate: startDate, EndDate: endDate, Granularity: types.Granularity(_granularity), Dimension: _groupby, Filter: filter, Metrics: metrics, FilterType: filterType}
		var output Output = StandardOutput{}
		if _output != "" {
			output = NewCSVOutput(_output)
		}
		QueryCostWithQuery(query, output)
	}

}

func QueryCostWithQuery(query Query, out Output) {
	var profile = query.Profile

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))

	if err != nil {
		panic("Cannot load aws profile.")
	}

	input := prepareAWSInput(query)
	output := executeQueryWithAWSInput(cfg, input)
	resultsCosts := Result{output, query}
	out.DisplayResult(SimpleFormatter{}.Format(resultsCosts))
}

func executeQueryWithAWSInput(cfg aws.Config, input *costexplorer.GetCostAndUsageInput) *costexplorer.GetCostAndUsageOutput {
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

	start := query.StartDate
	end := query.EndDate
	granularity := query.Granularity
	filter := query.Filter
	filterType := query.FilterType
	metrics := query.Metrics
	groupby := query.Dimension

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

	input := &costexplorer.GetCostAndUsageInput{
		Filter:      _filter,
		Granularity: granularity,
		TimePeriod: &types.DateInterval{
			Start: aws.String(start),
			End:   aws.String(end),
		},
		Metrics: metrics,
		GroupBy: []types.GroupDefinition{
			{
				Type: types.GroupDefinitionTypeDimension,
				Key:  aws.String(groupby),
			},
		},
	}
	return input
}

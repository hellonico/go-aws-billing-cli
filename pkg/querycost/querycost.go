package querycost

import (
	"context"
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
}

type Result struct {
	Output *costexplorer.GetCostAndUsageOutput
	Query  Query
}

func QueryCost(_profile string, _start string, _end string, _granularity string, _groupby string, _filter string, _metrics string, _output string) {

	profiles := arrayFromParameter(_profile)
	filter := arrayFromParameter(_filter)
	metrics := arrayFromParameter(_metrics)
	startDate, endDate := startDateEndDate(_start, _end)

	for _, __profile := range profiles {
		query := Query{Profile: __profile, StartDate: startDate, EndDate: endDate, Granularity: types.Granularity(_granularity), Dimension: _groupby, Filter: filter, Metrics: metrics}
		var output Output = StandardOutput{}
		if _output != "" {
			output = NewCSVOutput(_output)
		}
		QueryCostWithQuery(query, output)
	}

}

func QueryCostWithQuery(query Query, out Output) {
	var profile = query.Profile
	if query.Profile == "" {
		profile = os.Getenv("AWS_PROFILE")
		if profile == "" {
			panic("No profile found. Use either -a or set the AWS_PROFILE environment variable.")
		}
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))

	if err != nil {
		panic("Cannot load aws profile.")
	}

	input := prepareAWSInput(query.Filter, query.StartDate, query.EndDate, query.Granularity, query.Metrics, query.Dimension)
	output := executeQueryWithAWSInput(cfg, input)
	resultsCosts := Result{output, query}
	out.DisplayResult(SimpleFormatter{}.Format(resultsCosts))
}

func executeQueryWithAWSInput(cfg aws.Config, input *costexplorer.GetCostAndUsageInput) *costexplorer.GetCostAndUsageOutput {
	svc := costexplorer.NewFromConfig(cfg)
	output, _ := svc.GetCostAndUsage(context.Background(), input)
	return output
}

func prepareAWSInput(filter []string, start string, end string, granularity types.Granularity, metrics []string, groupby string) *costexplorer.GetCostAndUsageInput {
	var _filter *types.Expression
	if len(filter) != 0 {
		_filter = &types.Expression{
			//CostCategories: &types.CostCategoryValues{
			//
			//},
			Dimensions: &types.DimensionValues{
				Key:    "SERVICE",
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

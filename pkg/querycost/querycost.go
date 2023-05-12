package querycost

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"os"
	"strings"
)

func QueryCost(profile string, start string, end string, granularity types.Granularity, groupby string, filter []string, metrics []string) {
	if profile == "" {
		profile = os.Getenv("AWS_PROFILE")
		if profile == "" {
			panic("No profile found. Use either -a or set the AWS_PROFILE environment variable.")
		}
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))

	if err != nil {
		panic("Cannot load aws profile.")
	}

	input := prepareInput(filter, start, end, granularity, metrics, groupby)
	output := executeQuery(cfg, input)
	resultsCosts := handleResults(groupby, metrics, output)
	displayResults(resultsCosts)
}

func displayResults(results [][]string) {
	for _, row := range results {
		fmt.Println(strings.Join(row, ","))
	}
}

func executeQuery(cfg aws.Config, input *costexplorer.GetCostAndUsageInput) *costexplorer.GetCostAndUsageOutput {
	svc := costexplorer.NewFromConfig(cfg)
	output, _ := svc.GetCostAndUsage(context.Background(), input)
	return output
}

func prepareInput(filter []string, start string, end string, granularity types.Granularity, metrics []string, groupby string) *costexplorer.GetCostAndUsageInput {
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
		Granularity: "MONTHLY",
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

func handleResults(groupby string, metrics []string, output *costexplorer.GetCostAndUsageOutput) [][]string {
	var resultsCosts [][]string

	var headers = []string{"startDate", "endDate", groupby}
	for _, metric := range metrics {
		headers = append(headers, metric)
	}
	resultsCosts = append(resultsCosts, headers)

	for _, results := range output.ResultsByTime {
		startDate := *results.TimePeriod.Start
		endDate := *results.TimePeriod.End
		for _, groups := range results.Groups {
			var info = []string{startDate, endDate, groups.Keys[0]}
			for _, metrics := range groups.Metrics {
				info = append(info, *metrics.Amount)
			}
			resultsCosts = append(resultsCosts, info)
		}
	}
	return resultsCosts
}

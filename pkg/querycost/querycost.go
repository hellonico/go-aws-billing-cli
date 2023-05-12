package querycost

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"strings"
)

func displayResults(results [][]string) {
	for _, row := range results {
		fmt.Println(strings.Join(row, ","))
	}
}

func QueryCost(profile string, start string, end string, groupby string, filter []string, metrics []string) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))

	if err != nil {
		panic(err)
	}

	svc := costexplorer.NewFromConfig(cfg)

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
		Granularity: types.GranularityMonthly,
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

	output, _ := svc.GetCostAndUsage(context.Background(), input)
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

	displayResults(resultsCosts)
}

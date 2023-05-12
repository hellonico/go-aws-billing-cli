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

	// Create a CostExplorer client using the loaded AWS credentials and region
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

	for _, results := range output.ResultsByTime {
		startDate := *results.TimePeriod.Start
		for _, groups := range results.Groups {
			for _, metrics := range groups.Metrics {
				info := []string{startDate, groups.Keys[0], *metrics.Amount}
				resultsCosts = append(resultsCosts, info)
			}
		}
	}

	displayResults(resultsCosts)
}

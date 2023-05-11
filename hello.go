package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	types "github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"os"
	"strings"
	"time"
)

/*
*
DOCS:
https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/costexplorer
https://stackoverflow.com/questions/74623677/adding-filters-to-aws-cost-sdk-in-go
https://medium.com/@ramonesparza1234/go-aws-costexplorer-a-simple-aws-cost-analyzer-831baed2e125
https://docs.aws.amazon.com/sdk-for-go/api/service/costexplorer/#CostExplorer.GetCostForecast
https://docs.aws.amazon.com/ja_jp/aws-cost-management/latest/APIReference/API_GetCostAndUsage.html
*/
func main() {

	var profile = flag.String("a", "default", "profile name")
	var month = flag.Int("m", 0, "how many months back in time")
	var dimension = flag.String("g", "LINKED_ACCOUNT", "dimension, one of: AZ, INSTANCE_TYPE, LEGAL_ENTITY_NAME, INVOICING_ENTITY, LINKED_ACCOUNT, OPERATION, PLATFORM, PURCHASE_TYPE, SERVICE, TENANCY, RECORD_TYPE, and USAGE_TYPE")
	var help = flag.Bool("help", false, "print usage")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	queryCost(*profile, *month, *dimension)
}

func queryCost(profile string, month int, groupby string) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))

	if err != nil {
		panic(err)
	}

	// Create a CostExplorer client using the loaded AWS credentials and region
	svc := costexplorer.NewFromConfig(cfg)

	now := time.Now()
	var startTime time.Time
	if month == 0 {
		startTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	} else {
		startTime = now.AddDate(0, month, 0)
	}

	var endTime = now

	input := &costexplorer.GetCostAndUsageInput{
		//Filter: &types.Expression{
		//	CostCategories: &types.CostCategoryValues{
		//		Key:    aws.String("SERVICE"),
		//		Values: []string{"Amazon Route 53"},
		//	},
		//},
		Granularity: types.GranularityMonthly,
		TimePeriod: &types.DateInterval{
			Start: aws.String(startTime.Format("2006-01-02")),
			End:   aws.String(endTime.Format("2006-01-02")),
		},
		Metrics: []string{"UnblendedCost"},
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

func displayResults(results [][]string) {
	for _, row := range results {
		fmt.Println(strings.Join(row, ","))
	}
}

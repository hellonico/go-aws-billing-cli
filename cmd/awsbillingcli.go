package main

import (
	"flag"
	"github.com/hellonico/go-aws-billing-cli/pkg/querycost"
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

	var profile = flag.String("a", "default", "aws profile name")
	var month = flag.Int("m", 0, "how many months back in time")
	var start = flag.String("start", "", "start date. if this is set, month is ignored")
	var end = flag.String("end", "", "end date")
	var dimension = flag.String("g", "LINKED_ACCOUNT", "group by dimension, one of: AZ, INSTANCE_TYPE, LEGAL_ENTITY_NAME, INVOICING_ENTITY, LINKED_ACCOUNT, OPERATION, PLATFORM, PURCHASE_TYPE, SERVICE, TENANCY, RECORD_TYPE, and USAGE_TYPE")
	var _filter = flag.String("f", "", "Filter by services. Use , to separate; one of Amazon Route 53, AmazonCloudWatch, Amazon Route 53...")
	var _metrics = flag.String("metrics", "UnblendedCost", "Metrics. Default metric is: UnblendedCost")

	var help = flag.Bool("help", false, "print usage")
	flag.Parse()
	filter := arrayFromParameter(_filter)
	metrics := arrayFromParameter(_metrics)

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	startDate, endDate := startDateEndDate(*month, *start, *end)

	querycost.QueryCost(*profile, startDate, endDate, *dimension, filter, metrics)
}

func startDateEndDate(month int, start string, end string) (string, string) {
	now := time.Now()
	var startTime time.Time
	var endTime = now

	if start != "" {
		startTime, _ = time.Parse("2006-01-02", start)
		if end != "" {
			endTime, _ = time.Parse("2006-01-02", end)
		}
	} else {
		if month == 0 {
			startTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		} else {
			if month > 0 {
				startTime = now.AddDate(0, -1*month, 0)
			} else {
				startTime = now.AddDate(0, month, 0)
			}

		}
	}

	return startTime.Format("2006-01-02"), endTime.Format("2006-01-02")

}

func arrayFromParameter(_filter *string) []string {
	var filter []string
	if *_filter != "" {
		filter = strings.Split(*_filter, ",")
	}
	return filter
}

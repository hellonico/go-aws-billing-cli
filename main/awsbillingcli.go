package main

import (
	"flag"
	"github.com/hellonico/go-aws-billing-cli/pkg/querycost"
	"os"
	"strings"
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
	var _filter = flag.String("f", "", "use , to separate; one of Amazon Route 53, AmazonCloudWatch, Amazon Route 53")

	var help = flag.Bool("help", false, "print usage")
	flag.Parse()
	var filter = []string{}
	if *_filter != "" {
		filter = strings.Split(*_filter, ",")
	}

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	querycost.QueryCost(*profile, *month, *dimension, filter)
}

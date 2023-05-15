package main

import (
	"flag"
	"github.com/hellonico/go-aws-billing-cli/pkg/querycost"
	"os"
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

	var profile = flag.String("a", "", "aws profile name. can specify multiple")
	//var script = flag.String("s", "", "repeat query with lines read from a text file")
	var output = flag.String("o", "", "output. empty for console output")
	var granularity = flag.String("gr", "MONTHLY", "granularity. one of:\nMONTHLY\nDAILY\nYEARLY\n")
	var filterType = flag.String("ft", "SERVICE", "Dimension to filter by. (Ex: SERVICE, LINKED_ACCOUNT etc...")
	var start = flag.String("start", "", "start date. if this is set, month is ignored")
	var end = flag.String("end", "", "end date")
	var dimension = flag.String("g", "LINKED_ACCOUNT", "group by dimension, one of:\nAZ\nINSTANCE_TYPE\nLEGAL_ENTITY_NAME\nINVOICING_ENTITY\nLINKED_ACCOUNT\nOPERATION\nPLATFORM\nPURCHASE_TYPE\nSERVICE\nTENANCY\nRECORD_TYPE\nUSAGE_TYPE\n")
	var _filter = flag.String("f", "", "Filter by services. Use , to separate; one or more of (non-exhaustive):\nAWS CloudTrail\nAWS Config\nAWS Cost Explorer\nAWS Directory Service\nAWS Glue\nAWS Key Management Service\nAWS Lambda\nAWS Step Functions\nAWS Support (Developer)\nAmazon Chime\nAmazon Chime Dialin\nAmazon EC2 Container Registry (ECR)\nEC2 - Other\nAmazon Elastic Compute Cloud - Compute\nAmazon Elastic Load Balancing\nAmazon GuardDuty\nAmazon MQ\nAmazon Registrar\nAmazon Relational Database Service\nAmazon Route 53\nAmazon Simple Notification Service\nAmazon Simple Queue Service\nAmazon Simple Storage Service\nAmazon Virtual Private Cloud\nAmazonCloudWatch")
	var _metrics = flag.String("m", "UnblendedCost", "Metrics. One or more of:\nAmortizedCost\nBlendedCost\nNetAmortizedCost\nNetUnblendedCost\nNormalizedUsageAmount\nUnblendedCost\nUsageQuantity\n")
	var _formatter = flag.String("fm", "", "Formatting")

	//var la = flag.Bool("la", false, "listaccounts")

	var help = flag.Bool("help", false, "print usage")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	querycost.NewNewQuery(*profile, *start, *end, *granularity, *dimension, *_filter, *_metrics, *output, *filterType, *_formatter)

}

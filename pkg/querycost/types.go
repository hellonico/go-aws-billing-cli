package querycost

import (
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
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
	Formatter   Formatter
}

type Result struct {
	Output *costexplorer.GetCostAndUsageOutput
	Query  Query
}

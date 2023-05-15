package querycost

type FormattedResult struct {
	Value [][]string
	Query Query
}

type Formatter interface {
	Format(res Result) FormattedResult
}
type SimpleFormatter struct{}

func (s SimpleFormatter) Format(res Result) FormattedResult {
	var resultsCosts [][]string

	var headers = []string{"startDate", "endDate", res.Query.Dimension}
	for _, metric := range res.Query.Metrics {
		headers = append(headers, metric)
	}
	resultsCosts = append(resultsCosts, headers)

	for _, results := range res.Output.ResultsByTime {
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
	return FormattedResult{resultsCosts, res.Query}
}

/*
*
 */
type ReplaceAccountAliasFormatter struct{}

func (s ReplaceAccountAliasFormatter) Format(res Result) FormattedResult {
	var resultsCosts [][]string

	var headers = []string{"startDate", "endDate", res.Query.Dimension}
	for _, metric := range res.Query.Metrics {
		headers = append(headers, metric)
	}
	resultsCosts = append(resultsCosts, headers)

	if res.Query.Dimension == "LINKED_ACCOUNT" {
		accounts, _ := FetchAccountsMap(res.Query.Profile)

		for _, results := range res.Output.ResultsByTime {
			startDate := *results.TimePeriod.Start
			endDate := *results.TimePeriod.End
			for _, groups := range results.Groups {
				var info = []string{startDate, endDate, accounts[groups.Keys[0]]}
				for _, metrics := range groups.Metrics {
					info = append(info, *metrics.Amount)
				}
				resultsCosts = append(resultsCosts, info)
			}
		}
	} else {

		for _, results := range res.Output.ResultsByTime {
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

	}

	return FormattedResult{resultsCosts, res.Query}
}

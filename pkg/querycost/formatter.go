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

	if res.Query.Dimension == "LINKED_ACCOUNT" {
		var resultsCosts [][]string

		var headers = []string{"startDate", "endDate", res.Query.Dimension}
		for _, metric := range res.Query.Metrics {
			headers = append(headers, metric)
		}
		resultsCosts = append(resultsCosts, headers)

		accounts, _ := FetchAccountsMap(res.Query.Profile)

		for _, results := range res.Output.ResultsByTime {
			startDate := *results.TimePeriod.Start
			endDate := *results.TimePeriod.End
			for _, groups := range results.Groups {
				var info = []string{startDate, endDate}

				aa, ok := accounts[groups.Keys[0]]
				if !ok {
					info = append(info, groups.Keys[0])
				} else {
					info = append(info, aa)
				}

				for _, metrics := range groups.Metrics {
					info = append(info, *metrics.Amount)
				}
				resultsCosts = append(resultsCosts, info)
			}
		}

		return FormattedResult{resultsCosts, res.Query}
	} else {
		return SimpleFormatter{}.Format(res)
	}

}

type CustomFormatter struct{}

func (s CustomFormatter) Format(res Result) FormattedResult {

	if res.Query.Dimension == "LINKED_ACCOUNT" {
		var resultsCosts [][]string

		//var headers = []string{"連結アカウント,連結アカウント名,連結アカウント の合計", res.Query.StartDate}
		//resultsCosts = append(resultsCosts, headers)

		accounts, _ := FetchAccountsMap(res.Query.Profile)

		for _, results := range res.Output.ResultsByTime {
			startDate := *results.TimePeriod.Start
			for _, groups := range results.Groups {
				accountId := groups.Keys[0]
				var info = []string{accountId, startDate}

				aa, ok := accounts[accountId]
				if !ok {
					info = append(info, "")
				} else {
					info = append(info, aa)
				}
				// estimate
				info = append(info, "")

				for _, metrics := range groups.Metrics {
					info = append(info, *metrics.Amount)
				}

				resultsCosts = append(resultsCosts, info)
			}
		}

		return FormattedResult{resultsCosts, res.Query}
	} else {
		return SimpleFormatter{}.Format(res)
	}

}

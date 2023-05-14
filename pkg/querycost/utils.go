package querycost

import (
	"strconv"
	"strings"
	"time"
)

func parseTime(_time string) string {
	month, err := strconv.Atoi(_time)
	now := time.Now()
	var res time.Time
	if err != nil {
		// not an integer so parse as a date
		if _time == "" {
			// empty date is now
			res = now
		} else {
			res, _ = time.Parse("2006-01-02", _time)
		}
	} else {
		if month == 0 {
			res = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		} else {
			if month > 0 {
				res = now.AddDate(0, -1*month, 0)
			} else {
				res = now.AddDate(0, month, 0)
			}
		}
	}
	return res.Format("2006-01-02")
}
func startDateEndDate(start string, end string) (string, string) {
	if start == "" && end == "" {
		start = "1"
	}
	return parseTime(start), parseTime(end)
}

func arrayFromParameter(_filter string) []string {
	var filter []string
	if _filter != "" {
		filter = strings.Split(_filter, ",")
	}
	return filter
}

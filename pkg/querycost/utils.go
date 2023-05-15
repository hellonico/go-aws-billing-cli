package querycost

import (
	"strconv"
	"strings"
	"time"
)

func parseTime(_time string) string {
	now := time.Now()

	month, err := strconv.Atoi(_time)
	var res time.Time
	if err != nil {
		// not an integer so parse as a date
		switch _time {
		case "":
			// empty date is now
			res = now
		case "bom":
			res = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		case "boy":
			res = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
		case "eoy":
			res = time.Date(now.Year(), 12, 31, 23, 59, 0, 0, now.Location())
		case "bolm":
			res = time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, now.Location())
		case "eolm":
			// .AddDate(0, 0, -1)
			res = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		default:
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

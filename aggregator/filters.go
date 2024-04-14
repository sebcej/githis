package aggregator

import (
	"time"

	"github.com/sebcej/githis/utils"
	"github.com/spf13/cobra"
)

func filter(filters Filters, log Log) bool {
	cur := time.Now()
	logDate, _ := utils.ParseLogDate(log.Date)

	if filters.Offset != 0 {
		cur = cur.AddDate(0, 0, filters.Offset)

		y1, m1, d1 := logDate.Date()
		y2, m2, d2 := cur.Date()

		if y1 != y2 || m1 != m2 || d1 != d2 {
			return false
		}
	}

	if len(filters.Authors) > 0 {
		present := false

		for _, author := range filters.Authors {
			if author == log.Author.Name || author == log.Author.Email {
				present = true
				break
			}
		}

		if !present {
			return false
		}
	}

	if filters.FromDay != "" {
		date, _ := utils.ParseLogDate(log.Date)

		from, err := utils.ParseDate(filters.FromDay)
		cobra.CheckErr(err)
		to := time.Now()

		if filters.ToDay != "" {
			to, err = utils.ParseDate(filters.ToDay)
			cobra.CheckErr(err)
		}

		if !date.After(from) || !date.Before(to) {
			return false
		}
	}

	return true
}

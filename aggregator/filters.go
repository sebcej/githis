package aggregator

import (
	"time"

	"github.com/sebcej/githis/utils"
)

func filter(filters Filters, log Log) bool {
	cur := time.Now()
	logDate, _ := utils.ParseDate(log.Date)

	if filters.Offset != 0 {
		cur = cur.AddDate(0, 0, filters.Offset)

		y1, m1, d1 := logDate.Date()
		y2, m2, d2 := cur.Date()

		return y1 == y2 && m1 == m2 && d1 == d2
	}

	return true
}

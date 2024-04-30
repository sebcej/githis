package aggregator

import (
	"strconv"
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

	if filters.Day != "" {
		dayFilter := filters.Day

		// Allow to insert only partial dates to day filter
		// Example: 01 -> 2024-01-01 or 01-01 -> 2024-01-01
		if matchDatePartialDay.MatchString(dayFilter) {
			day, _ := strconv.Atoi(dayFilter)
			date := time.Now()

			date = time.Date(date.Year(), date.Month(), day, 0, 0, 0, 0, time.Local)
			dayFilter = utils.FormatDate(date)
		} else if matchDatePartialDayMonth.MatchString(dayFilter) {
			dateArr := matchDatePartialDayMonth.FindStringSubmatch(dayFilter)
			date := time.Now()

			day, _ := strconv.Atoi(dateArr[2])
			month, _ := strconv.Atoi(dateArr[1])

			date = time.Date(date.Year(), time.Month(month), day, 0, 0, 0, 0, time.Local)
			dayFilter = utils.FormatDate(date)
		}

		filters.From = dayFilter
		filters.To = dayFilter
	}

	if filters.From != "" {
		date, _ := utils.ParseLogDate(log.Date)

		from, err := utils.ParseDate(filters.From)
		cobra.CheckErr(err)
		to := time.Now()

		if filters.To != "" {
			to, err = utils.ParseDate(filters.To)
			cobra.CheckErr(err)

			to = to.AddDate(0, 0, 1)
		}

		if !date.After(from) || !date.Before(to) {
			return false
		}
	}

	return true
}

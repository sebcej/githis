package utils

import "time"

func ParseLogDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", date)
}

func ParseDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

func FormatDate(date time.Time) string {
	return date.Format("2006-01-02")
}

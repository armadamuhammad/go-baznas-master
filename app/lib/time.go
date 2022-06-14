package lib

import (
	"time"
)

// CurrentTime func
func CurrentTime(format string) string {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	return time.Now().Format(format)
}

// TimeNow func
func TimeNow() *time.Time {
	t := time.Now()
	return &t
}

// RangeDate func
func RangeDate(date1, date2, types string) float64 {
	// date_1 > date_2
	// format_date : 2006-01-02 15:04:05
	var result float64

	format := "2006-01-02 15:04:05"
	dateOne, _ := time.Parse(format, date1)
	dateTwo, _ := time.Parse(format, date2)

	diff := dateOne.Sub(dateTwo)
	// number of Hours
	if types == "hours" {
		result = diff.Hours()

		// number of Nanoseconds
	} else if types == "nanoseconds" {
		result = float64(diff.Nanoseconds())

		// number of Minutes
	} else if types == "minutes" {
		result = diff.Minutes()

		// number of Seconds
	} else if types == "seconds" {
		result = diff.Seconds()

		// number of Days
	} else if types == "days" {
		result = float64(diff.Hours() / 24)
	}
	return result
}

// DateTimeAhead func
func DateTimeAhead(fromDate, format string, years, month, days int) string {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	date, _ := time.Parse(format, fromDate)
	t2 := date.AddDate(years, month, days)
	CurrentTimeAhead := t2.Format(format)
	return CurrentTimeAhead
}

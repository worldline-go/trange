package trange

import (
	"time"
)

const (
	DateLayout = "2006-01-02"
)

func Day(dateStr string, loc *time.Location) (time.Time, time.Time, error) {
	start, err := time.ParseInLocation(DateLayout, dateStr, loc)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	end := time.Date(start.Year(), start.Month(), start.Day(), 23, 59, 59, 999_999_999, loc)

	return start, end, nil
}

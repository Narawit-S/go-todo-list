package utils

import (
	"time"
)

func DateFromNow(days int) time.Time {
	now := time.Now().UTC()

	date_from_now := now.AddDate(0, 0, days)

	return date_from_now
}

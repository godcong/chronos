package utils

import (
	"time"
)

const Day = 24

// BetweenTime returns the duration between two times.
func BetweenTime(d1, d2 time.Time) time.Duration {
	return d1.Sub(d2)
}

// BetweenDay returns the number of days between two times
func BetweenDay(current time.Time, start time.Time) int {
	return int(current.Sub(start).Hours() / 24)
}

package utils

import (
	"time"
)

const Day = 24

// BetweenTime returns the time.Duration between two times
// @param time.Time
// @param time.Time
// @return time.Duration
func BetweenTime(d1, d2 time.Time) time.Duration {
	return d1.Sub(d2)
}

// BetweenDay returns the number of days between two times
func BetweenDay(current time.Time, start time.Time) int {
	return int(current.Sub(start).Hours() / 24)
}

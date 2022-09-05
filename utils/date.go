package utils

import (
	"time"
)

const Day = 24

func DateDiff(d1, d2 time.Time) time.Duration {
	return d1.Sub(d2)
}

func DateDiffDay(d1, d2 time.Time) int {
	return int(d1.Sub(d2).Hours() / Day)
}

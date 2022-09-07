package utils

import (
	"math"
	"time"
)

const Day = 24

func DateDiff(d1, d2 time.Time) time.Duration {
	return d1.Sub(d2)
}

//func DateDiffDay(d1, d2 time.Time) int {
//	return int(d1.Sub(d2).Hours() / Day)
//}

//BetweenDay 计算两个时间差的天数
func BetweenDay(d time.Time, s time.Time) int {
	return int(math.Abs(d.Sub(s).Hours() / 24))
}

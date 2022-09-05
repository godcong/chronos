package chronos

import (
	"time"
)

// Calendar ...
type Calendar interface {
	Lunar() Lunar
	Solar() Solar
	FormatTime() string
	Time() time.Time
	Date() CalendarDate
}

// Solar ...
type Solar interface {
	Minute() int
	Second() int
	YearDay() int
	IsLeapYear() bool
	Year() int
	Month() time.Month
	Day() int
	Hour() int
}

type Lunar interface {
	Year() int
	Month() int
	Day() int
	Hour() int
	LeapMonth() int
	IsLeapMonth() bool
}

type ChineseSupport interface {
	Chinese() string
}

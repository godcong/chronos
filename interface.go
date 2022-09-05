package chronos

import (
	"time"
)

// Calendar returns a calendar
type Calendar interface {
	Lunar() Lunar
	Solar() Solar
	FormatTime() string
	Time() time.Time
	Date() CalendarDate
}

// Solar returns the solar time
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

// Lunar returns the lunar time
type Lunar interface {
	Year() int
	Month() int
	Day() int
	Hour() int
	LeapMonth() int
	IsLeapMonth() bool
}

// ChineseSupport implements the interface, if support chinese language output
type ChineseSupport interface {
	Chinese() string
}

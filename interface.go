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
	ViewData() View
}

// Solar ...
type Solar interface {
	Year() int
	Month() int
	Day() int
	Hour() int
	IsLeapYear() bool
}

type Lunar interface {
	Year() int
	Month() int
	Day() int
	Hour() int
	IsLeapYear() bool
	IsLeapMonth() bool
}

//go:generate go-enum --marshal -f constellation.go
//go:generate go-enum --marshal -f gan_zhi.go
//go:generate go-enum --marshal -f solar_term.go
//go:generate go-enum --marshal -f zodiac.go
package chronos

import (
	"fmt"
	"time"
)

const minYear = 1900
const maxYear = 3000

const (
	// DateFormatYMD ...
	DateFormatYMD = "2006/01/02"
	// DateFormatYMDHMS ...
	DateFormatYMDHMS = "2006/01/02 15:04:05"
)

var (
	loc, _ = time.LoadLocation("Asia/Shanghai")
	// startTimeUnix is 1900/01/01 00:00:00
	// startTimeUnix = uint64(0xFFFFFFFF7C558180)
	// startTime is 1900/01/01 00:00:00
	startTime = TimeFromYmd(1900, 1, 1)
	// lunarStartTimeUnix = uint64(0xFFFFFFFF7C7C9E00)
	// lunarStartTime = TimeFromYmd(1900, 1, 31)
)

type calendarTime struct {
	time time.Time

	solar Solar
	lunar Lunar
}

func (c *calendarTime) FormatTime() string {
	return c.time.Format(DateFormatYMDHMS)
}

func (c *calendarTime) Time() time.Time {
	return c.time
}

func (c *calendarTime) LocalTime() time.Time {
	return c.time
}

// Lunar ...
func (c *calendarTime) Lunar() Lunar {
	return c.lunar
}

// Solar ...
func (c *calendarTime) Solar() Solar {
	return c.solar
}

func (c *calendarTime) initialize() *calendarTime {
	c.solar = ParseSolarByTime(c.time)
	c.lunar = ParseLunarTime(c.time)
	return c
}

// NewSolarCalendar can input three type of time to create the calendarTime
// "2006/01/02 03:04" format string
// time.Time value
// or nil to create a new time.Now() value
func NewSolarCalendar(v ...any) Calendar {
	var c *calendarTime
	if len(v) == 0 {
		c = &calendarTime{
			time: localTime(),
		}
		return c.initialize()
	}

	first := parseFirstArg(v)
	args := parseArgs(v)
	switch vv := first.(type) {
	case int:
		c = parseIntDate(vv, args...)
	case string:
		c = parseStringDate(vv, args...)
	case time.Time:
		c = parseTime(vv.In(loc))
	default:
		c = parseTime(localTime())
	}
	return c.initialize()
}

// ParseSolarString returns Calendar parse from string(value,format?)
// @param string
// @param ...string
// @return Calendar
func ParseSolarString(s string, format ...string) Calendar {
	return parseStringDateFormat(s, format...).initialize()
}

// ParseSolarDate returns Calendar parse from date(year, month, day, hour, minute, second)
// @param int
// @param int
// @param int
// @param int
// @param int
// @param int
// @return Calendar
func ParseSolarDate(year, month, day, hour, minute, second int) Calendar {
	date := time.Date(year, time.Month(month), day, hour, minute, second, 0, loc)
	return parseTime(date).initialize()
}

// ParseSolarNow returns Calendar parse from solar time now(time.Now())
// @return Calendar
func ParseSolarNow() Calendar {
	return parseTime(localTime()).initialize()
}

// ParseSolarTime returns Calendar parse from solar time
// @param time.Time
// @return Calendar
func ParseSolarTime(t time.Time) Calendar {
	return parseTime(t.In(loc)).initialize()
}

func TimeFromY(y int) time.Time {
	return time.Date(y, 1, 1, 0, 0, 0, 0, loc)
}

func TimeFromYm(y int, m time.Month) time.Time {
	return time.Date(y, m, 1, 0, 0, 0, 0, loc)
}

func TimeFromYmd(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, loc)
}

func TimeFromYmdHms(Y int, M time.Month, D int, h, m, s int) time.Time {
	return time.Date(Y, M, D, h, m, s, 0, loc)
}

func checkYearSupport(year int) error {
	if year < minYear || year > maxYear {
		return fmt.Errorf("[chronos] year %d not supported", year)
	}
	return nil
}

var _ Calendar = (*calendarTime)(nil)

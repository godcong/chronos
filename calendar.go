//go:generate go-enum --marshal -f constellation.go
//go:generate go-enum --marshal -f gan_zhi.go
//go:generate go-enum --marshal -f solar_term.go
//go:generate go-enum --marshal -f zodiac.go
package chronos

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	// DateFormatYMD ...
	DateFormatYMD = "2006/01/02"
	// DateFormatYMDHMS ...
	DateFormatYMDHMS = "2006/01/02 15:04:05"
)

var (
	//startTimeUnix is 1900/01/01 00:00:00
	//startTimeUnix = uint64(0xFFFFFFFF7C558180)
	//startTime is 1900/01/01 00:00:00
	startTime = yearMonthDayDate(1900, 1, 1)
)

type calendar struct {
	loc     *time.Location
	time    time.Time
	lunar   *lunar
	solar   *solar
	isToday bool
}

func (c *calendar) Date() CalendarDate {
	st, ok := CheckSolarTermDay(c.Time())
	return CalendarDate{
		IsToday: c.isToday,
		Solar:   c.solar.Date(),
		Lunar:   c.lunar.Date(),
		EightCharacter: EightCharacter{
			NianZhu: NianZhu(c.Time()),
			YueZhu:  YueZhu(c.Time()),
			Rizhu:   RiZhu(c.Time()),
			ShiZhu:  ShiZhu(c.Time()),
		},
		Zodiac:         getZodiac(c.solar.year),
		Constellation:  getConstellation(c.time.Date()),
		IsSolarTermDay: ok,
		SolarTerm:      st,
	}
}

func isToday(t time.Time, now time.Time) bool {
	y, m, d := t.Date()
	ny, nm, nd := now.Date()
	return y == ny && m == nm && d == nd
}

func (c *calendar) FormatTime() string {
	return c.time.Format(DateFormatYMDHMS)
}

func (c *calendar) Time() time.Time {
	return c.time
}

func (c *calendar) LocalTime() time.Time {
	return c.time
}

func (c *calendar) String() string {
	vd, _ := json.Marshal(c.Date())
	return string(vd)
}

// Lunar ...
func (c *calendar) Lunar() Lunar {
	return c.lunar
}

// Solar ...
func (c *calendar) Solar() Solar {
	return c.solar
}

func (c *calendar) initializeCalendarDate() *calendar {
	if err := checkYearSupport(c.time.Year()); err != nil {
		panic(err)
	}
	c.isToday = isToday(c.time, time.Now())
	c.solar = solarByTime(c.time)
	c.lunar = lunarByTime(c.time)
	return c
}

//NewSolarCalendar can input three type of time to create the calendar
//"2006/01/02 03:04" format string
// time.Time value
// or nil to create a new time.Now() value
func NewSolarCalendar(v ...any) Calendar {
	var c *calendar
	if len(v) == 0 {
		return ParseTime(time.Now(), time.Local)
	}

	first := parseFirstArg(v)
	args := parseArgs(v)
	switch vv := first.(type) {
	case int:
		//todo(parseIntDate)
	case string:
		c = parseStringDate(vv, args...)
	case time.Time:
		c = parseTime(vv, time.Local)
	default:
		c = parseTime(time.Now(), time.Local)
	}
	return c.initializeCalendarDate()
}

// ParseSolarString returns Calendar parse from string(value,format?)
// @param string
// @param ...string
// @return Calendar
func ParseSolarString(s string, format ...string) Calendar {
	return parseStringDateFormat(s, format...).initializeCalendarDate()
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
	date := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
	return parseTime(date, time.Local).initializeCalendarDate()
}

// ParseSolarNow returns Calendar parse from solar time now(time.Now())
// @return Calendar
func ParseSolarNow() Calendar {
	return parseTime(time.Now(), time.Local).initializeCalendarDate()
}

// ParseSolarTime returns Calendar parse from solar time
// @param time.Time
// @return Calendar
func ParseSolarTime(t time.Time) Calendar {
	return parseTime(t, time.Local).initializeCalendarDate()
}

const minYear = 1900
const maxYear = 3000

func checkYearSupport(year int) error {
	if year < minYear || year > maxYear {
		return fmt.Errorf("[chronos] year %d not supported", year)
	}
	//if _, ok := solarTermTimes[year]; !ok {
	//	return fmt.Errorf("[chronos] year %d not supported", year)
	//}
	return nil
}

func yearDate(year int) time.Time {
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
}

func yearMonthDate(year int, month time.Month) time.Time {
	return time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
}

func yearMonthDayDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

package chronos

import (
	"encoding/json"
	"time"
)

// DefaultDateFormat ...
const DefaultDateFormat = "2006/01/02 15:04"
const LunarDateFormat = "2006/01/02"

type calendar struct {
	loc   *time.Location
	time  time.Time
	lunar *lunar
	solar *solar
}

func (c *calendar) FormatTime() string {
	return c.time.Format(DefaultDateFormat)
}

func (c *calendar) Time() time.Time {
	return c.time
}

func (c *calendar) LocalTime() time.Time {
	return c.time
}

func (c *calendar) ViewData() View {
	//todo(parse view data)
	return View{}
}

func (c *calendar) String() string {
	vd, _ := json.Marshal(c.ViewData())
	return string(vd)
}

// Lunar ...
func (c *calendar) Lunar() Lunar {
	return &lunar{}
}

// Solar ...
func (c *calendar) Solar() Solar {
	return &solar{}
}

func (c *calendar) initLunarAndSolar() *calendar {
	c.solar = &solar{}
	c.lunar = &lunar{}
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
	c.initLunarAndSolar()
	return c
}

func NewLunarCalendar() Calendar {
	return &calendar{}
}

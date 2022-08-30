package chronos

import (
	"encoding/json"
	"time"
)

// DefaultDateFormat ...
const DefaultDateFormat = "2006/01/02 15:04"
const LunarDateFormat = "2006/01/02"

type calendar struct {
	loc  *time.Location
	time time.Time
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
func (c *calendar) Lunar() *Lunar {
	return CalculateLunar(c.time.Format(DefaultDateFormat))
}

// Solar ...
func (c *calendar) Solar() *Solar {
	return &Solar{time: c.time}
}

// LunarDate ...
func (c *calendar) LunarDate() string {
	return c.Lunar().Date()
}

//NewSolarCalendar can input three type of time to create the calendar
//"2006/01/02 03:04" format string
// time.Time value
// or nil to create a new time.Now() value
func NewSolarCalendar(v ...any) Calendar {
	var c Calendar
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
		c = ParseTime(vv, time.Local)
	default:
		c = ParseTime(time.Now(), time.Local)
	}
	return c
}

func NewLunarCalendar() Calendar {

}

package chronos

import (
	"time"
)

// DefaultDateFormat ...
const DefaultDateFormat = "2006/01/02 15:04"
const LunarDateFormat = "2006/01/02"

type calendar struct {
	time time.Time
}

func (c *calendar) Time() time.Time {
	return c.time
}

//New can input three type of time to create the calendar
//"2006/01/02 03:04" format string
// time.Time value
// or nil to create a new time.Now() value
func New(v ...interface{}) Calendar {
	var c Calendar
	if v == nil {
		return ParseTime(time.Now())
	}

	switch vv := v[0].(type) {
	case string:
		c = parseDate(vv, v...)
	case time.Time:
		c = &calendar{vv}
	default:
		c = ParseTime(time.Now())
	}
	return c
}

func parseStringFormat(v interface{}) string {
	f := DefaultDateFormat
	df, ok := (v).(string)
	if ok {
		f = df
	}
	return f
}

func parseStringTime(v interface{}) string {
	t := ""
	df, ok := (v).(string)
	if ok {
		t = df
	}
	return t
}

func parseDate(t string, vv ...interface{}) *calendar {
	if t == "" {
		return &calendar{
			time: time.Now(),
		}
	}
	f := DefaultDateFormat
	if len(vv) > 1 {
		f = parseStringFormat(vv[1])
	}
	c := &calendar{
		time: time.Now(),
	}

	tt, err := time.Parse(f, t)
	if err == nil {
		c.time = tt
	}
	return c
}

func ParseTime(t time.Time) Calendar {
	return &calendar{
		time: t,
	}
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

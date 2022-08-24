package chronos

import (
	"time"
)

// DefaultDateFormat ...
const DefaultDateFormat = "2006/01/02 15:04"
const LunarDateFormat = "2006/01/02"

type calendar struct {
	loc  *time.Location
	time time.Time
}

func (c *calendar) String() string {
	return c.time.Format(DefaultDateFormat)
}

func (c *calendar) Time() time.Time {
	return c.time
}

func parseFirstArg(v []any) any {
	if len(v) > 0 {
		return v[0]
	}
	return nil
}

func parseArgs(v []any) []any {
	if len(v) > 1 {
		return v[1:]
	}
	return nil
}

//New can input three type of time to create the calendar
//"2006/01/02 03:04" format string
// time.Time value
// or nil to create a new time.Now() value
func New(v ...any) Calendar {
	var c Calendar
	first := parseFirstArg(v)
	args := parseArgs(v)
	switch vv := first.(type) {
	case string:
		c = parseStringDate(vv, args...)
	case time.Time:
		c = ParseTime(vv, time.Local)
	default:
		c = ParseTime(time.Now(), time.Local)
	}
	return c
}

func parseStringFormat(f string, v ...any) string {
	if len(v) == 0 {
		return f
	}
	df, ok := (v[0]).(string)
	if ok {
		f = df
	}
	return f
}

func parseStringTime(v any) string {
	t := ""
	df, ok := (v).(string)
	if ok {
		t = df
	}
	return t
}

func parseStringDate(t string, vv ...any) *calendar {
	c := &calendar{
		loc:  time.Local,
		time: time.Now(),
	}
	if t == "" {
		return c
	}

	f := parseStringFormat(DefaultDateFormat, vv...)
	tt, err := time.Parse(f, t)
	if err == nil {
		c.time = tt
	}
	return c
}

func ParseTime(t time.Time, local *time.Location) Calendar {
	return &calendar{
		loc:  local,
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

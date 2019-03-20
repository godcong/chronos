package chronos

import (
	"time"
)

// DateFormat ...
const DateFormat = "2006/01/02 15:04"

type calendar struct {
	time time.Time
}

// Calendar ...
type Calendar interface {
	Lunar() *Lunar
	Solar() *Solar
	LunarDate() string
}

// CalendarData ...
type CalendarData interface {
	Type() string
	Calendar() Calendar
}

//New can input three type of time to create the calendar
//"2006/01/02 03:04" format string
// time.Time value
// or nil to create a new time.Now() value
func New(v ...interface{}) Calendar {
	var c Calendar
	if v == nil {
		return &calendar{time.Now()}
	}
	switch vv := v[0].(type) {
	case string:
		c = formatDate(vv)
	case time.Time:
		c = &calendar{vv}
	}
	return c
}

func formatDate(s string) Calendar {
	t, err := time.Parse(DateFormat, s)
	if err != nil {
		t = time.Now()
	}
	return &calendar{
		time: t,
	}
}

// Lunar ...
func (c *calendar) Lunar() *Lunar {
	return CalculateLunar(c.time.Format(DateFormat))
}

// Solar ...
func (c *calendar) Solar() *Solar {
	return &Solar{time: c.time}
}

// LunarDate ...
func (c *calendar) LunarDate() string {
	return c.Lunar().Date()
}

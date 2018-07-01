package chronos

import (
	"time"
)

const DateFormat = "2006/01/02 15:04"

type calendar struct {
	time time.Time
}

type Calendar interface {
	Lunar() *Lunar
	Solar() *Solar
}

type CalendarData interface {
	Type() string
	Calendar() Calendar
}

//New can input two type of time to create the calendar
//"2006/01/02 03:04" format string or time.Time value
func New(v ...interface{}) Calendar {
	var c Calendar
	if v == nil {
		return &calendar{time.Now()}
	}
	switch vv := v[0].(type) {
	case string:
		c = formatString(vv)
	case time.Time:
		c = &calendar{vv}
	}
	return c
}

func formatString(s string) Calendar {
	t, err := time.Parse(DateFormat, s)
	if err != nil {
		t = time.Now()
	}
	return &calendar{
		time: t,
	}

}

//func NewCalendar(c CalendarData) Calendar {
//	if c != nil {
//		return c.Calendar()
//	}
//	return &calendar{
//		lunar: NewLunar(nil),
//		solar: NewSolar(nil),
//	}
//}

//func CalendarFromLunar(y, m, d int) Calendar {
//	return &calendar{
//		lunar: &Lunar{
//			year:  y,
//			month: m,
//			day:   d,
//		},
//	}
//}

//func CalendarFromSolar(time time.Time) Calendar {
//	return &calendar{
//		solar: &Solar{
//			time: time,
//		},
//	}
//}

func (c *calendar) Lunar() *Lunar {
	return CalculateLunar(c.time.Format(DateFormat))
}

func (c *calendar) Solar() *Solar {
	return &Solar{time: c.time}
}

func (c *calendar) LunarDate() string {
	return c.Lunar().Date()
}

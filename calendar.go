package chronos

import "time"

const DATE_FORMAT = "2006/01/02"

type calendar struct {
	lunar *Lunar
	solar *Solar
}

type Calendar interface {
	Lunar() *Lunar
	Solar() *Solar
}

type CalendarData interface {
	Type() string
	Calendar() Calendar
}

func NewCalendar(c CalendarData) Calendar {
	if c != nil {
		return c.Calendar()
	}
	return &calendar{
		lunar: NewLunar(nil),
		solar: NewSolar(nil),
	}
}

func CalendarFromLunar(y, m, d int) Calendar {
	return &calendar{
		lunar: &Lunar{
			year:  y,
			month: m,
			day:   d,
		},
	}
}

func CalendarFromSolar(time time.Time) Calendar {
	return &calendar{
		solar: &Solar{
			time: time,
		},
	}
}

func (c *calendar) Lunar() *Lunar {
	time := time.Now()
	if c.lunar != nil {
		return c.lunar
	}
	if c.solar != nil {
		time = c.solar.time
	}
	c.lunar = CalculateLunar(time.Format(DATE_FORMAT))
	return c.lunar
}

func (c *calendar) Solar() *Solar {
	if c.solar != nil {
		return c.solar
	}
	return nil
}

func (c *calendar) LunarDate() string {
	return c.Lunar().Date()
}

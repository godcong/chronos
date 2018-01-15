package lunar

import "time"

const DATE_FORMAT = "2006/01/02"

type Calendar struct {
	lunar *Lunar
	solar *Solar
}

func CalendarFromLunar(y, m, d int) Calendar {
	return Calendar{
		lunar: &Lunar{
			year:  y,
			month: m,
			day:   d,
		},
	}
}

func CalendarFromSolar(time time.Time) Calendar {
	return Calendar{
		solar: &Solar{
			time: time,
		},
	}
}

func (c *Calendar) Lunar() *Lunar {
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

package chronos

type LunarMonth struct {
	month          int
	dayCount       int
	firstJulianDay float64
}

func newLunarMonth(m, dc int, jd float64) *LunarMonth {
	return &LunarMonth{
		month:          m,
		dayCount:       dc,
		firstJulianDay: jd,
	}
}

func (l LunarMonth) FirstJulianDay() float64 {
	return l.firstJulianDay
}

func (l LunarMonth) DayCount() int {
	return l.dayCount
}

func (l LunarMonth) Month() int {
	return l.month
}

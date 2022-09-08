package utils

import (
	"time"
)

const Day = 24

// BetweenTime returns the time.Duration between two times
// @param time.Time
// @param time.Time
// @return time.Duration
func BetweenTime(d1, d2 time.Time) time.Duration {
	return d1.Sub(d2)
}

//BetweenDay returns the number of days between two times
func BetweenDay(current time.Time, start time.Time) int {
	return int(current.Sub(start).Hours() / 24)
}

func Lunar(date time.Time) {
	//lunarYear := 0
	//lunarMonth := 0
	//lunarDay := 0
	////solar := NewSolarFromDate(date)
	//hour := date.Hour()
	//minute := date.Minute()
	//second := date.Second()
	//currentYear := date.Year()
	//currentMonth := date.Month()
	//currentDay := date.Day()
	//ly := NewLunarYear(currentYear)
	//lunar := new(Lunar)
	//for i := ly.months.Front(); i != nil; i = i.Next() {
	//	m := i.Value.(*LunarMonth)
	//	firstDay := NewSolarFromJulianDay(m.GetFirstJulianDay())
	//	days := GetDaysBetween(firstDay.GetYear(), firstDay.GetMonth(), firstDay.GetDay(), currentYear, currentMonth, currentDay)
	//	if days < m.GetDayCount() {
	//		lunarYear = m.GetYear()
	//		lunarMonth = m.GetMonth()
	//		lunarDay = days + 1
	//		noon := NewSolarFromJulianDay(m.GetFirstJulianDay() + float64(lunarDay-1))
	//		lunar.solar = NewSolar(noon.GetYear(), noon.GetMonth(), noon.GetDay(), hour, minute, second)
	//		break
	//	}
	//}
	//lunar.year = lunarYear
	//lunar.month = lunarMonth
	//lunar.day = lunarDay
	//lunar.hour = hour
	//lunar.minute = minute
	//lunar.second = second
	//compute(lunar, ly)
	//return lunar
}

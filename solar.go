package chronos

import (
	"time"
)

// solar ...
type solar struct {
	t       time.Time
	year    int
	month   time.Month
	day     int
	hour    int
	minute  int
	second  int
	yearDay int
	weekDay time.Weekday
}

func (s *solar) Minute() int {
	return s.minute
}

func (s *solar) Second() int {
	return s.second
}

func (s *solar) YearDay() int {
	return s.yearDay
}

func (s *solar) IsLeapYear() bool {
	return s.year%4 == 0 && (s.year%100 != 0 || s.year%400 == 0)
}

func (s *solar) Year() int {
	return s.year
}

func (s *solar) Month() time.Month {
	return s.month
}

func (s *solar) Day() int {
	return s.day
}

func (s *solar) Hour() int {
	return s.hour
}

func (s *solar) Ymd() string {
	return s.t.Format(DateFormatYMD)
}

func (s *solar) YmdHms() string {
	return s.t.Format(DateFormatYMDHMS)
}

func (s *solar) Date() SolarDate {
	return SolarDate{
		Year:    s.year,
		Month:   s.month,
		Day:     s.day,
		Hour:    s.hour,
		Minute:  s.minute,
		Second:  s.second,
		WeekDay: s.weekDay,
	}
}

func solarByTime(t time.Time) *solar {
	s := &solar{
		t:       t,
		hour:    t.Hour(),
		minute:  t.Minute(),
		second:  t.Second(),
		yearDay: t.YearDay(),
		weekDay: t.Weekday(),
	}
	s.year, s.month, s.day = t.Date()
	return s
}

var _ Solar = &solar{}

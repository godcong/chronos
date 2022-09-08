package utils

import (
	"math"
	"time"
)

func JulianDayTime(julianDay float64) time.Time {
	d := int(julianDay + 0.5)
	f := julianDay + 0.5 - float64(d)

	if d >= 2299161 {
		c := int((float64(d) - 1867216.25) / 36524.25)
		d += 1 + c - c/4
	}
	d += 1524
	year := int((float64(d) - 122.1) / 365.25)
	d -= int(365.25 * float64(year))
	month := int(float64(d) / 30.601)
	d -= int(30.601 * float64(month))
	day := d
	if month > 13 {
		month -= 13
		year -= 4715
	} else {
		month -= 1
		year -= 4716
	}
	f *= 24
	hour := int(f)

	f -= float64(hour)
	f *= 60
	minute := int(f)

	f -= float64(minute)
	f *= 60
	second := int(math.Round(f))

	if second > 59 {
		second -= 60
		minute++
	}
	if minute > 59 {
		minute -= 60
		hour++
	}

	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
}

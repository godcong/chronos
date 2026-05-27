package chronos

import (
	"time"
)

// Zodiac represents one of the twelve Chinese Zodiac animals (生肖).
type Zodiac uint32

const defaultZodiac = "猫"

var zodiacChinese = [...]string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}

func (x Zodiac) Chinese() string {
	if x >= ZodiacMax {
		return ""
	}
	return zodiacChinese[x]
}

// YearZodiac returns the Zodiac for the given date, with LiChun (立春) boundary
// correction at second precision.
func YearZodiac(t time.Time, lichun time.Time) Zodiac {
	if t.Unix() >= lichun.Unix() {
		return getZodiac(t.Year())
	}
	return getZodiac(t.Year() - 1)
}

// YearZodiacDay returns the Zodiac for the given date, with LiChun boundary
// correction at day precision.
func YearZodiacDay(t time.Time, lichun time.Time) Zodiac {
	_, m, d := t.Date()
	_, sm, sd := lichun.Date()
	if m > sm || (m == sm && d >= sd) {
		return getZodiac(t.Year())
	}
	return getZodiac(t.Year() - 1)
}

// YearZodiacNoFix returns the Zodiac for the given year without LiChun boundary
// correction.
func YearZodiacNoFix(year int) Zodiac {
	return getZodiac(year)
}

func getZodiac(year int) Zodiac {
	return Zodiac((year - 4) % 12)
}

var _ ChineseSupport = Zodiac(0)

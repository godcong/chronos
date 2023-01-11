package chronos

import (
	"time"

	"github.com/godcong/chronos/v2/runes"
)

const defaultZodiac = "猫"

var zodiacs = runes.Runes("鼠牛虎兔龙蛇马羊猴鸡狗猪")

// Zodiac
// ENUM(Rat, Cow, Tiger, Rabbit, Dragon, Snake, Horse, Sheep, Monkey, Chicken, Dog, Pig, Max)
type Zodiac uint32

func (x Zodiac) Chinese() string {
	return ZodiacChineseV2(x)
}

// YearZodiac returns the zodiac of year.(ps: this will auto fix zodiac with LiChun check stopped at seconds)
// @param time.Time
// @return Zodiac
// @return error
func YearZodiac(t time.Time, lichun time.Time) Zodiac {
	if t.Unix() >= lichun.Unix() {
		return getZodiac(t.Year())
	}
	return getZodiac(t.Year() - 1)
}

// YearZodiacDay returns the zodiac of year.(ps: this will auto fix zodiac with LiChun check stopped at day)
// @param time.Time
// @return Zodiac
// @return error
func YearZodiacDay(t time.Time, lichun time.Time) Zodiac {
	_, m, d := t.Date()
	_, sm, sd := lichun.Date()
	if m >= sm && d >= sd {
		return getZodiac(t.Year())
	}
	return getZodiac(t.Year() - 1)
}

// YearZodiacNoFix returns the zodiac of year.(ps: this is not support LiChun day fix)
// @param int
// @return Zodiac
func YearZodiacNoFix(year int) Zodiac {
	return getZodiac(year)
}

func getZodiac(year int) Zodiac {
	return Zodiac((year - 4) % 12)
}

// ZodiacChineseV2 returns the chinese Zodiac string
// @param Zodiac
// @return string
func ZodiacChineseV2(zodiac Zodiac) string {
	return zodiacs.MustReadString(int(zodiac), 1)
}

// ZodiacChinese returns the chinese Zodiac string
// @param Zodiac
// @return string
// @return error
func ZodiacChinese(zodiac Zodiac) (string, error) {
	readString, err := zodiacs.ReadString(int(zodiac), 1)
	if err != nil {
		return "", ErrWrongZodiacTypes
	}
	return readString, nil
}

var _ ChineseSupport = Zodiac(0)

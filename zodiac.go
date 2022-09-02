package chronos

import (
	"errors"
	"time"

	"github.com/godcong/chronos/v2/runes"
)

const defaultZodiac = "猫"

var zodiacs = runes.Runes("鼠牛虎兔龙蛇马羊猴鸡狗猪")

// ErrWrongZodiacTypes returns an error
var ErrWrongZodiacTypes = errors.New("[chronos] wrong zodiac types")

//Zodiac
//ENUM(rat, cow, tiger, rabbit, dragon, snake, horse, sheep, monkey, chicken, dog, pig)
type Zodiac uint32

func (x Zodiac) Chinese() string {
	return ZodiacChineseV2(x)
}

// YearZodiac returns the zodiac of year.(pa: this will auto fix zodiac with LiChun check stopped at seconds)
// @param time.Time
// @return Zodiac
// @return error
func YearZodiac(t time.Time) (Zodiac, error) {
	if err := checkYearSupport(t.Year()); err != nil {
		return 0, err
	}
	if t.UTC().Unix() >= getYearSolarTermTime(t.Year(), SolarTermLiChun).Unix() {
		return getZodiac(t.Year()), nil
	}
	return getZodiac(t.Year() - 1), nil
}

// YearZodiacDay returns the zodiac of year.(pa: this will auto fix zodiac with LiChun check stopped at day)
// @param time.Time
// @return Zodiac
// @return error
func YearZodiacDay(t time.Time) (Zodiac, error) {
	if err := checkYearSupport(t.Year()); err != nil {
		return 0, err
	}
	_, m, d := t.UTC().Date()
	_, sm, sd := getYearSolarTermTime(t.Year(), SolarTermLiChun).Date()
	if m >= sm && d >= sd {
		return getZodiac(t.Year()), nil
	}
	return getZodiac(t.Year() - 1), nil
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

// YearZodiacNoFix ...
//func YearZodiacNoFix(lunar *lunar) string {
//	s := string([]rune(lunar.nianZhu(lunar.fixLiChun))[1])
//	for idx, v := range _DiZhiTable {
//		if strings.Compare(v, s) == 0 {
//			return zodiacs[idx]
//		}
//	}
//	return ""
//}

var _ ChineseSupport = Zodiac(0)

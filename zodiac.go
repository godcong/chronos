package chronos

import (
	"errors"
	"time"

	"github.com/godcong/chronos/v2/runes"
)

const defaultZodiac = "猫"

var zodiacs = runes.Runes("鼠牛虎兔龙蛇马羊猴鸡狗猪")

// ErrWrongZodiacTypes returns an error
var ErrWrongZodiacTypes = errors.New("wrong zodiac type error")

//Zodiac
//ENUM(rat, cow, tiger, rabbit, dragon, snake, horse, sheep, monkey, chicken, dog, pig)
type Zodiac uint32

// GetYearZodiac returns the zodiac of year.(pa: this will auto fix zodiac with lichun )
// @param time.Time
// @return Zodiac
// @return error
func GetYearZodiac(t time.Time) (Zodiac, error) {
	if err := checkYearSupport(t.Year()); err != nil {
		return 0, err
	}
	if t.UTC().Unix() > getSolarTermTime(t.Year(), SolarTermLiChun).Unix() {
		return getZodiac(t.Year()), nil
	}
	return getZodiac(t.Year() - 1), nil
}

// GetZodiac returns the zodiac of year.(ps: this is not support lichun day fix)
// @param int
// @return Zodiac
func GetZodiac(year int) Zodiac {
	return getZodiac(year)
}

func getZodiac(year int) Zodiac {
	return Zodiac(year%12 - 4)
}

func ZodiacChineseV2(zodiac Zodiac) string {
	readString, err := zodiacs.ReadString(int(zodiac), 1)
	if err != nil {
		return defaultZodiac
	}
	return readString
}

func ZodiacChinese(zodiac Zodiac) (string, error) {
	readString, err := zodiacs.ReadString(int(zodiac), 1)
	if err != nil {
		return "", ErrWrongZodiacTypes
	}
	return readString, nil
}

// GetZodiac ...
//func GetZodiac(lunar *lunar) string {
//	s := string([]rune(lunar.nianZhu(lunar.fixLiChun))[1])
//	for idx, v := range _DiZhiTable {
//		if strings.Compare(v, s) == 0 {
//			return zodiacs[idx]
//		}
//	}
//	return ""
//}

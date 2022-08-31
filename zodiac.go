package chronos

import (
	"errors"

	"github.com/godcong/chronos/v2/runes"
)

const defaultZodiac = "猫"

var zodiacs = runes.Runes("鼠牛虎兔龙蛇马羊猴鸡狗猪")

// ErrWrongZodiacTypes returns an error
var ErrWrongZodiacTypes = errors.New("wrong zodiac type error")

//Zodiac
//ENUM(rat, cow, tiger, rabbit, dragon, snake, horse, sheep, monkey, chicken, dog, pig)
type Zodiac uint32

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

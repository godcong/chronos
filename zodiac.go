package chronos

import (
	"errors"
)

var zodiacs = []rune("鼠牛虎兔龙蛇马羊猴鸡狗猪猫")

// ErrWrongZodiacTypes returns an error
var ErrWrongZodiacTypes = errors.New("error wrong zodiac types")

//Zodiac
//ENUM(rat, cow, tiger, rabbit, dragon, snake, horse, sheep, monkey, chicken, dog, pig, cat)
type Zodiac uint32

func ZodiacChinese(zodiac Zodiac) string {
	if zodiac >= ZodiacCat {
		zodiac = ZodiacCat
	}
	return string(zodiacs[int(zodiac)])
}

func ZodiacChineseV2(zodiac Zodiac) (string, error) {
	if zodiac >= ZodiacCat {
		return "", ErrWrongZodiacTypes
	}
	return string(zodiacs[int(zodiac)]), nil
}

// GetZodiac ...
//func GetZodiac(lunar *lunar) string {
//	s := string([]rune(lunar.nianZhu(lunar.fixLiChun))[1])
//	for idx, v := range earthyBranch {
//		if strings.Compare(v, s) == 0 {
//			return zodiacs[idx]
//		}
//	}
//	return ""
//}

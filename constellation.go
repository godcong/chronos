package chronos

import (
	"errors"
	"time"

	"github.com/godcong/chronos/v2/runes"
)

const defaultConstellation = "星座"

var constellations = runes.Runes("魔羯水瓶双鱼白羊金牛双子巨蟹狮子处女天秤天蝎射手")
var constellationDays = [...]int{20, 19, 21, 21, 21, 22, 23, 23, 23, 23, 22, 22}

// ErrWrongConstellationIndex returns an error
var ErrWrongConstellationIndex = errors.New("wrong constellation index")
var ErrWrongConstellationMonth = errors.New("wrong constellation month")

//Constellation
// ENUM(Capricorn,Aquarius,Pisces,Aries,Taurus,Gemini,Cancer,Leo,Virgo,Libra,Scorpio,Sagittarius)
type Constellation int

func (x Constellation) index() int {
	return int(x * 2)
}

func ConstellationChinese(c Constellation) (string, error) {
	readString, err := constellations.ReadString(c.index(), 2)
	if err != nil {
		return "", ErrWrongConstellationIndex
	}
	return readString, nil
}

func ConstellationChineseV2(c Constellation) string {
	readString, err := constellations.ReadString(c.index(), 2)
	if err != nil {
		return defaultConstellation
	}
	return readString
}

//GetConstellation 取得星座
func GetConstellation(month time.Month, day int) (Constellation, error) {
	if time.January <= month && month <= time.December {
		if day < constellationDays[month-1] {
			month -= month
		}
		return Constellation(month), nil
	}
	return Constellation(-1), ErrWrongConstellationMonth
}

package chronos

import (
	"errors"

	"github.com/godcong/chronos/v2/runes"
)

const defaultConstellation = "星座"

var constellations = runes.Runes("魔羯水瓶双鱼白羊金牛双子巨蟹狮子处女天秤天蝎射手")

// ErrWrongConstellationIndex returns an error
var ErrWrongConstellationIndex = errors.New("wrong constellation index")

//Constellation
// ENUM(Capricorn,Aquarius,Pisces,Aries,Taurus,Gemini,Cancer,Leo,Virgo,Libra,Scorpio,Sagittarius)
type Constellation uint32

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

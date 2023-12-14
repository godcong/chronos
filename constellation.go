package chronos

import (
	"time"

	"github.com/godcong/chronos/runes"
)

const defaultConstellation = "星座"

var constellations = runes.Runes("摩羯水瓶双鱼白羊金牛双子巨蟹狮子处女天秤天蝎射手")
var constellationDays = [ConstellationMax]int{20, 19, 21, 20, 21, 22, 23, 23, 23, 24, 23, 22}

var ()

//var ErrWrongConstellationMonth = errors.New("wrong constellation month")

// Constellation
// ENUM(Capricorn,Aquarius,Pisces,Aries,Taurus,Gemini,Cancer,Leo,Virgo,Libra,Scorpio,Sagittarius,Max)
type Constellation int

func (x Constellation) index() int {
	return int(x * 2)
}

func (x Constellation) Chinese() string {
	return ConstellationChineseV2(x)
}

// ConstellationChinese returns a constellation of the chinese
// @param Constellation
// @return string
// @return error
func ConstellationChinese(c Constellation) (string, error) {
	readString, err := constellations.ReadString(c.index(), 2)
	if err != nil {
		return "", ErrWrongConstellationTypes
	}
	return readString, nil
}

// ConstellationChineseV2 returns a constellation of the chinese
// @param Constellation
// @return string
func ConstellationChineseV2(c Constellation) string {
	return constellations.MustReadString(c.index(), 2)
}

// GetConstellation get the constellation of date
// @param time.Time
// @return Constellation
func GetConstellation(t time.Time) Constellation {
	return getConstellation(t.Date())
}

func getConstellation(_ int, month time.Month, day int) Constellation {
	if day < constellationDays[month-1] {
		month -= 1
	}
	month %= 12
	return Constellation(month)
}

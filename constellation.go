package chronos

import (
	"time"
)

// Constellation represents one of the twelve Western astrological constellations
// (星座).
type Constellation uint32

var constellationChinese = [...]string{"摩羯", "水瓶", "双鱼", "白羊", "金牛", "双子", "巨蟹", "狮子", "处女", "天秤", "天蝎", "射手"}

var constellationDays = [ConstellationMax]int{20, 19, 21, 20, 21, 22, 23, 23, 23, 24, 23, 22}

var constellationChineseMap = map[string]Constellation{
	"摩羯": ConstellationCapricorn, "水瓶": ConstellationAquarius,
	"双鱼": ConstellationPisces, "白羊": ConstellationAries,
	"金牛": ConstellationTaurus, "双子": ConstellationGemini,
	"巨蟹": ConstellationCancer, "狮子": ConstellationLeo,
	"处女": ConstellationVirgo, "天秤": ConstellationLibra,
	"天蝎": ConstellationScorpio, "射手": ConstellationSagittarius,
}

func (x Constellation) Chinese() string {
	if x < 0 || x >= ConstellationMax {
		return ""
	}
	return constellationChinese[x]
}

// GetConstellation returns the Constellation for the given date.
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

func constellationFromChinese(s string) Constellation {
	if c, ok := constellationChineseMap[s]; ok {
		return c
	}
	return ConstellationMax
}

var _ ChineseSupport = Constellation(0)

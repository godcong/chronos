package chronos

import (
	"github.com/6tail/lunar-go/calendar"
)

type eightChar struct {
	*calendar.EightChar
}

// EightCharIndex identifies a pillar position in the Eight Characters array.
type EightCharIndex int

const (
	EightCharYear EightCharIndex = iota
	EightCharMonth
	EightCharDay
	EightCharTime
	eightCharMax
)

var _ EightChar = (*eightChar)(nil)

func (e *eightChar) NaYin() [4]string {
	return [4]string{
		e.GetYearNaYin(),
		e.GetMonthNaYin(),
		e.GetDayNaYin(),
		e.GetTimeNaYin(),
	}
}

func (e *eightChar) FourPillars() [4]string {
	return [4]string{
		e.GetYear(),
		e.GetMonth(),
		e.GetDay(),
		e.GetTime(),
	}
}

func (e *eightChar) FiveElements() [4]string {
	return [4]string{
		e.GetYearWuXing(),
		e.GetMonthWuXing(),
		e.GetDayWuXing(),
		e.GetTimeWuXing(),
	}
}

func (e *eightChar) HiddenStems() [4][]string {
	return [4][]string{
		DiZhiHiddenStems[e.GetYearZhi()],
		DiZhiHiddenStems[e.GetMonthZhi()],
		DiZhiHiddenStems[e.GetDayZhi()],
		DiZhiHiddenStems[e.GetTimeZhi()],
	}
}

func (e *eightChar) TenGodsStems() [4]string {
	return [4]string{
		e.GetYearShiShenGan(),
		e.GetMonthShiShenGan(),
		e.GetDayShiShenGan(),
		e.GetTimeShiShenGan(),
	}
}

func (e *eightChar) TenGodsBranches() [4][]string {
	return [4][]string{
		listToStrings(e.GetYearShiShenZhi()),
		listToStrings(e.GetMonthShiShenZhi()),
		listToStrings(e.GetDayShiShenZhi()),
		listToStrings(e.GetTimeShiShenZhi()),
	}
}

func (e *eightChar) DaYun(sex int) []int {
	dayun := e.GetYun(sex).GetDaYunBy(11)
	if len(dayun) <= 0 {
		return nil
	}
	var result []int
	for i := 1; i < len(dayun); i++ {
		result = append(result, dayun[i].GetStartYear())
	}
	return result
}

// Pillar represents a single pillar in a BaZi chart, consisting of a Heavenly
// Stem (天干) and an Earthly Branch (地支).
type Pillar struct {
	Stem   string `json:"stem"`
	Branch string `json:"branch"`
}

// FourPillarsStruct represents the four pillars (四柱) of a BaZi chart: Year,
// Month, Day, and Hour.
type FourPillarsStruct struct {
	Year  Pillar `json:"year"`
	Month Pillar `json:"month"`
	Day   Pillar `json:"day"`
	Hour  Pillar `json:"hour"`
}

func (e *eightChar) ToFourPillars() FourPillarsStruct {
	siZhu := e.FourPillars()
	return FourPillarsFromArr(siZhu)
}

func FourPillarsFromArr(arr [4]string) FourPillarsStruct {
	pillar := func(s string) Pillar {
		runes := []rune(s)
		if len(runes) < 2 {
			return Pillar{}
		}
		return Pillar{Stem: string(runes[:1]), Branch: string(runes[1:])}
	}
	return FourPillarsStruct{
		Year:  pillar(arr[0]),
		Month: pillar(arr[1]),
		Day:   pillar(arr[2]),
		Hour:  pillar(arr[3]),
	}
}

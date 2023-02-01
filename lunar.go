package chronos

import (
	"time"

	"github.com/6tail/lunar-go/calendar"
)

// minLunarYear 最小可转换年
const minLunarYear = 1900

// maxLunarYear 最大可转换年
const maxLunarYear = 3000

// lunar ...
type lunar struct {
	*calendar.Lunar
}

func (l *lunar) GetSolarTerm() SolarTerm {
	jieQi := solarTerms.FindString(l.GetJieQi())
	if jieQi == 0 {
		return SolarTermMax
	}
	if !l.GetCurrentJieQi().IsJie() {
		return SolarTermMax
	}
	return SolarTerm(jieQi / 2)
}

func (l *lunar) GetSolarTermDetail() SolarTermDetail {
	return solarTermDetail(l.GetSolarTerm(), l.GetSolar().GetCalendar())
}

func (l *lunar) GetZodiac() Zodiac {
	jieQi := l.GetJieQiTable()
	liChun := jieQi["立春"]
	if liChun.GetYear() != l.GetSolar().GetYear() {
		liChun = jieQi["LI_CHUN"]
	}
	t := TimeFromYmdHms(liChun.GetYear(), (time.Month)(liChun.GetMonth()), liChun.GetDay(), liChun.GetHour(), liChun.GetMinute(), liChun.GetSecond())
	return YearZodiac(l.GetSolar().GetCalendar(), t)
}

func (l *lunar) GetEightChar() EightChar {
	return &eightChar{EightChar: l.Lunar.GetEightChar()}
}

func ParseLunarTime(t time.Time) Lunar {
	return &lunar{Lunar: calendar.NewLunarFromDate(t)}
}

var _ Lunar = &lunar{}

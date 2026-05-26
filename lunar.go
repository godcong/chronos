package chronos

import (
	"time"

	"github.com/6tail/lunar-go/calendar"
)

const minLunarYear = 1900

const maxLunarYear = 3000

type lunar struct {
	*calendar.Lunar
}

func (l *lunar) EightChar() EightChar {
	return &eightChar{EightChar: l.Lunar.GetEightChar()}
}

func (l *lunar) Zodiac() Zodiac {
	return l.GetZodiac()
}

func (l *lunar) YearXunKong() string {
	return l.GetYearXunKong()
}

func (l *lunar) MonthXunKong() string {
	return l.GetMonthXunKong()
}

func (l *lunar) DayXunKong() string {
	return l.GetDayXunKong()
}

func (l *lunar) TimeXunKong() string {
	return l.GetTimeXunKong()
}

func (l *lunar) JieQi() string {
	return l.GetJieQi()
}

func (l *lunar) JieQiTable() map[string]*calendar.Solar {
	return l.GetJieQiTable()
}

func (l *lunar) CurrentJieQi() *calendar.JieQi {
	return l.GetCurrentJieQi()
}

func (l *lunar) NextJie() *calendar.JieQi {
	return l.GetNextJie()
}

func (l *lunar) PrevJie() *calendar.JieQi {
	return l.GetPrevJie()
}

func (l *lunar) NextQi() *calendar.JieQi {
	return l.GetNextQi()
}

func (l *lunar) PrevQi() *calendar.JieQi {
	return l.GetPrevQi()
}

func (l *lunar) SolarTerm() SolarTerm {
	jieQi := solarTerms.FindString(l.GetJieQi())
	if jieQi == 0 {
		return SolarTermMax
	}
	if !l.GetCurrentJieQi().IsJie() {
		return SolarTermMax
	}
	return SolarTerm(jieQi / 2)
}

func (l *lunar) SolarTermDetail() SolarTermDetail {
	return solarTermDetail(l.SolarTerm(), l.GetSolar().ToYmdHms())
}

func (l *lunar) GetSolarTerm() SolarTerm {
	return l.SolarTerm()
}

func (l *lunar) GetSolarTermDetail() SolarTermDetail {
	return l.SolarTermDetail()
}

func (l *lunar) GetZodiac() Zodiac {
	jieQi := l.GetJieQiTable()
	liChun := jieQi["立春"]
	if liChun.GetYear() != l.GetSolar().GetYear() {
		liChun = jieQi["LI_CHUN"]
	}
	t := TimeFromYmdHms(liChun.GetYear(), (time.Month)(liChun.GetMonth()), liChun.GetDay(), 0, 0, 0)
	sl := l.GetSolar()
	t2 := TimeFromYmdHms(sl.GetYear(), (time.Month)(sl.GetMonth()), sl.GetDay(), 0, 0, 0)
	return YearZodiac(t2, t)
}

func (l *lunar) GetEightChar() EightChar {
	return &eightChar{EightChar: l.Lunar.GetEightChar()}
}

func ParseLunarTime(t time.Time) Lunar {
	return &lunar{Lunar: calendar.NewLunarFromDate(t)}
}

var _ Lunar = &lunar{}

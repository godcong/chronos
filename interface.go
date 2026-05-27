package chronos

import (
	"github.com/6tail/lunar-go/calendar"
)

// Calendar is the root interface for accessing both Solar and Lunar calendar
// data from a single date.
type Calendar interface {
	Lunar() Lunar
	Solar() Solar
}

// Solar provides Solar (Gregorian) calendar data and derived calculations.
type Solar interface {
	IsLeapYear() bool
	GetWeek() int
	GetWeekInChinese() string
	GetConstellation() Constellation
	GetFestivals() []string
	GetOtherFestivals() []string
	GetYear() int
	GetMonth() int
	GetDay() int
	GetHour() int
	GetMinute() int
	GetSecond() int
	GetJulianDay() float64
	ToYmd() string
	ToYmdHms() string
	String() string
	ToFullString() string
	Next(days int, onlyWorkDays bool) *calendar.Solar
	GetLunar() *calendar.Lunar
}

// BaziProvider provides the core data needed for BaZi (八字) analysis.
type BaziProvider interface {
	EightChar() EightChar
	Zodiac() Zodiac
	YearXunKong() string
	MonthXunKong() string
	DayXunKong() string
	TimeXunKong() string
}

// JieQiProvider provides Solar Term (节气) data and navigation.
type JieQiProvider interface {
	JieQi() string
	JieQiTable() map[string]*calendar.Solar
	CurrentJieQi() *calendar.JieQi
	NextJie() *calendar.JieQi
	PrevJie() *calendar.JieQi
	NextQi() *calendar.JieQi
	PrevQi() *calendar.JieQi
	SolarTerm() SolarTerm
	SolarTermDetail() SolarTermDetail
}

// Lunar provides comprehensive Lunar calendar data including Eight Characters,
// Zodiac, Solar Terms, and various traditional Chinese calendar attributes.
type Lunar interface {
	BaziProvider
	JieQiProvider
	GetGan() string
	GetYearGan() string
	GetYearGanByLiChun() string
	GetYearGanExact() string
	GetZhi() string
	GetYearZhi() string
	GetYearZhiByLiChun() string
	GetYearZhiExact() string
	GetYearInGanZhi() string
	GetYearInGanZhiByLiChun() string
	GetYearInGanZhiExact() string
	GetMonthGan() string
	GetMonthGanExact() string
	GetMonthZhi() string
	GetMonthZhiExact() string
	GetMonthInGanZhi() string
	GetMonthInGanZhiExact() string
	GetDayGan() string
	GetDayGanExact() string
	GetDayGanExact2() string
	GetDayZhi() string
	GetDayZhiExact() string
	GetDayZhiExact2() string
	GetDayInGanZhi() string
	GetDayInGanZhiExact() string
	GetDayInGanZhiExact2() string
	GetTimeGan() string
	GetTimeZhi() string
	GetTimeInGanZhi() string
	GetYearInChinese() string
	GetMonthInChinese() string
	GetDayInChinese() string
	GetSeason() string
	GetJie() string
	GetQi() string
	GetWeek() int
	GetWeekInChinese() string
	GetXiu() string
	GetXiuLuck() string
	GetXiuSong() string
	GetZheng() string
	GetAnimal() string
	GetGong() string
	GetShou() string
	GetFestivals() []string
	GetOtherFestivals() []string
	GetPengZuGan() string
	GetPengZuZhi() string
	GetPositionXi() string
	GetPositionXiDesc() string
	GetPositionYangGui() string
	GetPositionYangGuiDesc() string
	GetPositionYinGui() string
	GetPositionYinGuiDesc() string
	GetPositionFu() string
	GetPositionFuDesc() string
	GetPositionCai() string
	GetPositionCaiDesc() string
	GetDayPositionXi() string
	GetDayPositionXiDesc() string
	GetDayPositionYangGui() string
	GetDayPositionYangGuiDesc() string
	GetDayPositionYinGui() string
	GetDayPositionYinGuiDesc() string
	GetDayPositionFu() string
	GetDayPositionFuBySect(sect int) string
	GetDayPositionFuDesc() string
	GetDayPositionFuDescBySect(sect int) string
	GetDayPositionCai() string
	GetDayPositionCaiDesc() string
	GetYearPositionTaiSui() string
	GetYearPositionTaiSuiBySect(sect int) string
	GetYearPositionTaiSuiDesc() string
	GetYearPositionTaiSuiDescBySect(sect int) string
	GetMonthPositionTaiSuiBySect(sect int) string
	GetMonthPositionTaiSui() string
	GetMonthPositionTaiSuiDesc() string
	GetMonthPositionTaiSuiDescBySect(sect int) string
	GetDayPositionTaiSuiBySect(sect int) string
	GetDayPositionTaiSui() string
	GetDayPositionTaiSuiDesc() string
	GetDayPositionTaiSuiDescBySect(sect int) string
	GetTimePositionXi() string
	GetTimePositionXiDesc() string
	GetTimePositionYangGui() string
	GetTimePositionYangGuiDesc() string
	GetTimePositionYinGui() string
	GetTimePositionYinGuiDesc() string
	GetTimePositionFu() string
	GetTimePositionFuDesc() string
	GetTimePositionCai() string
	GetTimePositionCaiDesc() string
	GetChong() string
	GetDayChong() string
	GetChongGan() string
	GetDayChongGan() string
	GetChongGanTie() string
	GetDayChongGanTie() string
	GetChongShengXiao() string
	GetDayChongShengXiao() string
	GetChongDesc() string
	GetDayChongDesc() string
	GetSha() string
	GetDaySha() string
	GetYearNaYin() string
	GetMonthNaYin() string
	GetDayNaYin() string
	GetTimeNaYin() string
	GetEightChar() EightChar
	GetZhiXing() string
	GetDayTianShen() string
	GetTimeTianShen() string
	GetDayTianShenType() string
	GetTimeTianShenType() string
	GetDayTianShenLuck() string
	GetTimeTianShenLuck() string
	GetDayPositionTai() string
	GetMonthPositionTai() string
	GetTimeChong() string
	GetTimeSha() string
	GetTimeChongGan() string
	GetTimeChongGanTie() string
	GetTimeChongShengXiao() string
	GetTimeChongDesc() string
	GetSolarTerm() SolarTerm
	GetSolarTermDetail() SolarTermDetail
	GetDayYi() []string
	GetDayYiBySect(sect int) []string
	GetDayJi() []string
	GetDayJiBySect(sect int) []string
	GetDayJiShen() []string
	GetDayXiongSha() []string
	GetTimeYi() []string
	GetTimeJi() []string
	GetYueXiang() string
	GetYearNineStarBySect(sect int) *calendar.NineStar
	GetYearNineStar() *calendar.NineStar
	GetMonthNineStarBySect(sect int) *calendar.NineStar
	GetMonthNineStar() *calendar.NineStar
	GetDayNineStar() *calendar.NineStar
	GetTimeNineStar() *calendar.NineStar
	GetCurrentJieQi() *calendar.JieQi
	String() string
	ToFullString() string
	GetYear() int
	GetMonth() int
	GetDay() int
	GetHour() int
	GetMinute() int
	GetSecond() int
	GetTimeGanIndex() int
	GetTimeZhiIndex() int
	GetDayGanIndex() int
	GetDayGanIndexExact() int
	GetDayGanIndexExact2() int
	GetDayZhiIndex() int
	GetDayZhiIndexExact() int
	GetDayZhiIndexExact2() int
	GetMonthGanIndex() int
	GetMonthGanIndexExact() int
	GetMonthZhiIndex() int
	GetMonthZhiIndexExact() int
	GetYearGanIndex() int
	GetYearGanIndexByLiChun() int
	GetYearGanIndexExact() int
	GetYearZhiIndex() int
	GetYearZhiIndexByLiChun() int
	GetYearZhiIndexExact() int
	GetSolar() *calendar.Solar
	Next(days int) *calendar.Lunar
	GetYearXun() string
	GetYearXunByLiChun() string
	GetYearXunExact() string
	GetYearXunKong() string
	GetYearXunKongByLiChun() string
	GetYearXunKongExact() string
	GetMonthXun() string
	GetMonthXunExact() string
	GetMonthXunKong() string
	GetMonthXunKongExact() string
	GetDayXun() string
	GetDayXunExact() string
	GetDayXunExact2() string
	GetDayXunKong() string
	GetDayXunKongExact() string
	GetDayXunKongExact2() string
	GetTimeXun() string
	GetTimeXunKong() string
	GetShuJiu() *calendar.ShuJiu
	GetFu() *calendar.Fu
	GetLiuYao() string
	GetHou() string
	GetWuHou() string
	GetDayLu() string
	GetTime() *calendar.LunarTime
	GetTimes() []*calendar.LunarTime
	GetFoto() *calendar.Foto
	GetTao() *calendar.Tao
	GetJieQi() string
	GetJieQiTable() map[string]*calendar.Solar
	GetZodiac() Zodiac
}

// BaziLunar combines BaziProvider with additional Lunar data accessors needed
// for BaZi analysis.
type BaziLunar interface {
	BaziProvider
	GetEightChar() EightChar
	GetZodiac() Zodiac
	GetJieQiTable() map[string]*calendar.Solar
	GetSolar() *calendar.Solar
	String() string
}

// ChineseSupport is implemented by types that can return their Chinese string
// representation.
type ChineseSupport interface {
	Chinese() string
}

// EightChar provides access to the Eight Characters (八字) of a birth chart,
// including the Four Pillars, Five Elements, NaYin, Ten Gods, and Hidden Stems.
type EightChar interface {
	String() string
	FourPillars() [4]string
	FiveElements() [4]string
	NaYin() [4]string
	TenGodsStems() [4]string
	TenGodsBranches() [4][]string
	HiddenStems() [4][]string
	DaYun(sex int) []int
}

package chronos

import (
	"container/list"
	"time"

	"github.com/6tail/lunar-go/calendar"
)

// Calendar returns a calendarTime
type Calendar interface {
	Lunar() Lunar
	Solar() Solar
	//FormatTime() string
	//Time() time.Time
	//Date() CalendarDate
}

// Solar returns the solar time
type Solar interface {
	IsLeapYear() bool
	GetWeek() int
	GetWeekInChinese() string
	GetConstellation() Constellation
	GetFestivals() *list.List
	GetOtherFestivals() *list.List
	GetYear() int
	GetMonth() int
	GetDay() int
	GetHour() int
	GetMinute() int
	GetSecond() int
	GetCalendar() time.Time
	GetJulianDay() float64
	ToYmd() string
	ToYmdHms() string
	String() string
	ToFullString() string
	Next(days int) *calendar.Solar
	GetLunar() *calendar.Lunar
}

// Lunar returns the lunar time
type Lunar interface {
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
	GetZodiac() Zodiac
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
	GetFestivals() *list.List
	GetOtherFestivals() *list.List
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
	//GetJieQiTable() map[string]*calendar.Solar
	//GetJieQiList() *list.List
	GetDayYi() *list.List
	GetDayYiBySect(sect int) *list.List
	GetDayJi() *list.List
	GetDayJiBySect(sect int) *list.List
	GetDayJiShen() *list.List
	GetDayXiongSha() *list.List
	GetTimeYi() *list.List
	GetTimeJi() *list.List
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
}

// ChineseSupport implements the interface, if support chinese language output
type ChineseSupport interface {
	Chinese() string
}

// EightChar returns the lunar eight char
type EightChar interface {
	String() string
	//GetSect() int
	//SetSect(sect int)
	//GetDayGanIndex() int
	//GetDayZhiIndex() int
	//GetYear() string
	//GetYearGan() string
	//GetYearZhi() string
	//GetYearHideGan() []string
	//GetYearWuXing() string
	//GetYearNaYin() string
	//GetYearShiShenGan() string
	//GetYearShiShenZhi() *list.List
	//GetYearDiShi() string
	//GetMonth() string
	//GetMonthGan() string
	//GetMonthZhi() string
	//GetMonthHideGan() []string
	//GetMonthWuXing() string
	//GetMonthNaYin() string
	//GetMonthShiShenGan() string
	//GetMonthShiShenZhi() *list.List
	//GetMonthDiShi() string
	//GetDay() string
	//GetDayGan() string
	//GetDayZhi() string
	//GetDayHideGan() []string
	//GetDayWuXing() string
	//GetDayNaYin() string
	//GetDayShiShenGan() string
	//GetDayShiShenZhi() *list.List
	//GetDayDiShi() string
	//GetTime() string
	//GetTimeGan() string
	//GetTimeZhi() string
	//GetTimeHideGan() []string
	//GetTimeWuXing() string
	//GetTimeNaYin() string
	//GetTimeShiShenGan() string
	//GetTimeShiShenZhi() *list.List
	//GetTimeDiShi() string
	//GetTaiYuan() string
	//GetTaiYuanNaYin() string
	//GetTaiXi() string
	//GetTaiXiNaYin() string
	//GetMingGong() string
	//GetMingGongNaYin() string
	//GetShenGong() string
	//GetShenGongNaYin() string
	//GetLunar() *calendar.Lunar
	//GetYun(gender int) *calendar.Yun
	//GetYunBySect(gender int, sect int) *calendar.Yun
	//GetYearXun() string
	//GetYearXunKong() string
	//GetMonthXun() string
	//GetMonthXunKong() string
	//GetDayXun() string
	//GetDayXunKong() string
	//GetTimeXun() string
	//GetTimeXunKong() string
	GetWuXing() [4]string
	GetNaYin() [4]string
	GetSiZhu() [4]string
	GetShiShenGan() [4]string
	GetShiShenZhi() [4]string
}

package chronos

import (
	"sync"
	"time"

	"github.com/6tail/lunar-go/calendar"
)

// SolarTerm represents one of the twenty-four Solar Terms (节气) in the Chinese
// calendar.
type SolarTerm int

const defaultSolarTerm = "节气"

var solarTermChinese = [...]string{
	"小寒", "大寒", "立春", "雨水", "惊蛰", "春分",
	"清明", "谷雨", "立夏", "小满", "芒种", "夏至",
	"小暑", "大暑", "立秋", "处暑", "白露", "秋分",
	"寒露", "霜降", "立冬", "小雪", "大雪", "冬至",
}

var jieQiNames = [24]string{
	"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰",
	"春分", "清明", "谷雨", "立夏", "小满", "芒种",
	"夏至", "小暑", "大暑", "立秋", "处暑", "白露",
	"秋分", "寒露", "霜降", "立冬", "小雪", "大雪",
}

var solarTermChineseMap = map[string]SolarTerm{
	"小寒": SolarTermXiaoHan, "大寒": SolarTermDaHan,
	"立春": SolarTermLiChun, "雨水": SolarTermYuShui,
	"惊蛰": SolarTermJingZhe, "春分": SolarTermChunFen,
	"清明": SolarTermQingMing, "谷雨": SolarTermGuYu,
	"立夏": SolarTermLiXia, "小满": SolarTermXiaoMan,
	"芒种": SolarTermMangZhong, "夏至": SolarTermXiaZhi,
	"小暑": SolarTermXiaoShu, "大暑": SolarTermDaShu,
	"立秋": SolarTermLiQiu, "处暑": SolarTermChuShu,
	"白露": SolarTermBaiLu, "秋分": SolarTermQiuFen,
	"寒露": SolarTermHanLu, "霜降": SolarTermShuangJiang,
	"立冬": SolarTermLiDong, "小雪": SolarTermXiaoXue,
	"大雪": SolarTermDaXue, "冬至": SolarTermDongZhi,
}

var solarTermCache sync.Map

// SolarTermDetail holds detailed information about a solar term occurrence,
// including its index, date/time, three pentads (三候), and explanation.
type SolarTermDetail struct {
	Index       int       `json:"index"`
	SolarTerm   SolarTerm `json:"solar_term"`
	Time        string    `json:"time"`
	SanHou      string    `json:"san_hou"`
	Explanation string    `json:"explanation"`
}

func (x SolarTerm) Chinese() string {
	if x < 0 || x >= 24 {
		return ""
	}
	return solarTermChinese[x]
}

func (x SolarTerm) SanHou() string {
	return solarTermSanHous[x]
}

func (x SolarTerm) Explanation() string {
	return solarTermExplanations[x]
}

func solarTermDetail(st SolarTerm, ymdhms string) SolarTermDetail {
	return SolarTermDetail{
		Index:       int(st),
		SolarTerm:   st,
		Time:        ymdhms,
		SanHou:      solarTermSanHous[st],
		Explanation: solarTermExplanations[st],
	}
}

func (x SolarTerm) GetYearDate(year int) (month time.Month, day int) {
	_, month, day = getYearSolarTermTime(year, x).Date()
	return
}

// YearSolarTermDetail returns detailed information about a solar term in the
// specified year, including its date, three pentads (三候), and explanation.
func YearSolarTermDetail(t time.Time, st SolarTerm) (SolarTermDetail, error) {
	if st >= 24 {
		return SolarTermDetail{}, ErrWrongSolarTermFormat
	}
	if err := checkYearSupport(t.Year()); err != nil {
		return SolarTermDetail{}, err
	}
	ts := getYearSolarTermTime(t.Year(), st)
	return solarTermDetail(st, ts.Format(DateFormatYMDHMS)), nil
}

func YearSolarTermDate(t time.Time, st SolarTerm) (month time.Month, day int) {
	_, month, day = getYearSolarTermTime(t.Year(), st).Date()
	return
}

func YearSolarTermMonth(t time.Time, st SolarTerm) (month time.Month) {
	_, month, _ = getYearSolarTermTime(t.Year(), st).Date()
	return
}

func YearSolarTermDay(t time.Time, st SolarTerm) (day int) {
	_, _, day = getYearSolarTermTime(t.Year(), st).Date()
	return
}

func yearLiChunDay(year int) (day int) {
	_, _, day = getYearSolarTermTime(year, SolarTermLiChun).Date()
	return
}

func afterYearLiChunTime(t time.Time) bool {
	return getYearSolarTermTime(t.Year(), SolarTermLiChun).Sub(t) <= 0
}

// CheckSolarTermDay checks whether the given date falls on a solar term day,
// returning the SolarTerm if it does.
func CheckSolarTermDay(t time.Time) (SolarTerm, bool) {
	if err := checkYearSupport(t.Year()); err != nil {
		return SolarTermMax, false
	}

	var yst time.Time
	for i := 0; i < 24; i++ {
		yst = getYearSolarTermTime(t.Year(), SolarTerm(i))
		if yst.Month() == t.Month() && yst.Day() == t.Day() {
			return SolarTerm(i), true
		}
	}
	return SolarTermMax, false
}

func getYearSolarTermTime(year int, st SolarTerm) time.Time {
	return getYearSolarTermTimeFromLunar(year, st)
}

func getYearSolarTermTimeStr(year int, st SolarTerm) string {
	return getYearSolarTermTime(year, st).Format(DateFormatYMDHMS)
}

func solarTermFromChinese(s string) SolarTerm {
	if st, ok := solarTermChineseMap[s]; ok {
		return st
	}
	return SolarTermMax
}

func solarTermToSolar(year int, st SolarTerm) *solar {
	t := getYearSolarTermTime(year, st)
	return &solar{Solar: calendar.NewSolarFromDate(t)}
}

func getSolarTermDay(year int, month time.Month) (min, max int) {
	idx := (month - 1) * 2
	return getYearSolarTermTime(year, SolarTerm(idx)).Day(), getYearSolarTermTime(year, SolarTerm(idx)+1).Day()
}

func solarTermJieQiName(st SolarTerm) string {
	return jieQiNames[(int(st)+1)%24]
}

func loadYearSolarTermCache(year int) map[SolarTerm]time.Time {
	if v, ok := solarTermCache.Load(year); ok {
		return v.(map[SolarTerm]time.Time)
	}

	solar := calendar.NewSolar(year, 7, 1, 12, 0, 0)
	lunar := calendar.NewLunarFromSolar(solar)
	table := lunar.GetJieQiTable()

	m := make(map[SolarTerm]time.Time, 24)
	for i := 0; i < 24; i++ {
		st := SolarTerm(i)
		name := solarTermJieQiName(st)
		if s, ok := table[name]; ok && s.GetYear() == year {
			m[st] = time.Date(s.GetYear(), time.Month(s.GetMonth()), s.GetDay(), s.GetHour(), s.GetMinute(), s.GetSecond(), 0, loc)
		}
	}

	if s, ok := table["DONG_ZHI"]; ok && s.GetYear() == year {
		m[SolarTermDongZhi] = time.Date(s.GetYear(), time.Month(s.GetMonth()), s.GetDay(), s.GetHour(), s.GetMinute(), s.GetSecond(), 0, loc)
	}

	solarTermCache.Store(year, m)
	return m
}

func getYearSolarTermTimeFromLunar(year int, st SolarTerm) time.Time {
	m := loadYearSolarTermCache(year)
	return m[st]
}

var _ ChineseSupport = SolarTerm(0)

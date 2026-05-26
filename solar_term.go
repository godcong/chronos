package chronos

import (
	"sync"
	"time"

	"github.com/6tail/lunar-go/calendar"

	"github.com/godcong/chronos/v2/runes"
)

const defaultSolarTerm = "节气"

var solarTerms = runes.Runes("小寒大寒立春雨水惊蛰春分清明谷雨立夏小满芒种夏至小暑大暑立秋处暑白露秋分寒露霜降立冬小雪大雪冬至")

var jieQiNames = [24]string{
	"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰",
	"春分", "清明", "谷雨", "立夏", "小满", "芒种",
	"夏至", "小暑", "大暑", "立秋", "处暑", "白露",
	"秋分", "寒露", "霜降", "立冬", "小雪", "大雪",
}

var solarTermCache sync.Map

type SolarTerm uint32

type SolarTermDetail struct {
	Index       int       `json:"index"`
	SolarTerm   SolarTerm `json:"solar_term"`
	Time        string    `json:"time"`
	SanHou      string    `json:"san_hou"`
	Explanation string    `json:"explanation"`
}

func (x SolarTerm) index() int {
	return int(x * 2)
}

func (x SolarTerm) Chinese() string {
	return SolarTermChineseV2(x)
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

func SolarTermChineseV2(st SolarTerm) string {
	return solarTerms.MustReadString(st.index(), 2)
}

func SolarTermChinese(st SolarTerm) (string, error) {
	readString, err := solarTerms.ReadString(st.index(), 2)
	if err != nil {
		return "", ErrWrongSolarTermIndex
	}
	return readString, nil
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

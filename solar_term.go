package chronos

import (
	"errors"
	"fmt"
	"time"

	"github.com/godcong/chronos/v2/runes"
	"github.com/godcong/chronos/v2/utils"
)

const defaultSolarTerm = "节气"

var solarTerms = runes.Runes("小寒大寒立春雨水惊蛰春分清明谷雨立夏小满芒种夏至小暑大暑立秋处暑白露秋分寒露霜降立冬小雪大雪冬至")

var (
	// ErrSolarTermFormat returns an error
	ErrSolarTermFormat = errors.New("[chronos] solar term format not supported")
	// ErrWrongSolarTermIndex returns an error
	ErrWrongSolarTermIndex = errors.New("[chronos] wrong solar term index error")
)

// SolarTerm
//ENUM(XiaoHan,DaHan,LiChun,YuShui,JingZhe,ChunFen,QingMing,GuYu,LiXia,XiaoMan,MangZhong,XiaZhi,XiaoShu,DaShu,LiQiu,ChuShu,BaiLu,QiuFen,HanLu,ShuangJiang,LiDong,XiaoXue,DaXue,DongZhi,Max)
type SolarTerm uint32

// SolarTermDetail 24节气表
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

//func (x SolarTerm) detail() SolarTermDetail {
//	return SolarTermDetail{
//		Index:       int(x),
//		SolarTerm:   x,
//		SanHou:      solarTermSanHous[x],
//		Explanation: solarTermExplanations[x],
//	}
//}

func solarTermIndex(st SolarTerm) int {
	return int(st * 2)
}

func solarTermDetail(st SolarTerm, time string) SolarTermDetail {
	return SolarTermDetail{
		Index:       int(st),
		SolarTerm:   st,
		Time:        time,
		SanHou:      solarTermSanHous[st],
		Explanation: solarTermExplanations[st],
	}
}

func (x SolarTerm) GetYearDate(year int) (month time.Month, day int) {
	_, month, day = getYearSolarTermTime(year, x).Date()
	return
}

// YearSolarTermDetail get the details of year solar term
// @param time.Time
// @param SolarTerm
// @return SolarTermDetail
// @return error
func YearSolarTermDetail(t time.Time, st SolarTerm) (SolarTermDetail, error) {
	if st >= 24 {
		return SolarTermDetail{}, ErrSolarTermFormat
	}
	if err := checkYearSupport(t.Year()); err != nil {
		return SolarTermDetail{}, err
	}
	ts := getYearSolarTermTimeStr(t.Year(), st)
	return solarTermDetail(st, ts), nil
}

// YearSolarTermDate returns the year month day of the solar term
// @param time.Time
// @param SolarTerm
// @return month
// @return day
func YearSolarTermDate(t time.Time, st SolarTerm) (month time.Month, day int) {
	_, month, day = getYearSolarTermTime(t.Year(), st).Date()
	return
}

// YearSolarTermMonth returns the year month  of the solar term
// @param time.Time
// @param SolarTerm
// @return month
func YearSolarTermMonth(t time.Time, st SolarTerm) (month time.Month) {
	_, month, _ = getYearSolarTermTime(t.Year(), st).Date()
	return
}

// YearSolarTermDay returns the year day of the solar term
// @param time.Time
// @param SolarTerm
// @return day
func YearSolarTermDay(t time.Time, st SolarTerm) (day int) {
	_, _, day = getYearSolarTermTime(t.Year(), st).Date()
	return
}

func yearLiChunDay(year int) (day int) {
	fmt.Println("date", getYearSolarTermTime(year, SolarTermLiChun).Format(DefaultDateFormat))
	_, _, day = getYearSolarTermTime(year, SolarTermLiChun).Date()
	return
}

func CheckSolarTermDay(t time.Time) (SolarTerm, bool) {
	if err := checkYearSupport(t.Year()); err != nil {
		return SolarTermMax, false
	}

	var yst time.Time
	offset := yearOffset(t.Year())
	for i := 0; i < 24; i++ {
		yst = readSolarTermTime(offset, SolarTerm(i))
		if yst.Month() == t.Month() && yst.Day() == t.Day() {
			return SolarTerm(i), true
		}
	}
	return SolarTermMax, false
}

func getYearSolarTermTime(year int, st SolarTerm) time.Time {
	offset := yearOffset(year)
	return readSolarTermTime(offset, st)
}

func getYearSolarTermTimeStr(year int, st SolarTerm) string {
	offset := yearOffset(year)
	return readSolarTermTime(offset, st).Format(DefaultDateFormat)
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

func getSolarTermDay(year int, month time.Month) (min, max int) {
	year = yearOffset(year)
	idx := (month - 1) * 2
	return readSolarTermDay(year, SolarTerm(idx)), readSolarTermDay(year, SolarTerm(idx)+1)
}

func readSolarTermDay(offset int, st SolarTerm) int {
	sta := int(st) * 9
	return int(readYearSolarTermData(offset)[sta])
}

func readSolarTermTime(offset int, st SolarTerm) time.Time {
	sta := int(st)*9 + 1
	return utils.BytesToTime(readYearSolarTermData(offset)[sta : sta+8])
}

func readYearSolarTermData(offset int) []byte {
	sta := offset * SolarTermDataOffset
	return SolarTermData[sta : sta+SolarTermDataOffset]
}

var _ ChineseSupport = SolarTerm(0)

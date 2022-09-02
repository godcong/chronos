package chronos

import (
	"errors"
	"time"

	"github.com/godcong/chronos/v2/runes"
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
//ENUM(XiaoHan,DaHan,LiChun,YuShui,JingZhe,ChunFen,QingMing,GuYu,LiXia,XiaoMan,MangZhong,XiaZhi,XiaoShu,DaShu,LiQiu,ChuShu,BaiLu,QiuFen,HanLu,ShuangJiang,LiDong,XiaoXue,DaXue,DongZhi)
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
	_, month, day = getSolarTermTime(year, x).Date()
	return
}

// YearSolarTermDetail get the details of year solar term
// @param int
// @param SolarTerm
// @return SolarTermDetail
// @return error
func YearSolarTermDetail(year int, st SolarTerm) (SolarTermDetail, error) {
	if st >= 24 {
		return SolarTermDetail{}, ErrSolarTermFormat
	}
	if err := checkYearSupport(year); err != nil {
		return SolarTermDetail{}, err
	}
	t := getSolarTermTimeStr(year, st)
	return solarTermDetail(st, t), nil
}

func IsSolarTermDetailDay(t time.Time) bool {
	if _, ok := solarTermTimes[t.Year()]; !ok {
		return false
	}
	var tmpT time.Time
	for i := range solarTermTimes[t.Year()] {
		tmpT = getSolarTermTime(t.Year(), SolarTerm(i))
		if tmpT.Month() == t.Month() && tmpT.Day() == t.Day() {
			return true
		}
	}
	return false
}

func getSolarTermTime(year int, st SolarTerm) time.Time {
	return time.Unix(int64(solarTermTimes[year][st]), 0).UTC()
}

func getSolarTermTimeStr(year int, st SolarTerm) string {
	return time.Unix(int64(solarTermTimes[year][st]), 0).UTC().Format(DefaultDateFormat)
}

func SolarTermChineseV2(st SolarTerm) string {
	readString, err := solarTerms.ReadString(st.index(), 2)
	if err != nil {
		return defaultSolarTerm
	}
	return readString
}

func SolarTermChinese(st SolarTerm) (string, error) {
	readString, err := solarTerms.ReadString(st.index(), 2)
	if err != nil {
		return "", ErrWrongSolarTermIndex
	}
	return readString, nil
}

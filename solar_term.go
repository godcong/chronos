package chronos

import (
	"errors"

	"github.com/godcong/chronos/v2/runes"
)

const defaultSolarTerm = "节气"

var solarTerms = runes.Runes("小寒大寒立春雨水惊蛰春分清明谷雨立夏小满芒种夏至小暑大暑立秋处暑白露秋分寒露霜降立冬小雪大雪冬至")

// SolarTerm
//ENUM(XiaoHan,DaHan,LiChun,YuShui,JingZhe,ChunFen,QingMing,GuYu,LiXia,XiaoMan,MangZhong,XiaZhi,XiaoShu,DaShu,LiQiu,ChuShu,BaiLu,QiuFen,HanLu,ShuangJiang,LiDong,XiaoXue,DaXue,DongZhi)
type SolarTerm uint32

// SolarTermDetail 24节气表
type SolarTermDetail struct {
	Index       int       `json:"index"`
	SolarTerm   SolarTerm `json:"solar_term"`
	Time        int64     `json:"time"`
	SanHou      string    `json:"san_hou"`
	Explanation string    `json:"explanation"`
}

// ErrWrongSolarTermIndex returns an error
var ErrWrongSolarTermIndex = errors.New("wrong solar term index error")

func (x SolarTerm) index() int {
	return int(x * 2)
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

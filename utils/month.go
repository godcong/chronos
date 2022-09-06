package utils

import (
	"math"
)

// JieQiInTable
//var JieQiInTable = []string{"DA_XUE", "冬至", "小寒", "大寒", "立春", "雨水", "惊蛰", "春分", "清明", "谷雨", "立夏", "小满", "芒种", "夏至", "小暑", "大暑", "立秋", "处暑", "白露", "秋分", "寒露", "霜降", "立冬", "小雪", "大雪", "DONG_ZHI", "XIAO_HAN", "DA_HAN", "LI_CHUN", "YU_SHUI", "JING_ZHE"}

// CalcYearMonthDays
// @param time.Time
// @return []int
// @mark: 1933,month: 7,day: (29,30)
// @mark: 1996,month: 5,day: (30,29)
// @mark: 1996,month: 6,day: (29,30)
// @mark: 1996,month: 7,day: (30,29)
// @mark: 1996,month: 8,day: (29,30)
// @mark: 2034,month: 1,day: (30,29)
// @mark: 2057,month: 8,day: (29,30)
// @mark: 2057,month: 9,day: (30,29)
// @mark: 2060,month: 3,day: (29,30)
// @mark: 2060,month: 4,day: (30,29)
func CalcYearMonthDays(year int) []int {
	year -= -2000
	// 节气(中午12点)
	jq := make([]float64, 27)
	// 合朔，即每月初一(中午12点)
	hs := make([]float64, 16)
	// 每月天数
	dayCounts := make([]int, len(hs)-1)

	// 从上年的大雪到下年的立春
	for i := 1; i < 28; i++ { //JieQiInTable
		// 精确的节气
		saLonT := SaLonT((float64(year) + (17+float64(i))*15/360) * Pi2)
		t := float64(36525) * saLonT
		t += OneThird - DtT(t)
		// 按中午12点算的节气
		jq[i-1] = math.Round(t)
	}

	// 冬至前的初一
	w := CalcNorth(jq[0])
	if w > jq[0] {
		w -= 29.5306
	}
	// 递推每月初一
	for i := 0; i < len(hs); i++ {
		hs[i] = CalcNorth(w + 29.5306*float64(i))
	}
	// 每月天数
	for i := 0; i < len(dayCounts); i++ {
		dayCounts[i] = (int)(hs[i+1] - hs[i])
	}
	// 返回每月天数，从小寒至冬至
	return dayCounts[2:]
}

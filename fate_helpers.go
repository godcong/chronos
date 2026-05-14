package chronos

func calculateWuxingStrength(siZhu [4]string) map[string]*WuxingStrength {
	wuxingMap := map[string]string{
		"甲": "木", "乙": "木",
		"丙": "火", "丁": "火",
		"戊": "土", "己": "土",
		"庚": "金", "辛": "金",
		"壬": "水", "癸": "水",
		"子": "水", "丑": "土", "寅": "木", "卯": "木",
		"辰": "土", "巳": "火", "午": "火", "未": "土",
		"申": "金", "酉": "金", "戌": "土", "亥": "水",
	}

	cangGan := map[string][]struct {
		wuxing string
		weight float64
	}{
		"子": {{"水", 1.0}},
		"丑": {{"土", 0.6}, {"金", 0.2}, {"水", 0.2}},
		"寅": {{"木", 0.6}, {"火", 0.2}, {"土", 0.2}},
		"卯": {{"木", 1.0}},
		"辰": {{"土", 0.6}, {"木", 0.2}, {"水", 0.2}},
		"巳": {{"火", 0.6}, {"土", 0.2}, {"金", 0.2}},
		"午": {{"火", 0.7}, {"土", 0.3}},
		"未": {{"土", 0.6}, {"火", 0.2}, {"木", 0.2}},
		"申": {{"金", 0.6}, {"土", 0.2}, {"水", 0.2}},
		"酉": {{"金", 1.0}},
		"戌": {{"土", 0.6}, {"金", 0.2}, {"火", 0.2}},
		"亥": {{"水", 0.7}, {"木", 0.3}},
	}

	scores := map[string]float64{
		"木": 0, "火": 0, "土": 0, "金": 0, "水": 0,
	}

	for i, zhu := range siZhu {
		if len(zhu) < 2 {
			continue
		}
		tianGan := string([]rune(zhu)[:1])
		diZhi := string([]rune(zhu)[1:])

		if wx, ok := wuxingMap[tianGan]; ok {
			weight := 1.0
			if i == 2 {
				weight = 1.5
			}
			scores[wx] += weight
		}

		if cgs, ok := cangGan[diZhi]; ok {
			for _, cg := range cgs {
				scores[cg.wuxing] += cg.weight
			}
		}
	}

	result := make(map[string]*WuxingStrength)
	totalScore := 0.0
	for _, score := range scores {
		totalScore += score
	}

	for wx, score := range scores {
		result[wx] = &WuxingStrength{
			Wuxing:  wx,
			Score:   score,
			Percent: score / totalScore * 100,
		}
	}

	type wxScore struct {
		wx    string
		score float64
	}
	var sorted []wxScore
	for wx, score := range scores {
		sorted = append(sorted, wxScore{wx, score})
	}
	for i := 0; i < len(sorted); i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j].score > sorted[i].score {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	for rank, item := range sorted {
		result[item.wx].Rank = rank + 1
	}

	return result
}

func judgeRizhuQiangRuo(riZhuGan string, strengths map[string]*WuxingStrength) string {
	riZhuWuxing := getWuxingOfTianGan(riZhuGan)
	if riZhuWuxing == "" {
		return "中和"
	}

	tonglei := getTongleiWuxing(riZhuWuxing)
	tongleiScore := 0.0
	for _, wx := range tonglei {
		if s, ok := strengths[wx]; ok {
			tongleiScore += s.Score
		}
	}

	yidangScore := 0.0
	for wx, s := range strengths {
		isTonglei := false
		for _, t := range tonglei {
			if wx == t {
				isTonglei = true
				break
			}
		}
		if !isTonglei {
			yidangScore += s.Score
		}
	}

	ratio := tongleiScore / (tongleiScore + yidangScore) * 100
	if ratio > 60 {
		return "强"
	} else if ratio < 40 {
		return "弱"
	}
	return "中和"
}

func findTiaoHouShen(riZhuGan string, yueZhi string) []string {
	tiaoHouTable := map[string]map[string][]string{
		"甲": {
			"子": {"火"}, "丑": {"火"}, "寅": {"金", "火"}, "卯": {"金"},
			"辰": {"金", "火"}, "巳": {"水"}, "午": {"水"}, "未": {"水"},
			"申": {"丁", "庚"}, "酉": {"丁", "庚"}, "戌": {"水", "木"}, "亥": {"金", "火"},
		},
		"乙": {
			"子": {"火"}, "丑": {"火"}, "寅": {"火"}, "卯": {"金"},
			"辰": {"金", "火"}, "巳": {"水"}, "午": {"水"}, "未": {"水"},
			"申": {"丙", "癸"}, "酉": {"丙", "癸"}, "戌": {"水", "木"}, "亥": {"金", "火"},
		},
		"丙": {
			"子": {"壬"}, "丑": {"壬"}, "寅": {"壬", "庚"}, "卯": {"壬"},
			"辰": {"壬", "甲"}, "巳": {"壬"}, "午": {"壬"}, "未": {"壬"},
			"申": {"壬", "戊"}, "酉": {"壬", "戊"}, "戌": {"甲", "壬"}, "亥": {"戊", "丁"},
		},
		"丁": {
			"子": {"甲"}, "丑": {"甲"}, "寅": {"庚", "甲"}, "卯": {"庚"},
			"辰": {"甲", "庚"}, "巳": {"庚"}, "午": {"庚"}, "未": {"庚"},
			"申": {"甲", "戊"}, "酉": {"甲", "戊"}, "戌": {"甲", "壬"}, "亥": {"戊", "庚"},
		},
		"戊": {
			"子": {"丙"}, "丑": {"丙"}, "寅": {"丙", "甲"}, "卯": {"丙"},
			"辰": {"丙", "甲"}, "巳": {"壬"}, "午": {"壬"}, "未": {"壬"},
			"申": {"丙", "戊"}, "酉": {"丙", "戊"}, "戌": {"甲", "丙"}, "亥": {"甲", "丙"},
		},
		"己": {
			"子": {"丙"}, "丑": {"丙"}, "寅": {"丙", "甲"}, "卯": {"丙"},
			"辰": {"丙", "甲"}, "巳": {"癸"}, "午": {"癸"}, "未": {"癸"},
			"申": {"丙", "戊"}, "酉": {"丙", "戊"}, "戌": {"甲", "丙"}, "亥": {"甲", "丙"},
		},
		"庚": {
			"子": {"丁"}, "丑": {"丁"}, "寅": {"戊", "壬"}, "卯": {"丁"},
			"辰": {"甲", "丁"}, "巳": {"壬"}, "午": {"壬"}, "未": {"丁"},
			"申": {"丁", "甲"}, "酉": {"丁", "甲"}, "戌": {"甲", "壬"}, "亥": {"丁", "庚"},
		},
		"辛": {
			"子": {"丙"}, "丑": {"丙"}, "寅": {"己", "壬"}, "卯": {"壬"},
			"辰": {"壬", "甲"}, "巳": {"壬"}, "午": {"壬"}, "未": {"庚"},
			"申": {"壬", "戊"}, "酉": {"壬", "戊"}, "戌": {"甲", "壬"}, "亥": {"丙", "辛"},
		},
		"壬": {
			"子": {"戊"}, "丑": {"丙"}, "寅": {"庚", "丙"}, "卯": {"庚"},
			"辰": {"甲", "庚"}, "巳": {"壬"}, "午": {"癸"}, "未": {"辛"},
			"申": {"戊", "丁"}, "酉": {"戊", "丁"}, "戌": {"甲", "丙"}, "亥": {"戊", "庚"},
		},
		"癸": {
			"子": {"丙"}, "丑": {"丙"}, "寅": {"辛"}, "卯": {"庚"},
			"辰": {"丙", "辛"}, "巳": {"辛"}, "午": {"庚"}, "未": {"庚"},
			"申": {"丁", "甲"}, "酉": {"辛", "甲"}, "戌": {"甲", "辛"}, "亥": {"庚", "辛"},
		},
	}

	if tiaoHou, ok := tiaoHouTable[riZhuGan]; ok {
		if wx, ok := tiaoHou[yueZhi]; ok {
			return wx
		}
	}
	return []string{}
}

func calculateTongYiPoints(riZhuGan string, strengths map[string]*WuxingStrength) (float64, float64) {
	riZhuWuxing := getWuxingOfTianGan(riZhuGan)
	if riZhuWuxing == "" {
		return 0, 0
	}

	tonglei := getTongleiWuxing(riZhuWuxing)
	similarPoint := 0.0
	heteroPoint := 0.0

	for wx, s := range strengths {
		isTonglei := false
		for _, t := range tonglei {
			if wx == t {
				isTonglei = true
				break
			}
		}
		if isTonglei {
			similarPoint += s.Score
		} else {
			heteroPoint += s.Score
		}
	}

	return similarPoint, heteroPoint
}

func calculateWuXingFen(strengths map[string]*WuxingStrength) map[string]int {
	result := make(map[string]int)
	for wx, s := range strengths {
		result[wx] = int(s.Percent)
	}
	return result
}

func getTiaoHouTianGan(tiaoHouWuxing []string) []string {
	wuxingToTianGan := map[string][]string{
		"木": {"甲", "乙"},
		"火": {"丙", "丁"},
		"土": {"戊", "己"},
		"金": {"庚", "辛"},
		"水": {"壬", "癸"},
	}

	var result []string
	for _, wx := range tiaoHouWuxing {
		if tgs, ok := wuxingToTianGan[wx]; ok {
			result = append(result, tgs...)
		}
	}
	return result
}

func getWuxingOfTianGan(tianGan string) string {
	wuxingMap := map[string]string{
		"甲": "木", "乙": "木",
		"丙": "火", "丁": "火",
		"戊": "土", "己": "土",
		"庚": "金", "辛": "金",
		"壬": "水", "癸": "水",
	}
	return wuxingMap[tianGan]
}

func getTongleiWuxing(wuxing string) []string {
	return []string{wuxing, shengWoMap[wuxing]}
}

func joinStrings(strs []string, sep string) string {
	result := ""
	for i, s := range strs {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}

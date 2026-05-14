package chronos

import (
	"fmt"
	"time"
)

// FateInput is the input for fate calculation
type FateInput struct {
	BirthDate time.Time `json:"birth_date"`
	Gender    int       `json:"gender"` // 1=male, 2=female
	IsLunar   bool      `json:"is_lunar"`
	Surname   string    `json:"surname"`
}

// FateData contains all fate-related information for naming
type FateData struct {
	SolarDate  string          `json:"solar_date"`
	LunarDate  string          `json:"lunar_date"`
	Gender     int             `json:"gender"`
	Bazi       *BaziInfo       `json:"bazi"`
	WuxingXiji *WuxingXijiInfo `json:"wuxing_xiji"`
}

// BaziInfo contains eight character (BaZi) information
type BaziInfo struct {
	Sizhu       [4]string `json:"sizhu"`       // 四柱 (year, month, day, hour)
	Wuxing      [4]string `json:"wuxing"`      // 五行
	Nayin       [4]string `json:"nayin"`       // 纳音
	Shishen     [4]string `json:"shishen"`     // 十神
	Canggan     [4][]string `json:"canggan"`   // 藏干
	Xunkong     [4]string `json:"xunkong"`     // 旬空
	Zodiac      string    `json:"zodiac"`      // 生肖
	Constellation string  `json:"constellation"` // 星座
}

// WuxingXijiInfo contains five elements preference information
type WuxingXijiInfo struct {
	DayGan         string                `json:"day_gan"`          // 日干
	DayWuxing      string                `json:"day_wuxing"`       // 日主五行
	YueZhi         string                `json:"yue_zhi"`          // 月支
	QiangRuo       string                `json:"qiang_ruo"`        // 强弱 (强/弱/中和)
	SimilarPoint   float64               `json:"similar_point"`    // 同党得分
	HeteroPoint    float64               `json:"hetero_point"`     // 异党得分
	TongleiRatio   float64               `json:"tonglei_ratio"`    // 同党比例
	XiWuxing       []string              `json:"xi_wuxing"`        // 喜用五行
	YongWuxing     string                `json:"yong_wuxing"`      // 用神
	JiWuxing       []string              `json:"ji_wuxing"`        // 忌神
	TiaoHouWuxing  []string              `json:"tiao_hou_wuxing"`  // 调候五行
	TiaoHouTianGan []string              `json:"tiao_hou_tian_gan"` // 调候天干
	WuxingStrengths map[string]*WuxingStrength `json:"wuxing_strengths"` // 五行强度
	Analysis       string                `json:"analysis"`         // 分析说明
	SuggestWuxing  []string              `json:"suggest_wuxing"`   // 建议五行
	WuXingFen      map[string]int        `json:"wuxing_fen"`       // 五行分数
}

// WuxingStrength represents the strength of a five element
type WuxingStrength struct {
	Wuxing string  `json:"wuxing"`
	Score  float64 `json:"score"`
	Percent float64 `json:"percent"`
	Rank   int     `json:"rank"`
}

// FateError is a custom error type for fate calculation
type FateError struct {
	Code    int
	Message string
	Module  string
}

func (e *FateError) Error() string {
	return fmt.Sprintf("[%s] error %d: %s", e.Module, e.Code, e.Message)
}

// Error codes
const (
	ErrCodeInputInvalid  = 1001
	ErrCodeDateRange     = 1002
	ErrCodeCalculateBazi = 2001
	ErrCodeCalculateWuxing = 2002
)

// GetFateData is the main entry point for fate calculation
func GetFateData(input *FateInput) (*FateData, error) {
	if input == nil {
		return nil, &FateError{
			Code:    ErrCodeInputInvalid,
			Message: "input cannot be nil",
			Module:  "fate",
		}
	}

	if input.BirthDate.IsZero() {
		return nil, &FateError{
			Code:    ErrCodeInputInvalid,
			Message: "birth date cannot be empty",
			Module:  "fate",
		}
	}

	if input.BirthDate.Year() < 1900 || input.BirthDate.Year() > 2100 {
		return nil, &FateError{
			Code:    ErrCodeDateRange,
			Message: "birth date must be between 1900 and 2100",
			Module:  "fate",
		}
	}

	calendar := ParseSolarTime(input.BirthDate)

	baziInfo, err := calculateBazi(calendar)
	if err != nil {
		return nil, err
	}

	wuxingXiji := calculateWuxingXiji(calendar, baziInfo)

	return &FateData{
		SolarDate:  calendar.Solar().ToYmdHms(),
		LunarDate:  calendar.Lunar().String(),
		Gender:     input.Gender,
		Bazi:       baziInfo,
		WuxingXiji: wuxingXiji,
	}, nil
}

func calculateBazi(calendar Calendar) (*BaziInfo, error) {
	lunar := calendar.Lunar()
	eightChar := lunar.GetEightChar()

	siZhu := eightChar.GetSiZhu()
	wuXing := eightChar.GetWuXing()
	naYin := eightChar.GetNaYin()
	shiShen := eightChar.GetShiShenGan()
	cangGan := eightChar.GetCangGan()

	zodiac := lunar.GetZodiac().String()
	constellation := calendar.Solar().GetConstellation().String()

	// Calculate xunkong for each pillar
	xunkong := [4]string{
		lunar.GetYearXunKong(),
		lunar.GetMonthXunKong(),
		lunar.GetDayXunKong(),
		lunar.GetTimeXunKong(),
	}

	return &BaziInfo{
		Sizhu:       siZhu,
		Wuxing:      wuXing,
		Nayin:       naYin,
		Shishen:     shiShen,
		Canggan:     cangGan,
		Xunkong:     xunkong,
		Zodiac:      zodiac,
		Constellation: constellation,
	}, nil
}

func calculateWuxingXiji(calendar Calendar, baziInfo *BaziInfo) *WuxingXijiInfo {
	siZhu := baziInfo.Sizhu

	// Calculate five elements strengths
	strengths := calculateWuxingStrength(siZhu)

	// Determine day master strength
	riZhuGan := string([]rune(baziInfo.Sizhu[2])[:1])
	yueZhi := string([]rune(baziInfo.Sizhu[1])[1:])

	qiangRuo := judgeRizhuQiangRuo(riZhuGan, strengths)

	// Find tiao hou yong shen
	tiaoHou := findTiaoHouShen(riZhuGan, yueZhi)

	// Calculate similar and hetero points
	similarPoint, heteroPoint := calculateTongYiPoints(riZhuGan, strengths)

	// Determine xi yong ji wu xing
	xiWuxing, yongWuxing, jiWuxing := determineXiYongJi(qiangRuo, strengths, tiaoHou)

	// Calculate wu xing fen
	wuXingFen := calculateWuXingFen(strengths)

	// Generate analysis
	analysis := generateAnalysis(riZhuGan, qiangRuo, xiWuxing, jiWuxing, strengths)

	// Suggest wu xing
	suggestWuxing := xiWuxing

	return &WuxingXijiInfo{
		DayGan:    riZhuGan,
		DayWuxing:      getWuxingOfTianGan(riZhuGan),
		YueZhi:    yueZhi,
		QiangRuo:       qiangRuo,
		SimilarPoint:   similarPoint,
		HeteroPoint:    heteroPoint,
		TongleiRatio:   similarPoint / (similarPoint + heteroPoint) * 100,
		XiWuxing:       xiWuxing,
		YongWuxing:     yongWuxing,
		JiWuxing:       jiWuxing,
		TiaoHouWuxing:  tiaoHou,
		TiaoHouTianGan: getTiaoHouTianGan(tiaoHou),
		WuxingStrengths: strengths,
		Analysis:       analysis,
		SuggestWuxing:  suggestWuxing,
		WuXingFen:      wuXingFen,
	}
}

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

	// Hidden stems in earthly branches
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

		// Tian gan weight
		if wx, ok := wuxingMap[tianGan]; ok {
			weight := 1.0
			if i == 2 { // Day master has higher weight
				weight = 1.5
			}
			scores[wx] += weight
		}

		// Di zhi weight (including hidden stems)
		if cgs, ok := cangGan[diZhi]; ok {
			for _, cg := range cgs {
				scores[cg.wuxing] += cg.weight
			}
		}
	}

	// Convert to WuxingStrength map
	result := make(map[string]*WuxingStrength)
	totalScore := 0.0
	for _, score := range scores {
		totalScore += score
	}

	for wx, score := range scores {
		result[wx] = &WuxingStrength{
			Wuxing: wx,
			Score:  score,
			Percent: score / totalScore * 100,
		}
	}

	// Rank by score
	type wxScore struct {
		wx    string
		score float64
	}
	var sorted []wxScore
	for wx, score := range scores {
		sorted = append(sorted, wxScore{wx, score})
	}
	// Simple bubble sort
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

	// Same party: generates day master + same as day master
	tonglei := getTongleiWuxing(riZhuWuxing)
	tongleiScore := 0.0
	for _, wx := range tonglei {
		if s, ok := strengths[wx]; ok {
			tongleiScore += s.Score
		}
	}

	// Different party: controls day master + controlled by day master + generated by day master
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
	// Tiao hou table: day master x month branch
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

func determineXiYongJi(qiangRuo string, strengths map[string]*WuxingStrength, tiaoHou []string) ([]string, string, []string) {
	// Find the weakest five element
	var weakestWuxing string
	minScore := 999.0
	for wx, s := range strengths {
		if s.Score < minScore {
			minScore = s.Score
			weakestWuxing = wx
		}
	}

	var xiWuxing []string
	var yongWuxing string
	var jiWuxing []string

	switch qiangRuo {
	case "强":
		// Strong day master: need to control/weaken
		xiWuxing = getKeXieHaoWuxing(weakestWuxing)
		yongWuxing = weakestWuxing
		jiWuxing = getTongleiWuxing(weakestWuxing)
	case "弱":
		// Weak day master: need to support/strengthen
		xiWuxing = getShengZhuWuxing(weakestWuxing)
		yongWuxing = weakestWuxing
		jiWuxing = getKeZhuWuxing(weakestWuxing)
	default:
		// Neutral: balance
		xiWuxing = tiaoHou
		if len(xiWuxing) == 0 {
			xiWuxing = []string{weakestWuxing}
		}
		yongWuxing = weakestWuxing
		jiWuxing = []string{}
	}

	return xiWuxing, yongWuxing, jiWuxing
}

func calculateWuXingFen(strengths map[string]*WuxingStrength) map[string]int {
	result := make(map[string]int)
	for wx, s := range strengths {
		result[wx] = int(s.Percent)
	}
	return result
}

func generateAnalysis(riZhuGan, qiangRuo string, xiWuxing, jiWuxing []string, strengths map[string]*WuxingStrength) string {
	return fmt.Sprintf("日主%s，五行%s，格局%s。喜用神为%s，忌神为%s。",
		riZhuGan, getWuxingOfTianGan(riZhuGan), qiangRuo,
		joinStrings(xiWuxing, "、"), joinStrings(jiWuxing, "、"))
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

// Helper functions
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
	// Same party: generates this wuxing + same as this wuxing
	shengWo := map[string]string{
		"木": "水", "火": "木", "土": "火", "金": "土", "水": "金",
	}
	return []string{wuxing, shengWo[wuxing]}
}

func getKeXieHaoWuxing(wuxing string) []string {
	// Control, drain, exhaust
	woKe := map[string]string{
		"木": "土", "火": "金", "土": "水", "金": "木", "水": "火",
	}
	woSheng := map[string]string{
		"木": "火", "火": "土", "土": "金", "金": "水", "水": "木",
	}
	keWo := map[string]string{
		"木": "金", "火": "水", "土": "木", "金": "火", "水": "土",
	}
	return []string{woKe[wuxing], woSheng[wuxing], keWo[wuxing]}
}

func getShengZhuWuxing(wuxing string) []string {
	// Generate day master + same as day master
	return getTongleiWuxing(wuxing)
}

func getKeZhuWuxing(wuxing string) []string {
	// Control day master
	keWo := map[string]string{
		"木": "金", "火": "水", "土": "木", "金": "火", "水": "土",
	}
	return []string{keWo[wuxing]}
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

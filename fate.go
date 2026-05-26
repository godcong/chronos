package chronos

import (
	"fmt"
	"time"
)

type FateInput struct {
	BirthDate time.Time    `json:"birth_date"`
	Gender    int          `json:"gender"`
	IsLunar   bool         `json:"is_lunar"`
	Surname   string       `json:"surname"`
	Method    XiYongMethod `json:"method"`
}

type FateData struct {
	SolarDate  string          `json:"solar_date"`
	LunarDate  string          `json:"lunar_date"`
	Gender     int             `json:"gender"`
	Bazi       *BaziInfo       `json:"bazi"`
	WuxingXiji *WuxingXijiInfo `json:"wuxing_xiji"`
}

type BaziInfo struct {
	FourPillars  [4]string   `json:"four_pillars"`
	FiveElements [4]string   `json:"five_elements"`
	NaYin        [4]string   `json:"na_yin"`
	TenGods      [4]string   `json:"ten_gods"`
	HiddenStems  [4][]string `json:"hidden_stems"`
	VoidPillars  [4]string   `json:"void_pillars"`
	Zodiac       string      `json:"zodiac"`
	Constellation string     `json:"constellation"`
}

type WuxingXijiInfo struct {
	DayGan            string                     `json:"day_gan"`
	DayWuxing         string                     `json:"day_wuxing"`
	YueZhi            string                     `json:"yue_zhi"`
	Strength          string                     `json:"strength"`
	SimilarPoint      float64                    `json:"similar_point"`
	HeteroPoint       float64                    `json:"hetero_point"`
	TongleiRatio      float64                    `json:"tonglei_ratio"`
	FavorableElements []string                   `json:"favorable_elements"`
	UsefulElement     string                     `json:"useful_element"`
	UnfavorableElements []string                 `json:"unfavorable_elements"`
	HostileElements   []string                   `json:"hostile_elements"`
	IdleElements      []string                   `json:"idle_elements"`
	AdjustmentElements []string                  `json:"adjustment_elements"`
	AdjustmentStems   []string                   `json:"adjustment_stems"`
	WuxingStrengths   map[string]*WuxingStrength `json:"wuxing_strengths"`
	Analysis          string                     `json:"analysis"`
	SuggestWuxing     []string                   `json:"suggest_wuxing"`
	ElementScores     map[string]int             `json:"element_scores"`
	Method            XiYongMethod               `json:"method"`
	MethodName        string                     `json:"method_name"`
	GeJu              *GeJuInfo                  `json:"geju,omitempty"`
}

type WuxingStrength struct {
	Element string  `json:"element"`
	Score   float64 `json:"score"`
	Percent float64 `json:"percent"`
	Rank    int     `json:"rank"`
}

type FateError struct {
	Code    int
	Message string
	Module  string
}

func (e *FateError) Error() string {
	return fmt.Sprintf("[%s] error %d: %s", e.Module, e.Code, e.Message)
}

const (
	ErrCodeInputInvalid    = 1001
	ErrCodeDateRange       = 1002
	ErrCodeCalculateBazi   = 2001
	ErrCodeCalculateWuxing = 2002
)

func GetFateData(input *FateInput) (*FateData, error) {
	if input == nil {
		return nil, &FateError{Code: ErrCodeInputInvalid, Message: "input cannot be nil", Module: "fate"}
	}
	if input.BirthDate.IsZero() {
		return nil, &FateError{Code: ErrCodeInputInvalid, Message: "birth date cannot be empty", Module: "fate"}
	}
	if input.BirthDate.Year() < 1900 || input.BirthDate.Year() > 2100 {
		return nil, &FateError{Code: ErrCodeDateRange, Message: "birth date must be between 1900 and 2100", Module: "fate"}
	}

	method := input.Method
	if method < XiYongMethodBalance || method > XiYongMethodGeJu {
		method = XiYongMethodBalance
	}

	calendar := ParseSolarTime(input.BirthDate)
	baziInfo, err := calculateBazi(calendar)
	if err != nil {
		return nil, err
	}
	wuxingXiji := calculateWuxingXiji(baziInfo, method)

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
	zodiacObj := lunar.GetZodiac()
	zodiac, _ := ZodiacChinese(zodiacObj)
	constellationObj := calendar.Solar().GetConstellation()
	constellation, _ := ConstellationChinese(constellationObj)
	xunkong := [4]string{
		lunar.GetYearXunKong(),
		lunar.GetMonthXunKong(),
		lunar.GetDayXunKong(),
		lunar.GetTimeXunKong(),
	}

	return &BaziInfo{
		FourPillars:   siZhu,
		FiveElements:  wuXing,
		NaYin:         naYin,
		TenGods:       shiShen,
		HiddenStems:   cangGan,
		VoidPillars:   xunkong,
		Zodiac:        zodiac,
		Constellation: constellation,
	}, nil
}

func calculateWuxingXiji(baziInfo *BaziInfo, method XiYongMethod) *WuxingXijiInfo {
	siZhu := baziInfo.FourPillars
	strengths := calculateWuxingStrength(siZhu)

	riZhuGan := string([]rune(baziInfo.FourPillars[2])[:1])
	yueZhi := string([]rune(baziInfo.FourPillars[1])[1:])
	qiangRuo := judgeRizhuQiangRuo(riZhuGan, strengths)
	tiaoHou := findTiaoHouShen(riZhuGan, yueZhi)
	similarPoint, heteroPoint := calculateTongYiPoints(riZhuGan, strengths)
	wuXingFen := calculateWuXingFen(strengths)

	var xyj *XiYongJiChou
	var geJu *GeJuInfo
	var analysis string
	var methodName string

	switch method {
	case XiYongMethodGeJu:
		geJu = determineGeJu(riZhuGan, yueZhi, siZhu)
		xyj = geJuXiYongJi(riZhuGan, geJu, qiangRuo, strengths, tiaoHou)
		xyj.IdleElements = findXianWuxing(xyj)
		analysis = generateGeJuAnalysis(riZhuGan, qiangRuo, geJu, xyj)
		methodName = "格局用神法"
	default:
		xyj = balanceXiYongJi(riZhuGan, qiangRuo, strengths, tiaoHou)
		xyj.IdleElements = findXianWuxing(xyj)
		analysis = generateBalanceAnalysis(riZhuGan, qiangRuo, xyj)
		methodName = "平衡用神法"
	}

	suggestWuxing := append([]string{xyj.UsefulElement}, xyj.FavorableElements...)

	return &WuxingXijiInfo{
		DayGan:            riZhuGan,
		DayWuxing:         wuxingOfTianGan(riZhuGan),
		YueZhi:            yueZhi,
		Strength:          qiangRuo,
		SimilarPoint:      similarPoint,
		HeteroPoint:       heteroPoint,
		TongleiRatio:      similarPoint / (similarPoint + heteroPoint) * 100,
		FavorableElements: xyj.FavorableElements,
		UsefulElement:     xyj.UsefulElement,
		UnfavorableElements: xyj.UnfavorableElements,
		HostileElements:   xyj.HostileElements,
		IdleElements:      xyj.IdleElements,
		AdjustmentElements: tiaoHou,
		AdjustmentStems:   getTiaoHouTianGan(tiaoHou),
		WuxingStrengths:   strengths,
		Analysis:          analysis,
		SuggestWuxing:     suggestWuxing,
		ElementScores:     wuXingFen,
		Method:            method,
		MethodName:        methodName,
		GeJu:              geJu,
	}
}

func generateBalanceAnalysis(riZhuGan, qiangRuo string, xyj *XiYongJiChou) string {
	return fmt.Sprintf("日主%s，五行%s，格局%s（平衡用神法）。用神为%s，喜神为%s，忌神为%s，仇神为%s。",
		riZhuGan, wuxingOfTianGan(riZhuGan), qiangRuo,
		xyj.UsefulElement, joinStrings(xyj.FavorableElements, "、"),
		joinStrings(xyj.UnfavorableElements, "、"), joinStrings(xyj.HostileElements, "、"))
}

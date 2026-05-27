package fate

import (
	"github.com/godcong/chronos/v2"
)

// XiYongMethod specifies which method to use for determining favorable elements.
type XiYongMethod int

const (
	// XiYongMethodBalance uses the Balance (平衡用神) method: strengthens weak
	// elements and weakens strong ones.
	XiYongMethodBalance XiYongMethod = iota
	// XiYongMethodGeJu uses the Pattern (格局用神) method: determines favorable
	// elements based on the monthly branch's Ten God pattern.
	XiYongMethodGeJu
)

// FateInput specifies the input parameters for fate analysis: birth date,
// gender, and calculation method.
type FateInput struct {
	Calendar      chronos.Calendar
	Gender        int
	XiYongMethod  XiYongMethod
}

// FateData contains the complete fate analysis result including BaZi information
// and Five Element analysis.
type FateData struct {
	BaziInfo       BaziInfo       `json:"bazi_info"`
	WuxingXijiInfo WuxingXijiInfo `json:"wuxing_xiji_info"`
	WuxingStrength WuxingStrength `json:"wuxing_strength"`
	XiYongJiChou   XiYongJiChou   `json:"xi_yong_ji_chou"`
	GeJuInfo       *GeJuInfo      `json:"ge_ju_info,omitempty"`
}

// BaziInfo contains the Four Pillars and derived BaZi information.
type BaziInfo struct {
	SiZhu        [4]string   `json:"si_zhu"`
	WuXing       [4]string   `json:"wu_xing"`
	NaYin        [4]string   `json:"na_yin"`
	ShiShenGan   [4]string   `json:"shi_shen_gan"`
	ShiShenZhi   [4][]string `json:"shi_shen_zhi"`
	CangGan      [4][]string `json:"cang_gan"`
	DaYun        []int       `json:"da_yun"`
	Zodiac       string      `json:"zodiac"`
	Constellation string     `json:"constellation"`
}

// WuxingXijiInfo contains the Five Element favorability analysis results.
type WuxingXijiInfo struct {
	Xi  string `json:"xi"`
	Ji  string `json:"ji"`
	RiZhuQiangRuo string `json:"ri_zhu_qiang_ruo"`
	TiaoHouShen string `json:"tiao_hou_shen"`
}

// WuxingStrength represents the strength score and ranking of a single Five
// Element.
type WuxingStrength struct {
	WuxingFen map[string]float64 `json:"wuxing_fen"`
	Total     float64            `json:"total"`
}

// XiYongJiChou contains the favorable, unfavorable, hostile, and idle elements
// determined by the XiYong method.
type XiYongJiChou struct {
	Xi   string `json:"xi"`
	Yong string `json:"yong"`
	Ji   string `json:"ji"`
	Chou string `json:"chou"`
}

// GeJuType enumerates the Pattern (格局) types used in the GeJu XiYong method.
type GeJuType int

const (
	GeJuZhengGuan GeJuType = iota
	GeJuQiSha
	GeJuZhengCai
	GeJuPianCai
	GeJuZhengYin
	GeJuPianYin
	GeJuShiShen
	GeJuShangGuan
	GeJuUnknown
)

// GeJuInfo contains the Pattern (格局) analysis result.
type GeJuInfo struct {
	Type      GeJuType `json:"type"`
	Name      string   `json:"name"`
	YongShen  string   `json:"yong_shen"`
	XiShen    string   `json:"xi_shen"`
	JiShen    string   `json:"ji_shen"`
	ChouShen  string   `json:"chou_shen"`
	Analysis  string   `json:"analysis"`
}

// FateError represents an error from the fate analysis package.
type FateError struct {
	Code    int
	Message string
}

func (e FateError) Error() string {
	return e.Message
}

const (
	ErrCodeInputInvalid = iota + 1
	ErrCodeDateRange
	ErrCodeCalculateBazi
	ErrCodeCalculateWuxing
)

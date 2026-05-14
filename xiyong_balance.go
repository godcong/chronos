package chronos

type XiYongMethod int

const (
	XiYongMethodBalance XiYongMethod = iota
	XiYongMethodGeJu
)

type XiYongJiChou struct {
	YongWuxing string   `json:"yong_wuxing"`
	XiWuxing   []string `json:"xi_wuxing"`
	JiWuxing   []string `json:"ji_wuxing"`
	ChouWuxing []string `json:"chou_wuxing"`
	XianWuxing []string `json:"xian_wuxing"`
}

func balanceXiYongJi(riZhuGan string, qiangRuo string, strengths map[string]*WuxingStrength, tiaoHou []string) *XiYongJiChou {
	riZhuWuxing := getWuxingOfTianGan(riZhuGan)
	result := &XiYongJiChou{}

	switch qiangRuo {
	case "强":
		woKe := keWoMap[riZhuWuxing]
		woSheng := woShengMap[riZhuWuxing]
		keWo := shengWoMap[riZhuWuxing]

		result.YongWuxing = woKe
		result.XiWuxing = []string{woSheng, keWo}
		result.JiWuxing = []string{riZhuWuxing, shengWoMap[riZhuWuxing]}
		result.ChouWuxing = []string{shengWoMap[shengWoMap[riZhuWuxing]]}
		result.XianWuxing = findXianWuxing(result)

	case "弱":
		shengWo := shengWoMap[riZhuWuxing]
		result.YongWuxing = shengWo
		result.XiWuxing = []string{riZhuWuxing}
		result.JiWuxing = []string{keWoMap[riZhuWuxing], woShengMap[riZhuWuxing]}
		result.ChouWuxing = []string{shengWoMap[keWoMap[riZhuWuxing]]}
		result.XianWuxing = findXianWuxing(result)

	default:
		if len(tiaoHou) > 0 {
			result.YongWuxing = tiaoHou[0]
			result.XiWuxing = tiaoHou[1:]
		} else {
			weakest := findWeakestWuxing(strengths)
			result.YongWuxing = weakest
			result.XiWuxing = []string{shengWoMap[weakest]}
		}
		result.JiWuxing = []string{keWoMap[result.YongWuxing]}
		result.ChouWuxing = []string{shengWoMap[keWoMap[result.YongWuxing]]}
		result.XianWuxing = findXianWuxing(result)
	}

	return result
}

var (
	shengWoMap = map[string]string{
		"木": "水", "火": "木", "土": "火", "金": "土", "水": "金",
	}
	woShengMap = map[string]string{
		"木": "火", "火": "土", "土": "金", "金": "水", "水": "木",
	}
	keWoMap = map[string]string{
		"木": "金", "火": "水", "土": "木", "金": "火", "水": "土",
	}
	woKeMap = map[string]string{
		"木": "土", "火": "金", "土": "水", "金": "木", "水": "火",
	}
)

func findWeakestWuxing(strengths map[string]*WuxingStrength) string {
	minScore := 999.0
	weakest := ""
	for wx, s := range strengths {
		if s.Score < minScore {
			minScore = s.Score
			weakest = wx
		}
	}
	return weakest
}

func findXianWuxing(xyj *XiYongJiChou) []string {
	allWuxing := []string{"木", "火", "土", "金", "水"}
	used := map[string]bool{
		xyj.YongWuxing: true,
	}
	for _, wx := range xyj.XiWuxing {
		used[wx] = true
	}
	for _, wx := range xyj.JiWuxing {
		used[wx] = true
	}
	for _, wx := range xyj.ChouWuxing {
		used[wx] = true
	}
	var xian []string
	for _, wx := range allWuxing {
		if !used[wx] {
			xian = append(xian, wx)
		}
	}
	return xian
}

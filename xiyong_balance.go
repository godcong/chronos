package chronos

type XiYongMethod int

const (
	XiYongMethodBalance XiYongMethod = iota
	XiYongMethodGeJu
)

type XiYongJiChou struct {
	UsefulElement     string   `json:"useful_element"`
	FavorableElements []string `json:"favorable_elements"`
	UnfavorableElements []string `json:"unfavorable_elements"`
	HostileElements   []string `json:"hostile_elements"`
	IdleElements      []string `json:"idle_elements"`
}

func balanceXiYongJi(riZhuGan string, qiangRuo string, strengths map[string]*WuxingStrength, tiaoHou []string) *XiYongJiChou {
	riZhuWuxing := wuxingOfTianGan(riZhuGan)
	result := &XiYongJiChou{}

	switch qiangRuo {
	case "强":
		woKe := wuxingRelations.WoKe[riZhuWuxing]
		woSheng := wuxingRelations.WoSheng[riZhuWuxing]
		keWo := wuxingRelations.ShengWo[riZhuWuxing]

		result.UsefulElement = woKe
		result.FavorableElements = []string{woSheng, keWo}
		result.UnfavorableElements = []string{riZhuWuxing, wuxingRelations.ShengWo[riZhuWuxing]}
		result.HostileElements = []string{wuxingRelations.ShengWo[wuxingRelations.ShengWo[riZhuWuxing]]}
		result.IdleElements = findXianWuxing(result)

	case "弱":
		shengWo := wuxingRelations.ShengWo[riZhuWuxing]
		result.UsefulElement = shengWo
		result.FavorableElements = []string{riZhuWuxing}
		result.UnfavorableElements = []string{wuxingRelations.KeWo[riZhuWuxing], wuxingRelations.WoSheng[riZhuWuxing]}
		result.HostileElements = []string{wuxingRelations.ShengWo[wuxingRelations.KeWo[riZhuWuxing]]}
		result.IdleElements = findXianWuxing(result)

	default:
		if len(tiaoHou) > 0 {
			result.UsefulElement = tiaoHou[0]
			result.FavorableElements = tiaoHou[1:]
		} else {
			weakest := findWeakestWuxing(strengths)
			result.UsefulElement = weakest
			result.FavorableElements = []string{wuxingRelations.ShengWo[weakest]}
		}
		result.UnfavorableElements = []string{wuxingRelations.KeWo[result.UsefulElement]}
		result.HostileElements = []string{wuxingRelations.ShengWo[wuxingRelations.KeWo[result.UsefulElement]]}
		result.IdleElements = findXianWuxing(result)
	}

	return result
}

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
		xyj.UsefulElement: true,
	}
	for _, wx := range xyj.FavorableElements {
		used[wx] = true
	}
	for _, wx := range xyj.UnfavorableElements {
		used[wx] = true
	}
	for _, wx := range xyj.HostileElements {
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

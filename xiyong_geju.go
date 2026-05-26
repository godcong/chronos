package chronos

import (
	"fmt"
)

type GeJuType int

const (
	GeJuZhengGuan  GeJuType = iota + 1
	GeJuQiSha
	GeJuZhengYin
	GeJuPianYin
	GeJuShiShen
	GeJuShangGuan
	GeJuZhengCai
	GeJuPianCai
	GeJuJianLu
	GeJuYangRen
	GeJuSpecial
)

func (g GeJuType) String() string {
	names := map[GeJuType]string{
		GeJuZhengGuan: "正官格",
		GeJuQiSha:     "七杀格",
		GeJuZhengYin:  "正印格",
		GeJuPianYin:   "偏印格",
		GeJuShiShen:   "食神格",
		GeJuShangGuan: "伤官格",
		GeJuZhengCai:  "正财格",
		GeJuPianCai:   "偏财格",
		GeJuJianLu:    "建禄格",
		GeJuYangRen:   "阳刃格",
		GeJuSpecial:   "特殊格",
	}
	if name, ok := names[g]; ok {
		return name
	}
	return "未知格局"
}

type GeJuInfo struct {
	Type            GeJuType `json:"type"`
	Name            string   `json:"name"`
	MonthBranchStem string   `json:"month_branch_stem"`
	TenGod          string   `json:"ten_god"`
	UsefulElement   string   `json:"useful_element"`
}

func getShiShen(dayGan, targetGan string) string {
	if m, ok := tianGanShiShenMap[dayGan]; ok {
		if ss, ok := m[targetGan]; ok {
			return ss
		}
	}
	return ""
}

func determineGeJu(riZhuGan string, yueZhi string, siZhu [4]string) *GeJuInfo {
	cangGans, ok := diZhiHiddenStems[yueZhi]
	if !ok || len(cangGans) == 0 {
		return &GeJuInfo{Type: GeJuSpecial, Name: "特殊格"}
	}

	mainGan := cangGans[0]
	shiShen := getShiShen(riZhuGan, mainGan)

	geJuType := shiShenToGeJu(shiShen, riZhuGan, yueZhi)

	return &GeJuInfo{
		Type:            geJuType,
		Name:            geJuType.String(),
		MonthBranchStem: mainGan,
		TenGod:          shiShen,
	}
}

func shiShenToGeJu(shiShen string, riZhuGan, yueZhi string) GeJuType {
	switch shiShen {
	case "正官":
		return GeJuZhengGuan
	case "七杀":
		return GeJuQiSha
	case "正印":
		return GeJuZhengYin
	case "偏印":
		return GeJuPianYin
	case "食神":
		return GeJuShiShen
	case "伤官":
		return GeJuShangGuan
	case "正财":
		return GeJuZhengCai
	case "偏财":
		return GeJuPianCai
	case "比肩":
		return GeJuJianLu
	case "劫财":
		yangRenZhi := getYangRenZhi(riZhuGan)
		if yueZhi == yangRenZhi {
			return GeJuYangRen
		}
		return GeJuJianLu
	default:
		return GeJuSpecial
	}
}

func getYangRenZhi(riZhuGan string) string {
	m := map[string]string{
		"甲": "卯", "丙": "午", "戊": "午",
		"庚": "酉", "壬": "子",
	}
	if v, ok := m[riZhuGan]; ok {
		return v
	}
	return ""
}

func geJuXiYongJi(riZhuGan string, geJu *GeJuInfo, qiangRuo string, strengths map[string]*WuxingStrength, tiaoHou []string) *XiYongJiChou {
	riZhuWuxing := wuxingOfTianGan(riZhuGan)
	result := &XiYongJiChou{}

	switch geJu.Type {
	case GeJuZhengGuan:
		result = geJuZhengGuanYong(riZhuGan, riZhuWuxing, qiangRuo, strengths)
	case GeJuQiSha:
		result = geJuQiShaYong(riZhuGan, riZhuWuxing, qiangRuo, strengths)
	case GeJuZhengYin, GeJuPianYin:
		result = geJuYinXingYong(riZhuGan, riZhuWuxing, qiangRuo, strengths)
	case GeJuShiShen:
		result = geJuShiShenYong(riZhuGan, riZhuWuxing, qiangRuo, strengths)
	case GeJuShangGuan:
		result = geJuShangGuanYong(riZhuGan, riZhuWuxing, qiangRuo, strengths)
	case GeJuZhengCai, GeJuPianCai:
		result = geJuCaiXingYong(riZhuGan, riZhuWuxing, qiangRuo, strengths)
	case GeJuJianLu, GeJuYangRen:
		result = geJuJianLuYangRenYong(riZhuGan, riZhuWuxing, qiangRuo, strengths)
	default:
		result = balanceXiYongJi(riZhuGan, qiangRuo, strengths, tiaoHou)
	}

	return result
}

func geJuZhengGuanYong(riZhuGan, riZhuWuxing, qiangRuo string, strengths map[string]*WuxingStrength) *XiYongJiChou {
	guanWX := wuxingRelations.KeWo[riZhuWuxing]
	caiWX := wuxingRelations.WoKe[riZhuWuxing]
	yinWX := wuxingRelations.ShengWo[riZhuWuxing]

	if qiangRuo == "强" {
		return &XiYongJiChou{
			UsefulElement:     guanWX,
			FavorableElements: []string{caiWX, riZhuWuxing},
			UnfavorableElements: []string{yinWX},
			HostileElements:   []string{wuxingRelations.ShengWo[yinWX]},
		}
	}
	return &XiYongJiChou{
		UsefulElement:     yinWX,
		FavorableElements: []string{guanWX},
		UnfavorableElements: []string{caiWX, wuxingRelations.WoSheng[riZhuWuxing]},
		HostileElements:   []string{wuxingRelations.ShengWo[caiWX]},
	}
}

func geJuQiShaYong(riZhuGan, riZhuWuxing, qiangRuo string, strengths map[string]*WuxingStrength) *XiYongJiChou {
	shaWX := wuxingRelations.KeWo[riZhuWuxing]
	shiWX := wuxingRelations.WoSheng[riZhuWuxing]
	yinWX := wuxingRelations.ShengWo[riZhuWuxing]

	if qiangRuo == "强" {
		return &XiYongJiChou{
			UsefulElement:     shiWX,
			FavorableElements: []string{shaWX, yinWX},
			UnfavorableElements: []string{riZhuWuxing},
			HostileElements:   []string{wuxingRelations.ShengWo[riZhuWuxing]},
		}
	}
	return &XiYongJiChou{
		UsefulElement:     yinWX,
		FavorableElements: []string{shaWX},
		UnfavorableElements: []string{shiWX, wuxingRelations.WoKe[riZhuWuxing]},
		HostileElements:   []string{wuxingRelations.ShengWo[shiWX]},
	}
}

func geJuYinXingYong(riZhuGan, riZhuWuxing, qiangRuo string, strengths map[string]*WuxingStrength) *XiYongJiChou {
	yinWX := wuxingRelations.ShengWo[riZhuWuxing]
	guanWX := wuxingRelations.KeWo[riZhuWuxing]

	if qiangRuo == "强" {
		return &XiYongJiChou{
			UsefulElement:     guanWX,
			FavorableElements: []string{wuxingRelations.WoKe[riZhuWuxing]},
			UnfavorableElements: []string{yinWX, riZhuWuxing},
			HostileElements:   []string{wuxingRelations.ShengWo[yinWX]},
		}
	}
	return &XiYongJiChou{
		UsefulElement:     yinWX,
		FavorableElements: []string{guanWX},
		UnfavorableElements: []string{wuxingRelations.WoKe[riZhuWuxing], wuxingRelations.KeWo[riZhuWuxing]},
		HostileElements:   []string{wuxingRelations.ShengWo[wuxingRelations.KeWo[riZhuWuxing]]},
	}
}

func geJuShiShenYong(riZhuGan, riZhuWuxing, qiangRuo string, strengths map[string]*WuxingStrength) *XiYongJiChou {
	shiWX := wuxingRelations.WoSheng[riZhuWuxing]
	caiWX := wuxingRelations.WoKe[riZhuWuxing]

	if qiangRuo == "强" {
		return &XiYongJiChou{
			UsefulElement:     shiWX,
			FavorableElements: []string{caiWX},
			UnfavorableElements: []string{wuxingRelations.ShengWo[riZhuWuxing], riZhuWuxing},
			HostileElements:   []string{wuxingRelations.ShengWo[wuxingRelations.ShengWo[riZhuWuxing]]},
		}
	}
	return &XiYongJiChou{
		UsefulElement:     riZhuWuxing,
		FavorableElements: []string{wuxingRelations.ShengWo[riZhuWuxing]},
		UnfavorableElements: []string{shiWX, caiWX},
		HostileElements:   []string{wuxingRelations.ShengWo[shiWX]},
	}
}

func geJuShangGuanYong(riZhuGan, riZhuWuxing, qiangRuo string, strengths map[string]*WuxingStrength) *XiYongJiChou {
	shangWX := wuxingRelations.WoSheng[riZhuWuxing]
	yinWX := wuxingRelations.ShengWo[riZhuWuxing]
	caiWX := wuxingRelations.WoKe[riZhuWuxing]

	if qiangRuo == "强" {
		return &XiYongJiChou{
			UsefulElement:     yinWX,
			FavorableElements: []string{shangWX, caiWX},
			UnfavorableElements: []string{riZhuWuxing, wuxingRelations.KeWo[riZhuWuxing]},
			HostileElements:   []string{wuxingRelations.ShengWo[wuxingRelations.KeWo[riZhuWuxing]]},
		}
	}
	return &XiYongJiChou{
		UsefulElement:     yinWX,
		FavorableElements: []string{wuxingRelations.ShengWo[yinWX]},
		UnfavorableElements: []string{shangWX, wuxingRelations.KeWo[riZhuWuxing]},
		HostileElements:   []string{wuxingRelations.ShengWo[wuxingRelations.KeWo[riZhuWuxing]]},
	}
}

func geJuCaiXingYong(riZhuGan, riZhuWuxing, qiangRuo string, strengths map[string]*WuxingStrength) *XiYongJiChou {
	caiWX := wuxingRelations.WoKe[riZhuWuxing]
	shiWX := wuxingRelations.WoSheng[riZhuWuxing]
	guanWX := wuxingRelations.KeWo[riZhuWuxing]

	if qiangRuo == "强" {
		return &XiYongJiChou{
			UsefulElement:     caiWX,
			FavorableElements: []string{shiWX, guanWX},
			UnfavorableElements: []string{riZhuWuxing, wuxingRelations.ShengWo[riZhuWuxing]},
			HostileElements:   []string{wuxingRelations.ShengWo[wuxingRelations.ShengWo[riZhuWuxing]]},
		}
	}
	return &XiYongJiChou{
		UsefulElement:     wuxingRelations.ShengWo[riZhuWuxing],
		FavorableElements: []string{riZhuWuxing},
		UnfavorableElements: []string{caiWX, guanWX},
		HostileElements:   []string{wuxingRelations.ShengWo[caiWX]},
	}
}

func geJuJianLuYangRenYong(riZhuGan, riZhuWuxing, qiangRuo string, strengths map[string]*WuxingStrength) *XiYongJiChou {
	guanWX := wuxingRelations.KeWo[riZhuWuxing]
	caiWX := wuxingRelations.WoKe[riZhuWuxing]
	shiWX := wuxingRelations.WoSheng[riZhuWuxing]

	if qiangRuo == "强" {
		return &XiYongJiChou{
			UsefulElement:     guanWX,
			FavorableElements: []string{caiWX, shiWX},
			UnfavorableElements: []string{riZhuWuxing, wuxingRelations.ShengWo[riZhuWuxing]},
			HostileElements:   []string{wuxingRelations.ShengWo[wuxingRelations.ShengWo[riZhuWuxing]]},
		}
	}
	return &XiYongJiChou{
		UsefulElement:     wuxingRelations.ShengWo[riZhuWuxing],
		FavorableElements: []string{riZhuWuxing},
		UnfavorableElements: []string{guanWX, caiWX},
		HostileElements:   []string{wuxingRelations.ShengWo[guanWX]},
	}
}

func generateGeJuAnalysis(riZhuGan, qiangRuo string, geJu *GeJuInfo, xyj *XiYongJiChou) string {
	return fmt.Sprintf("日主%s，五行%s，格局%s（%s），%s格。用神为%s，喜神为%s，忌神为%s。",
		riZhuGan, wuxingOfTianGan(riZhuGan), geJu.Name, geJu.TenGod, qiangRuo,
		xyj.UsefulElement, joinStrings(xyj.FavorableElements, "、"), joinStrings(xyj.UnfavorableElements, "、"))
}

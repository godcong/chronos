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
	Type       GeJuType `json:"type"`
	Name       string   `json:"name"`
	YueZhiGan  string   `json:"yue_zhi_gan"`
	ShiShen    string   `json:"shi_shen"`
	YongShenWX string   `json:"yong_shen_wx"`
}

var diZhiCangGan = map[string][]string{
	"子": {"癸"},
	"丑": {"己", "癸", "辛"},
	"寅": {"甲", "丙", "戊"},
	"卯": {"乙"},
	"辰": {"戊", "乙", "癸"},
	"巳": {"丙", "庚", "戊"},
	"午": {"丁", "己"},
	"未": {"己", "丁", "乙"},
	"申": {"庚", "壬", "戊"},
	"酉": {"辛"},
	"戌": {"戊", "辛", "丁"},
	"亥": {"壬", "甲"},
}

var tianGanShiShenMap = map[string]map[string]string{
	"甲": {"甲": "比肩", "乙": "劫财", "丙": "食神", "丁": "伤官", "戊": "偏财", "己": "正财", "庚": "七杀", "辛": "正官", "壬": "偏印", "癸": "正印"},
	"乙": {"甲": "劫财", "乙": "比肩", "丙": "伤官", "丁": "食神", "戊": "正财", "己": "偏财", "庚": "正官", "辛": "七杀", "壬": "正印", "癸": "偏印"},
	"丙": {"甲": "偏印", "乙": "正印", "丙": "比肩", "丁": "劫财", "戊": "食神", "己": "伤官", "庚": "偏财", "辛": "正财", "壬": "七杀", "癸": "正官"},
	"丁": {"甲": "正印", "乙": "偏印", "丙": "劫财", "丁": "比肩", "戊": "伤官", "己": "食神", "庚": "正财", "辛": "偏财", "壬": "正官", "癸": "七杀"},
	"戊": {"甲": "七杀", "乙": "正官", "丙": "偏印", "丁": "正印", "戊": "比肩", "己": "劫财", "庚": "食神", "辛": "伤官", "壬": "偏财", "癸": "正财"},
	"己": {"甲": "正官", "乙": "七杀", "丙": "正印", "丁": "偏印", "戊": "劫财", "己": "比肩", "庚": "伤官", "辛": "食神", "壬": "正财", "癸": "偏财"},
	"庚": {"甲": "偏财", "乙": "正财", "丙": "七杀", "丁": "正官", "戊": "偏印", "己": "正印", "庚": "比肩", "辛": "劫财", "壬": "食神", "癸": "伤官"},
	"辛": {"甲": "正财", "乙": "偏财", "丙": "正官", "丁": "七杀", "戊": "正印", "己": "偏印", "庚": "劫财", "辛": "比肩", "壬": "伤官", "癸": "食神"},
	"壬": {"甲": "食神", "乙": "伤官", "丙": "偏财", "丁": "正财", "戊": "七杀", "己": "正官", "庚": "偏印", "辛": "正印", "壬": "比肩", "癸": "劫财"},
	"癸": {"甲": "伤官", "乙": "食神", "丙": "正财", "丁": "偏财", "戊": "正官", "己": "七杀", "庚": "正印", "辛": "偏印", "壬": "劫财", "癸": "比肩"},
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
	cangGans, ok := diZhiCangGan[yueZhi]
	if !ok || len(cangGans) == 0 {
		return &GeJuInfo{Type: GeJuSpecial, Name: "特殊格"}
	}

	mainGan := cangGans[0]
	shiShen := getShiShen(riZhuGan, mainGan)

	geJuType := shiShenToGeJu(shiShen, riZhuGan, yueZhi)

	return &GeJuInfo{
		Type:      geJuType,
		Name:      geJuType.String(),
		YueZhiGan: mainGan,
		ShiShen:   shiShen,
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
	riZhuWuxing := getWuxingOfTianGan(riZhuGan)
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
	guanWX := keWoMap[riZhuWuxing]
	caiWX := woKeMap[riZhuWuxing]
	yinWX := shengWoMap[riZhuWuxing]

	if qiangRuo == "强" {
		return &XiYongJiChou{
			YongWuxing: guanWX,
			XiWuxing:   []string{caiWX, riZhuWuxing},
			JiWuxing:   []string{yinWX},
			ChouWuxing: []string{shengWoMap[yinWX]},
		}
	}
	return &XiYongJiChou{
		YongWuxing: yinWX,
		XiWuxing:   []string{guanWX},
		JiWuxing:   []string{caiWX, woShengMap[riZhuWuxing]},
		ChouWuxing: []string{shengWoMap[caiWX]},
	}
}

func geJuQiShaYong(riZhuGan, riZhuWuxing, qiangRuo string, strengths map[string]*WuxingStrength) *XiYongJiChou {
	shaWX := keWoMap[riZhuWuxing]
	shiWX := woShengMap[riZhuWuxing]
	yinWX := shengWoMap[riZhuWuxing]

	if qiangRuo == "强" {
		return &XiYongJiChou{
			YongWuxing: shiWX,
			XiWuxing:   []string{shaWX, yinWX},
			JiWuxing:   []string{riZhuWuxing},
			ChouWuxing: []string{shengWoMap[riZhuWuxing]},
		}
	}
	return &XiYongJiChou{
		YongWuxing: yinWX,
		XiWuxing:   []string{shaWX},
		JiWuxing:   []string{shiWX, caiWX(riZhuWuxing)},
		ChouWuxing: []string{shengWoMap[shiWX]},
	}
}

func caiWX(riZhuWuxing string) string {
	return woKeMap[riZhuWuxing]
}

func geJuYinXingYong(riZhuGan, riZhuWuxing, qiangRuo string, strengths map[string]*WuxingStrength) *XiYongJiChou {
	yinWX := shengWoMap[riZhuWuxing]
	guanWX := keWoMap[riZhuWuxing]

	if qiangRuo == "强" {
		return &XiYongJiChou{
			YongWuxing: guanWX,
			XiWuxing:   []string{caiWX(riZhuWuxing)},
			JiWuxing:   []string{yinWX, riZhuWuxing},
			ChouWuxing: []string{shengWoMap[yinWX]},
		}
	}
	return &XiYongJiChou{
		YongWuxing: yinWX,
		XiWuxing:   []string{guanWX},
		JiWuxing:   []string{caiWX(riZhuWuxing), keWoMap[riZhuWuxing]},
		ChouWuxing: []string{shengWoMap[keWoMap[riZhuWuxing]]},
	}
}

func geJuShiShenYong(riZhuGan, riZhuWuxing, qiangRuo string, strengths map[string]*WuxingStrength) *XiYongJiChou {
	shiWX := woShengMap[riZhuWuxing]
	caiWX_ := woKeMap[riZhuWuxing]

	if qiangRuo == "强" {
		return &XiYongJiChou{
			YongWuxing: shiWX,
			XiWuxing:   []string{caiWX_},
			JiWuxing:   []string{shengWoMap[riZhuWuxing], riZhuWuxing},
			ChouWuxing: []string{shengWoMap[shengWoMap[riZhuWuxing]]},
		}
	}
	return &XiYongJiChou{
		YongWuxing: riZhuWuxing,
		XiWuxing:   []string{shengWoMap[riZhuWuxing]},
		JiWuxing:   []string{shiWX, caiWX_},
		ChouWuxing: []string{shengWoMap[shiWX]},
	}
}

func geJuShangGuanYong(riZhuGan, riZhuWuxing, qiangRuo string, strengths map[string]*WuxingStrength) *XiYongJiChou {
	shangWX := woShengMap[riZhuWuxing]
	yinWX := shengWoMap[riZhuWuxing]
	caiWX_ := woKeMap[riZhuWuxing]

	if qiangRuo == "强" {
		return &XiYongJiChou{
			YongWuxing: yinWX,
			XiWuxing:   []string{shangWX, caiWX_},
			JiWuxing:   []string{riZhuWuxing, keWoMap[riZhuWuxing]},
			ChouWuxing: []string{shengWoMap[keWoMap[riZhuWuxing]]},
		}
	}
	return &XiYongJiChou{
		YongWuxing: yinWX,
		XiWuxing:   []string{shengWoMap[yinWX]},
		JiWuxing:   []string{shangWX, keWoMap[riZhuWuxing]},
		ChouWuxing: []string{shengWoMap[keWoMap[riZhuWuxing]]},
	}
}

func geJuCaiXingYong(riZhuGan, riZhuWuxing, qiangRuo string, strengths map[string]*WuxingStrength) *XiYongJiChou {
	caiWX_ := woKeMap[riZhuWuxing]
	shiWX := woShengMap[riZhuWuxing]
	guanWX := keWoMap[riZhuWuxing]

	if qiangRuo == "强" {
		return &XiYongJiChou{
			YongWuxing: caiWX_,
			XiWuxing:   []string{shiWX, guanWX},
			JiWuxing:   []string{riZhuWuxing, shengWoMap[riZhuWuxing]},
			ChouWuxing: []string{shengWoMap[shengWoMap[riZhuWuxing]]},
		}
	}
	return &XiYongJiChou{
		YongWuxing: shengWoMap[riZhuWuxing],
		XiWuxing:   []string{riZhuWuxing},
		JiWuxing:   []string{caiWX_, guanWX},
		ChouWuxing: []string{shengWoMap[caiWX_]},
	}
}

func geJuJianLuYangRenYong(riZhuGan, riZhuWuxing, qiangRuo string, strengths map[string]*WuxingStrength) *XiYongJiChou {
	guanWX := keWoMap[riZhuWuxing]
	caiWX_ := woKeMap[riZhuWuxing]
	shiWX := woShengMap[riZhuWuxing]

	if qiangRuo == "强" {
		return &XiYongJiChou{
			YongWuxing: guanWX,
			XiWuxing:   []string{caiWX_, shiWX},
			JiWuxing:   []string{riZhuWuxing, shengWoMap[riZhuWuxing]},
			ChouWuxing: []string{shengWoMap[shengWoMap[riZhuWuxing]]},
		}
	}
	return &XiYongJiChou{
		YongWuxing: shengWoMap[riZhuWuxing],
		XiWuxing:   []string{riZhuWuxing},
		JiWuxing:   []string{guanWX, caiWX_},
		ChouWuxing: []string{shengWoMap[guanWX]},
	}
}

func generateGeJuAnalysis(riZhuGan, qiangRuo string, geJu *GeJuInfo, xyj *XiYongJiChou) string {
	return fmt.Sprintf("日主%s，五行%s，格局%s（%s），%s格。用神为%s，喜神为%s，忌神为%s。",
		riZhuGan, getWuxingOfTianGan(riZhuGan), geJu.Name, geJu.ShiShen, qiangRuo,
		xyj.YongWuxing, joinStrings(xyj.XiWuxing, "、"), joinStrings(xyj.JiWuxing, "、"))
}

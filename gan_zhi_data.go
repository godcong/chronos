package chronos

//var _TianGanDiZhiGanZhiTable [TianGanMax][DiZhiMax]GanZhi
var _TianGanDiZhiGanZhiTable = [TianGanMax][DiZhiMax]GanZhi{
	TianGanJia:  {GanZhiJiaZi, GanZhiMax, GanZhiJiaYin, GanZhiMax, GanZhiJiaChen, GanZhiMax, GanZhiJiaWu, GanZhiMax, GanZhiJiaShen, GanZhiMax, GanZhiJiaXu, GanZhiMax},
	TianGanYi:   {GanZhiMax, GanZhiYiChou, GanZhiMax, GanZhiYiMao, GanZhiMax, GanZhiYiSi, GanZhiMax, GanZhiYiWei, GanZhiMax, GanZhiYiYou, GanZhiMax, GanZhiYiHai},
	TianGanBing: {GanZhiBingZi, GanZhiMax, GanZhiBingYin, GanZhiMax, GanZhiBingChen, GanZhiMax, GanZhiBingWu, GanZhiMax, GanZhiBingShen, GanZhiMax, GanZhiBingXu, GanZhiMax},
	TianGanDing: {GanZhiMax, GanZhiDingChou, GanZhiMax, GanZhiDingMao, GanZhiMax, GanZhiDingSi, GanZhiMax, GanZhiDingWei, GanZhiMax, GanZhiDingYou, GanZhiMax, GanZhiDingHai},
	TianGanWu:   {GanZhiWuZi, GanZhiMax, GanZhiWuYin, GanZhiMax, GanZhiWuChen, GanZhiMax, GanZhiWuWu, GanZhiMax, GanZhiWuShen, GanZhiMax, GanZhiWuXu, GanZhiMax},
	TianGanJi:   {GanZhiMax, GanZhiJiChou, GanZhiMax, GanZhiJiMao, GanZhiMax, GanZhiJiSi, GanZhiMax, GanZhiJiWei, GanZhiMax, GanZhiJiYou, GanZhiMax, GanZhiJiHai},
	TianGanGeng: {GanZhiGengZi, GanZhiMax, GanZhiGengYin, GanZhiMax, GanZhiGengChen, GanZhiMax, GanZhiGengWu, GanZhiMax, GanZhiGengShen, GanZhiMax, GanZhiGengXu, GanZhiMax},
	TianGanXin:  {GanZhiMax, GanZhiXinChou, GanZhiMax, GanZhiXinMao, GanZhiMax, GanZhiXinSi, GanZhiMax, GanZhiXinWei, GanZhiMax, GanZhiXinYou, GanZhiMax, GanZhiXinHai},
	TianGanRen:  {GanZhiRenZi, GanZhiMax, GanZhiRenYin, GanZhiMax, GanZhiRenChen, GanZhiMax, GanZhiRenWu, GanZhiMax, GanZhiRenShen, GanZhiMax, GanZhiRenXu, GanZhiMax},
	TianGanGui:  {GanZhiMax, GanZhiGuiChou, GanZhiMax, GanZhiGuiMao, GanZhiMax, GanZhiGuiSi, GanZhiMax, GanZhiGuiWei, GanZhiMax, GanZhiGuiYou, GanZhiMax, GanZhiGuiHai},
}

func init() {
	//var ganzhi GanZhi
	//var err error
	//for gi := range _TianGanDiZhiGanZhiTable {
	//	for zi := range _TianGanDiZhiGanZhiTable[gi] {
	//		_TianGanDiZhiGanZhiTable[gi][zi] = GanZhiMax
	//		ganzhi, err = ParseGanZhi(ganZhiStr(TianGan(gi), DiZhi(zi)))
	//		if err == nil {
	//			_TianGanDiZhiGanZhiTable[gi][zi] = ganzhi
	//		}
	//	}
	//}
}

//func PrintGanZhiTable() ([]byte, error) {
//	return json.Marshal(_TianGanDiZhiGanZhiTable)
//}

package chronos

var _TianGanTable = []string{
	`甲`, `乙`, `丙`, `丁`, `戊`, `己`, `庚`, `辛`, `壬`, `癸`,
}

var _DiZhiTable = []string{
	`子`, `丑`, `寅`, `卯`, `辰`, `巳`, `午`, `未`, `申`, `酉`, `戌`, `亥`,
}

var _GanZhiTable = []string{
	`甲子`, `乙丑`, `丙寅`, `丁卯`, `戊辰`, `己巳`, `庚午`, `辛未`, `壬申`, `癸酉`, `甲戌`, `乙亥`, //甲或己日
	`丙子`, `丁丑`, `戊寅`, `己卯`, `庚辰`, `辛巳`, `壬午`, `癸未`, `甲申`, `乙酉`, `丙戌`, `丁亥`, //乙或庚日
	`戊子`, `己丑`, `庚寅`, `辛卯`, `壬辰`, `癸巳`, `甲午`, `乙未`, `丙申`, `丁酉`, `戊戌`, `己亥`, //丙或辛日
	`庚子`, `辛丑`, `壬寅`, `癸卯`, `甲辰`, `乙巳`, `丙午`, `丁未`, `戊申`, `己酉`, `庚戌`, `辛亥`, //丁或壬日
	`壬子`, `癸丑`, `甲寅`, `乙卯`, `丙辰`, `丁巳`, `戊午`, `己未`, `庚申`, `辛酉`, `壬戌`, `癸亥`, //戊或癸日
}

// PillarHour is an alias name of ShiZhu
var PillarHour = ShiZhu
var PillarDay = RiZhu
var PillarMonty = YueZhu
var PillarYear = NianZhu

//GetGanZhi 取得干支
func GetGanZhi(y int) string {
	return _TianGanTable[y%10] + _DiZhiTable[y%12]
}

//ShiZhu 获取时柱
//　	子 　　丑 　　寅 　　卯 　　辰 　　己
//　　　23-01：01-03：03-05 :05-07：07-09：09-11
//　　　午 　　未 　　申 　　酉 　　戊 　　亥
//　　　11-13：13-15：15-17：17-19：19-21：21-23
//`甲子`, `乙丑`, `丙寅`, `丁卯`, `戊辰`, `己巳`, `庚午`, `辛未`, `壬申`, `癸酉`, `甲戌`, `乙亥`, //甲或己日
//`丙子`, `丁丑`, `戊寅`, `己卯`, `庚辰`, `辛巳`, `壬午`, `癸未`, `甲申`, `乙酉`, `丙戌`, `丁亥`, //乙或庚日
//`戊子`, `己丑`, `庚寅`, `辛卯`, `壬辰`, `癸巳`, `甲午`, `乙未`, `丙申`, `丁酉`, `戊戌`, `己亥`, //丙或辛日
//`庚子`, `辛丑`, `壬寅`, `癸卯`, `甲辰`, `乙巳`, `丙午`, `丁未`, `戊申`, `己酉`, `庚戌`, `辛亥`, //丁或壬日
//`壬子`, `癸丑`, `甲寅`, `乙卯`, `丙辰`, `丁巳`, `戊午`, `己未`, `庚申`, `辛酉`, `壬戌`, `癸亥`, //戊或癸日
func ShiZhu(y, m, d, h int) string {
	i := stemBranchIndex(y, m, d) % 5 * 12
	idx := (h + 1) / 2 % 12
	return _GanZhiTable[fixDayNext(i, idx, h)]
}

// RiZhu 获取日柱
func RiZhu(y, m, d int) string {
	return _GanZhiTable[stemBranchIndex(y, m, d)]
}

//YueZhu 获取月柱
func YueZhu(y, m, d int) string {
	//月柱 1900年1月小寒以前为 丙子月(60进制12)
	fir := GetTermInfo(y, m*2-1) //返回当月「节」为几日开始

	//依据12节气修正干支月
	var sb = GetGanZhi(fixSuffix(y)*12 + m + 11)
	if d >= fir {
		sb = GetGanZhi(fixSuffix(y)*12 + m + 12)
	}
	return sb
}

//NianZhu 获取年柱
func NianZhu(y int) string {
	num := y - 4
	return _TianGanTable[num%10] + _DiZhiTable[num%12]
}

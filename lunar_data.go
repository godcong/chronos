package lunar

import (
	"log"
	"strconv"
)

var heavenlyStem = []string{
	`甲`, `乙`, `丙`, `丁`, `戊`, `己`, `庚`, `辛`, `壬`, `癸`,
}

var earthyBranch = []string{
	`子`, `丑`, `寅`, `卯`, `辰`, `巳`, `午`, `未`, `申`, `酉`, `戌`, `亥`,
}

var hourStemBranch = []string{
	`甲子`, `乙丑`, `丙寅`, `丁卯`, `戊辰`, `己巳`, `庚午`, `辛未`, `壬申`, `癸酉`, `甲戌`, `乙亥`, //甲或己日
	`丙子`, `丁丑`, `戊寅`, `己卯`, `庚辰`, `辛巳`, `壬午`, `癸未`, `甲申`, `乙酉`, `丙戌`, `丁亥`, //乙或庚日
	`戊子`, `己丑`, `庚寅`, `辛卯`, `壬辰`, `癸巳`, `甲午`, `乙未`, `丙申`, `丁酉`, `戊戌`, `己亥`, //丙或辛日
	`庚子`, `辛丑`, `壬寅`, `癸卯`, `甲辰`, `乙巳`, `丙午`, `丁未`, `戊申`, `己酉`, `庚戌`, `辛亥`, //丁或壬日
	`壬子`, `癸丑`, `甲寅`, `乙卯`, `丙辰`, `丁巳`, `戊午`, `己未`, `庚申`, `辛酉`, `壬戌`, `癸亥`, //戊或癸日
}

var yearTable = [][]int{
	//19
	{50, 31}, {60, 24}, {70, 16}, {80, 9}, {90, 1},
	{51, 36}, {61, 29}, {71, 21}, {81, 14}, {91, 6},
	{52, 42}, {62, 34}, {72, 27}, {82, 19}, {92, 12},
	{53, 47}, {63, 39}, {73, 32}, {83, 24}, {93, 17},
	{54, 52}, {64, 45}, {74, 37}, {84, 30}, {94, 22},
	{55, 57}, {65, 50}, {75, 42}, {85, 35}, {95, 27},
	{56, 3}, {66, 55}, {76, 48}, {86, 40}, {96, 33},
	{57, 8}, {67, 0}, {77, 53}, {87, 45}, {97, 38},
	{58, 13}, {68, 6}, {78, 58}, {88, 51}, {98, 43},
	{59, 18}, {69, 11}, {79, 3}, {89, 56}, {99, 48},
	//20
	{0, 54}, {10, 46}, {20, 39}, {30, 31}, {40, 24},
	{1, 59}, {11, 51}, {21, 44}, {31, 36}, {41, 29},
	{2, 4}, {12, 57}, {22, 49}, {32, 42}, {42, 34},
	{3, 9}, {13, 2}, {23, 54}, {33, 47}, {43, 39},
	{4, 15}, {14, 7}, {24, 0}, {34, 52}, {44, 45},
	{5, 20}, {15, 12}, {25, 5}, {35, 57}, {45, 50},
	{6, 25}, {16, 18}, {26, 10}, {36, 3}, {46, 55},
	{7, 30}, {17, 23}, {27, 15}, {37, 8}, {47, 0},
	{8, 36}, {18, 28}, {28, 21}, {38, 13}, {48, 6},
	{9, 41}, {19, 33}, {29, 26}, {39, 18}, {49, 11},
}

var lunarInfoList = []int{
	0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2, //1900-1909
	0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977, //1910-1919
	0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970, //1920-1929
	0x06566, 0x0d4a0, 0x0ea50, 0x06e95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950, //1930-1939
	0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557, //1940-1949
	0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5b0, 0x14573, 0x052b0, 0x0a9a8, 0x0e950, 0x06aa0, //1950-1959
	0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0, //1960-1969
	0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b6a0, 0x195a6, //1970-1979
	0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570, //1980-1989
	0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x055c0, 0x0ab60, 0x096d5, 0x092e0, //1990-1999
	0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5, //2000-2009
	0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930, //2010-2019
	0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530, //2020-2029
	0x05aa0, 0x076a3, 0x096d0, 0x04afb, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45, //2030-2039
	0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0, //2040-2049
	0x14b63, 0x09370, 0x049f8, 0x04970, 0x064b0, 0x168a6, 0x0ea50, 0x06b20, 0x1a6c4, 0x0aae0, //2050-2059
	0x0a2e0, 0x0d2e3, 0x0c960, 0x0d557, 0x0d4a0, 0x0da50, 0x05d55, 0x056a0, 0x0a6d0, 0x055d4, //2060-2069
	0x052d0, 0x0a9b8, 0x0a950, 0x0b4a0, 0x0b6a6, 0x0ad50, 0x055a0, 0x0aba4, 0x0a5b0, 0x052b0, //2070-2079
	0x0b273, 0x06930, 0x07337, 0x06aa0, 0x0ad50, 0x14b55, 0x04b60, 0x0a570, 0x054e4, 0x0d160, //2080-2089
	0x0e968, 0x0d520, 0x0daa0, 0x16aa6, 0x056d0, 0x04ae0, 0x0a9d4, 0x0a2d0, 0x0d150, 0x0f252, //2090-2099
	0x0d520, //2100
}

var termInfoList = []string{
	`9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`, `97bcf97c3598082c95f8c965cc920f`, `97bd0b06bdb0722c965ce1cfcc920f`, `b027097bd097c36b0b6fc9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`, `97bcf97c359801ec95f8c965cc920f`, `97bd0b06bdb0722c965ce1cfcc920f`, `b027097bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`,
	`97bcf97c359801ec95f8c965cc920f`, `97bd0b06bdb0722c965ce1cfcc920f`, `b027097bd097c36b0b6fc9274c91aa`, `9778397bd19801ec9210c965cc920e`, `97b6b97bd19801ec95f8c965cc920f`,
	`97bd09801d98082c95f8e1cfcc920f`, `97bd097bd097c36b0b6fc9210c8dc2`, `9778397bd197c36c9210c9274c91aa`, `97b6b97bd19801ec95f8c965cc920e`, `97bd09801d98082c95f8e1cfcc920f`,
	`97bd097bd097c36b0b6fc9210c8dc2`, `9778397bd097c36c9210c9274c91aa`, `97b6b97bd19801ec95f8c965cc920e`, `97bcf97c3598082c95f8e1cfcc920f`, `97bd097bd097c36b0b6fc9210c8dc2`,
	`9778397bd097c36c9210c9274c91aa`, `97b6b97bd19801ec9210c965cc920e`, `97bcf97c3598082c95f8c965cc920f`, `97bd097bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`, `97bcf97c3598082c95f8c965cc920f`, `97bd097bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`,
	`97bcf97c359801ec95f8c965cc920f`, `97bd097bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`, `97bcf97c359801ec95f8c965cc920f`,
	`97bd097bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`, `97bcf97c359801ec95f8c965cc920f`, `97bd097bd07f595b0b6fc920fb0722`,
	`9778397bd097c36b0b6fc9210c8dc2`, `9778397bd19801ec9210c9274c920e`, `97b6b97bd19801ec95f8c965cc920f`, `97bd07f5307f595b0b0bc920fb0722`, `7f0e397bd097c36b0b6fc9210c8dc2`,
	`9778397bd097c36c9210c9274c920e`, `97b6b97bd19801ec95f8c965cc920f`, `97bd07f5307f595b0b0bc920fb0722`, `7f0e397bd097c36b0b6fc9210c8dc2`, `9778397bd097c36c9210c9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`, `97bd07f1487f595b0b0bc920fb0722`, `7f0e397bd097c36b0b6fc9210c8dc2`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`,
	`97bcf7f1487f595b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`, `97bcf7f1487f595b0b0bb0b6fb0722`,
	`7f0e397bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`, `97bcf7f1487f531b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc920fb0722`,
	`9778397bd097c36b0b6fc9274c91aa`, `97b6b97bd19801ec9210c965cc920e`, `97bcf7f1487f531b0b0bb0b6fb0722`, `7f0e397bd07f595b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`,
	`97b6b97bd19801ec9210c9274c920e`, `97bcf7f0e47f531b0b0bb0b6fb0722`, `7f0e397bd07f595b0b0bc920fb0722`, `9778397bd097c36b0b6fc9210c91aa`, `97b6b97bd197c36c9210c9274c920e`,
	`97bcf7f0e47f531b0b0bb0b6fb0722`, `7f0e397bd07f595b0b0bc920fb0722`, `9778397bd097c36b0b6fc9210c8dc2`, `9778397bd097c36c9210c9274c920e`, `97b6b7f0e47f531b0723b0b6fb0722`,
	`7f0e37f5307f595b0b0bc920fb0722`, `7f0e397bd097c36b0b6fc9210c8dc2`, `9778397bd097c36b0b70c9274c91aa`, `97b6b7f0e47f531b0723b0b6fb0721`, `7f0e37f1487f595b0b0bb0b6fb0722`,
	`7f0e397bd097c35b0b6fc9210c8dc2`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f595b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc920fb0722`,
	`9778397bd097c36b0b6fc9274c91aa`, `97b6b7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`,
	`97b6b7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b7f0e47f531b0723b0b6fb0721`,
	`7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e397bd07f595b0b0bc920fb0722`, `9778397bd097c36b0b6fc9274c91aa`, `97b6b7f0e47f531b0723b0787b0721`, `7f0e27f0e47f531b0b0bb0b6fb0722`,
	`7f0e397bd07f595b0b0bc920fb0722`, `9778397bd097c36b0b6fc9210c91aa`, `97b6b7f0e47f149b0723b0787b0721`, `7f0e27f0e47f531b0723b0b6fb0722`, `7f0e397bd07f595b0b0bc920fb0722`,
	`9778397bd097c36b0b6fc9210c8dc2`, `977837f0e37f149b0723b0787b0721`, `7f07e7f0e47f531b0723b0b6fb0722`, `7f0e37f5307f595b0b0bc920fb0722`, `7f0e397bd097c35b0b6fc9210c8dc2`,
	`977837f0e37f14998082b0787b0721`, `7f07e7f0e47f531b0723b0b6fb0721`, `7f0e37f1487f595b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc9210c8dc2`, `977837f0e37f14998082b0787b06bd`,
	`7f07e7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc920fb0722`, `977837f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`,
	`7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e397bd097c35b0b6fc920fb0722`, `977837f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f531b0b0bb0b6fb0722`,
	`7f0e397bd07f595b0b0bc920fb0722`, `977837f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e397bd07f595b0b0bc920fb0722`,
	`977837f0e37f14998082b0787b06bd`, `7f07e7f0e47f149b0723b0787b0721`, `7f0e27f0e47f531b0b0bb0b6fb0722`, `7f0e397bd07f595b0b0bc920fb0722`, `977837f0e37f14998082b0723b06bd`,
	`7f07e7f0e37f149b0723b0787b0721`, `7f0e27f0e47f531b0723b0b6fb0722`, `7f0e397bd07f595b0b0bc920fb0722`, `977837f0e37f14898082b0723b02d5`, `7ec967f0e37f14998082b0787b0721`,
	`7f07e7f0e47f531b0723b0b6fb0722`, `7f0e37f1487f595b0b0bb0b6fb0722`, `7f0e37f0e37f14898082b0723b02d5`, `7ec967f0e37f14998082b0787b0721`, `7f07e7f0e47f531b0723b0b6fb0722`,
	`7f0e37f1487f531b0b0bb0b6fb0722`, `7f0e37f0e37f14898082b0723b02d5`, `7ec967f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`, `7f0e37f1487f531b0b0bb0b6fb0722`,
	`7f0e37f0e37f14898082b072297c35`, `7ec967f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e37f0e37f14898082b072297c35`,
	`7ec967f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`, `7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e37f0e366aa89801eb072297c35`, `7ec967f0e37f14998082b0787b06bd`,
	`7f07e7f0e47f149b0723b0787b0721`, `7f0e27f1487f531b0b0bb0b6fb0722`, `7f0e37f0e366aa89801eb072297c35`, `7ec967f0e37f14998082b0723b06bd`, `7f07e7f0e47f149b0723b0787b0721`,
	`7f0e27f0e47f531b0723b0b6fb0722`, `7f0e37f0e366aa89801eb072297c35`, `7ec967f0e37f14998082b0723b06bd`, `7f07e7f0e37f14998083b0787b0721`, `7f0e27f0e47f531b0723b0b6fb0722`,
	`7f0e37f0e366aa89801eb072297c35`, `7ec967f0e37f14898082b0723b02d5`, `7f07e7f0e37f14998082b0787b0721`, `7f07e7f0e47f531b0723b0b6fb0722`, `7f0e36665b66aa89801e9808297c35`,
	`665f67f0e37f14898082b0723b02d5`, `7ec967f0e37f14998082b0787b0721`, `7f07e7f0e47f531b0723b0b6fb0722`, `7f0e36665b66a449801e9808297c35`, `665f67f0e37f14898082b0723b02d5`,
	`7ec967f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`, `7f0e36665b66a449801e9808297c35`, `665f67f0e37f14898082b072297c35`, `7ec967f0e37f14998082b0787b06bd`,
	`7f07e7f0e47f531b0723b0b6fb0721`, `7f0e26665b66a449801e9808297c35`, `665f67f0e37f1489801eb072297c35`, `7ec967f0e37f14998082b0787b06bd`, `7f07e7f0e47f531b0723b0b6fb0721`,
	`7f0e27f1487f531b0b0bb0b6fb0722`, //2100
}

var number = []string{`一`, `二`, `三`, `四`, `五`, `六`, `七`, `八`, `九`, `十`, `十一`, `十二`}
var ten = []string{`初`, `十`, `廿`, `卅`}
var chineseNumber = []string{`正`, `二`, `三`, `四`, `五`, `六`, `七`, `八`, `九`, `十`, `十一`, `腊`}

func fixSuffix(y int) int {
	return y - 1900
}

func GetLunarInfo(y int) int {
	y = fixSuffix(y)
	if y < 0 || y > len(lunarInfoList) {
		return 0
	}
	return lunarInfoList[y]
}

func GetTermInfo(y, n int) int {
	y = fixSuffix(y)
	if y < 0 || y > len(termInfoList) {
		return -1
	}
	if n < 1 || n > 24 {
		return -1
	}

	t := termInfoList[y]
	var day []string
	for i := 0; i < 30; i += 5 {
		i, _ := strconv.ParseInt("0x"+t[i:i+5], 0, 64)
		a := strconv.Itoa(int(i))
		day = append(day, a[0:1], a[1:3], a[3:4], a[4:6])

	}

	i, _ := strconv.Atoi(day[n-1])
	return i
}

func GetChineseMonth(m int) string {
	if m > 12 || m < 1 {
		return "?月"
	}
	return chineseNumber[m-1] + "月" //加上月字
}

func GetChineseDay(d int) string {
	if d < 0 || d > 31 {
		return "?日"
	}
	var s string
	switch d {
	case 10:
		s = `初十`
	case 20:
		s = `二十`
	case 30:
		s = `三十`
	default:
		n := (d - 1) % 10
		s = ten[d/10] + number[n]
	}
	return s + "日"

}

//GetStemBranch
func GetStemBranch(y int) string {
	return heavenlyStem[y%10] + earthyBranch[y%12]
}

//StemBranchHour
//　子 　　丑 　　寅 　　卯 　　辰 　　己
//　　　23-01：01-03：03-05 :05-07：07-09：09-11
//　　　　午 　　未 　　申 　　酉 　　戊 　　亥
//　　　11-13：13-15：15-17：17-19：19-21：21-23
//`甲子`, `乙丑`, `丙寅`, `丁卯`, `戊辰`, `己巳`, `庚午`, `辛未`, `壬申`, `癸酉`, `甲戌`, `乙亥`, //甲或己日
//`丙子`, `丁丑`, `戊寅`, `己卯`, `庚辰`, `辛巳`, `壬午`, `癸未`, `甲申`, `乙酉`, `丙戌`, `丁亥`, //乙或庚日
//`戊子`, `己丑`, `庚寅`, `辛卯`, `壬辰`, `癸巳`, `甲午`, `乙未`, `丙申`, `丁酉`, `戊戌`, `己亥`, //丙或辛日
//`庚子`, `辛丑`, `壬寅`, `癸卯`, `甲辰`, `乙巳`, `丙午`, `丁未`, `戊申`, `己酉`, `庚戌`, `辛亥`, //丁或壬日
//`壬子`, `癸丑`, `甲寅`, `乙卯`, `丙辰`, `丁巳`, `戊午`, `己未`, `庚申`, `辛酉`, `壬戌`, `癸亥`, //戊或癸日
func StemBranchHour(y, m, d, h int) string {
	g, _ := calculateStemBranch(y, m, d)
	h = h / 2 % 12
	g = (g + h - 8) % 10
	//g = (g%10)%5*12 + h

	//g = (g + 2 + h) % 10 * 12
	log.Print("g:", g%10%5)
	//log.Print(hourStemBranch[g])
	return ""
	return heavenlyStem[g%10] + earthyBranch[h]
}
func calcHourStemBranch(g, h int) string {
	//hourStemBranch
	return ""
}

// StemBranchDay
//　  G = 4C + [C / 4] + 5y + [y / 4] + [3 * (M + 1) / 5] + d - 3
//　　Z = 8C + [C / 4] + 5y + [y / 4] + [3 * (M + 1) / 5] + d + 7 + i
//　　其中 C 是世纪数减一，y 是年份后两位，M 是月份，d 是日数。1月和2月按上一年的13月和14月来算。奇数月i=0，偶数月i=6。G 除以10的余数是天干，Z 除以12的余数是地支。
//　　计算时带[ ]的数表示取整。
func StemBranchDay(y, m, d int) string {
	g, z := calculateStemBranch(y, m, d)
	return heavenlyStem[g%10] + earthyBranch[z%12]
}

func calculateStemBranch(y, m, d int) (int, int) {
	c := y / 100
	y = y % 100
	i := 0
	if m%2 == 0 {
		i = 6
	}
	g := 4*c + c/4 + 5*y + y/4 + 3*(m+1)/5 + d - 2
	z := 8*c + c/4 + 5*y + y/4 + 3*(m+1)/5 + d + 8 + i
	return g, z
}

//StemBranchMonth
func StemBranchMonth(y, m, d int) string {
	//月柱 1900年1月小寒以前为 丙子月(60进制12)
	fir := GetTermInfo(y, m*2-1) //返回当月「节」为几日开始
	//sec := GetTermInfo(y, m*2)   //返回当月「节」为几日开始

	//依据12节气修正干支月
	var sb = GetStemBranch((y-1900)*12 + m + 11)
	if d >= fir {
		sb = GetStemBranch((y-1900)*12 + m + 12)
	}
	return sb
}

//StemBranchYear
func StemBranchYear(y int) string {
	num := y - 4
	return heavenlyStem[num%10] + earthyBranch[num%12]
}

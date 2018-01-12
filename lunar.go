package lunar

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Lunar struct {
	time      time.Time
	Year      int
	Month     int
	Day       int
	leapMonth int
	isLeap    bool
}

func GetZodiac(time time.Time) string {
	return Zodiac[(time.Year()-4)%12]
}

func yearDay(y int) int {
	i, sum := 348, 348
	for i = 0x8000; i > 0x8; i >>= 1 {
		if (lunarInfo[y-1900] & i) != 0 {
			sum++
		}
	}
	return sum + leapDay(y)
}

func leapDay(y int) int {
	if leapMonth(y) != 0 {
		if (lunarInfo[y-1900] & 0x10000) != 0 {
			return 30
		}
		return 29
	}
	return 0
}

func leapMonth(y int) int {
	return lunarInfo[y-1900] & 0xf
}

func monthDays(y int, m int) int {
	//月份参数从1至12，参数错误返回-1
	if m > 12 || m < 1 {
		return -1
	}
	if lunarInfo[y-1900]&(0x10000>>uint32(m)) != 0 {
		return 30
	}
	return 29
}

func solarDays(y, m int) int {
	//若参数错误 返回-1
	if m > 12 || m < 1 {
		return -1
	}
	var idx = m - 1
	if idx == 1 { //2月份的闰平规律测算后确认返回28或29
		if (y%4 == 0) && (y%100 != 0) || (y%400 == 0) {
			return 29
		}
		return 28
	} else {
		return monthDay[idx]
	}
}

//　子 　　丑 　　寅 　　卯 　　辰 　　己
//　　　23-01：01-03：03-05 :05-07：07-09：09-11
//　　　　午 　　未 　　申 　　酉 　　戊 　　亥
//　　　11-13：13-15：15-17：17-19：19-21：21-23

func StemBranchHour(h int) string {
	h = h % 12
	return EarthyBranch[h] + "时"
}

//　  G = 4C + [C / 4] + 5y + [y / 4] + [3 * (M + 1) / 5] + d - 3
//　　Z = 8C + [C / 4] + 5y + [y / 4] + [3 * (M + 1) / 5] + d + 7 + i
//　　其中 C 是世纪数减一，y 是年份后两位，M 是月份，d 是日数。1月和2月按上一年的13月和14月来算。奇数月i=0，偶数月i=6。G 除以10的余数是天干，Z 除以12的余数是地支。
//　　计算时带[ ]的数表示取整。
func StemBranchDay(y, m, d int) string {
	c := y / 100
	y = y % 100
	i := 0
	if m%2 == 0 {
		i = 6
	}
	g := 4*c + c/4 + 5*y + y/4 + 3*(m+1)/5 + d - 4
	z := 8*c + c/4 + 5*y + y/4 + 3*(m+1)/5 + d + 6 + i
	return HeavenlyStem[g%10] + EarthyBranch[z%12]
}

func StemBranchMonth(y, m int) string {
	GetChineseMonth(m)
	return ChinesMonth[(y+1)%5*12+m-1]
}

func StemBranchYear(y int) string {
	num := y - 4
	return HeavenlyStem[num%10] + EarthyBranch[num%12]
}

func GetAstro(m, d int) string {
	arr := []int{20, 19, 21, 21, 21, 22, 23, 23, 23, 23, 22, 22}
	idx := d < arr[m-1]
	index := m * 2
	if idx {
		index = m*2 - 2
	}
	return Astro[index] + "座"
}

func GetTerm(y, n int) int {
	if y < 1900 || y > 2100 {
		return -1
	}
	if n < 1 || n > 24 {
		return -1
	}

	t := termInfo[y-1900]
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
		return ""
	}
	return ChineseNumber[m-1] + "月" //加上月字
}

func GetChineseDay(d int) string {
	var s string
	switch d {
	case 10:
		s = `初十`
	case 20:
		s = `二十`
		break
	case 30:
		s = `三十`
		break
	default:
		n := d % 10
		if n == 0 {
			n = 10
		}
		log.Println(d, n)
		s = Ten[d/10] + Number[n-1]
	}
	return s + "日"

}

func lunarYear(offset int) int {
	day := 0
	i := 0
	//求当年农历年天数
	for i = YEAR_MIN; i <= YEAR_MAX; i++ {
		day = yearDay(i)
		if offset-day < 1 {
			break
		}
		offset -= day

	}
	return i
}

func CalculateLunar(date string) Lunar {
	lunar := Lunar{
		isLeap: false,
	}
	loc, _ := time.LoadLocation("Local")
	i := 0
	day := 0
	isLeapYear := false
	input, err := time.ParseInLocation(DATE_FORMAT, date, loc)
	if err != nil {
		fmt.Println(err.Error())
	}
	start, err := time.ParseInLocation(DATE_FORMAT, "1900/01/30", loc)
	if err != nil {
		fmt.Println(err.Error())
	}

	offset := BetweenDay(input, start)
	year := lunarYear(offset)

	lunar.leapMonth = leapMonth(year) //计算该年闰哪个月

	//设定当年是否有闰月
	if lunar.leapMonth > 0 {
		isLeapYear = true
	}

	for i = 1; i <= 12; i++ {
		if i == lunar.leapMonth+1 && isLeapYear {
			day = leapDay(year)
			isLeapYear = false
			lunar.isLeap = true
			i--
		} else {
			day = monthDays(year, i)
		}
		offset -= day
		if offset <= 0 {
			break
		}
	}
	offset += day
	lunar.Month = i
	lunar.Day = offset
	lunar.Year = year
	return lunar

}

func Solar2Lunar(time time.Time) string {
	lunar := CalculateLunar(time.Format(DATE_FORMAT))
	log.Println(lunar)
	result := StemBranchYear(lunar.Year) + "年"
	if lunar.isLeap && (lunar.Month == lunar.leapMonth) {
		result += "闰"
	}
	result += GetChineseMonth(lunar.Month)
	result += GetChineseDay(lunar.Day)
	return result
}

// 计算差的天数
func BetweenDay(d time.Time, s time.Time) int {
	subValue := float64(d.Unix()-s.Unix())/86400.0 + 0.5
	return int(subValue)
}

const DATE_FORMAT = "2006/01/02"
const YEAR_MIN = 1900
const YEAR_MAX = 2100

var Number = []string{`一`, `二`, `三`, `四`, `五`, `六`, `七`, `八`, `九`, `十`, `十一`, `十二`}
var ChineseNumber = []string{`正`, `二`, `三`, `四`, `五`, `六`, `七`, `八`, `九`, `十`, `十一`, `腊`}
var Ten = []string{`初`, `十`, `廿`, `卅`}

var HeavenlyStem = []string{
	`甲`, `乙`, `丙`, `丁`, `戊`, `己`, `庚`, `辛`, `壬`, `癸`,
}

var EarthyBranch = []string{
	`子`, `丑`, `寅`, `卯`, `辰`, `巳`, `午`, `未`, `申`, `酉`, `戌`, `亥`,
}

var Zodiac = []string{
	`鼠`, `牛`, `虎`, `兔`, `龙`, `蛇`, `马`, `羊`, `猴`, `鸡`, `狗`, `猪`,
}

var SolarTerm = []string{
	`小寒`, `大寒`, `立春`, `雨水`, `惊蛰`, `春分`, `清明`, `谷雨`, `立夏`, `小满`, `芒种`, `夏至`, `小暑`, `大暑`, `立秋`, `处暑`, `白露`, `秋分`, `寒露`, `霜降`, `立冬`, `小雪`, `大雪`, `冬至`,
}

var Astro = []string{
	`魔羯`, `水瓶`, `双鱼`, `白羊`, `金牛`, `双子`, `巨蟹`, `狮子`, `处女`, `天秤`, `天蝎`, `射手`,
}

//一月	二月	三月	四月	五月	六月	七月	八月	九月	十月	十一月	十二月 :年份
var ChinesMonth = []string{
	`丙寅`, `丁卯`, `戊辰`, `己巳`, `庚午`, `辛未`, `壬申`, `癸酉`, `甲戌`, `乙亥`, `丙子`, `丁丑`, //甲、己
	`戊寅`, `己卯`, `庚辰`, `辛巳`, `壬午`, `癸未`, `甲申`, `乙酉`, `丙戌`, `丁亥`, `戊子`, `己丑`, //乙、庚
	`庚寅`, `辛卯`, `壬辰`, `癸巳`, `甲午`, `乙未`, `丙申`, `丁酉`, `戊戌`, `己亥`, `庚子`, `辛丑`, //丙、辛
	`壬寅`, `癸卯`, `甲辰`, `乙巳`, `丙午`, `丁未`, `戊申`, `己酉`, `庚戌`, `辛亥`, `壬子`, `癸丑`, //丁、壬
	`甲寅`, `乙卯`, `丙辰`, `丁巳`, `戊午`, `己未`, `庚申`, `辛酉`, `壬戌`, `癸亥`, `甲子`, `乙丑`, //戊、癸
}

/**
 * 公历每个月份的天数普通表
 * @Array Of Property
 * @return Number
 */
var monthDay = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

var lunarInfo = []int{
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
var termInfo = []string{
	`9778397bd097c36b0b6fc9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`,
	`97bcf97c3598082c95f8c965cc920f`,
	`97bd0b06bdb0722c965ce1cfcc920f`,
	`b027097bd097c36b0b6fc9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`,
	`97bcf97c359801ec95f8c965cc920f`,
	`97bd0b06bdb0722c965ce1cfcc920f`,
	`b027097bd097c36b0b6fc9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`,
	`97bcf97c359801ec95f8c965cc920f`,
	`97bd0b06bdb0722c965ce1cfcc920f`,
	`b027097bd097c36b0b6fc9274c91aa`,
	`9778397bd19801ec9210c965cc920e`,
	`97b6b97bd19801ec95f8c965cc920f`,
	`97bd09801d98082c95f8e1cfcc920f`,
	`97bd097bd097c36b0b6fc9210c8dc2`,
	`9778397bd197c36c9210c9274c91aa`,
	`97b6b97bd19801ec95f8c965cc920e`,
	`97bd09801d98082c95f8e1cfcc920f`,
	`97bd097bd097c36b0b6fc9210c8dc2`,
	`9778397bd097c36c9210c9274c91aa`,
	`97b6b97bd19801ec95f8c965cc920e`,
	`97bcf97c3598082c95f8e1cfcc920f`,
	`97bd097bd097c36b0b6fc9210c8dc2`,
	`9778397bd097c36c9210c9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`,
	`97bcf97c3598082c95f8c965cc920f`,
	`97bd097bd097c35b0b6fc920fb0722`,
	`9778397bd097c36b0b6fc9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`,
	`97bcf97c3598082c95f8c965cc920f`,
	`97bd097bd097c35b0b6fc920fb0722`,
	`9778397bd097c36b0b6fc9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`,
	`97bcf97c359801ec95f8c965cc920f`,
	`97bd097bd097c35b0b6fc920fb0722`,
	`9778397bd097c36b0b6fc9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`,
	`97bcf97c359801ec95f8c965cc920f`,
	`97bd097bd097c35b0b6fc920fb0722`,
	`9778397bd097c36b0b6fc9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`,
	`97bcf97c359801ec95f8c965cc920f`,
	`97bd097bd07f595b0b6fc920fb0722`,
	`9778397bd097c36b0b6fc9210c8dc2`,
	`9778397bd19801ec9210c9274c920e`,
	`97b6b97bd19801ec95f8c965cc920f`,
	`97bd07f5307f595b0b0bc920fb0722`,
	`7f0e397bd097c36b0b6fc9210c8dc2`,
	`9778397bd097c36c9210c9274c920e`,
	`97b6b97bd19801ec95f8c965cc920f`,
	`97bd07f5307f595b0b0bc920fb0722`,
	`7f0e397bd097c36b0b6fc9210c8dc2`,
	`9778397bd097c36c9210c9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`,
	`97bd07f1487f595b0b0bc920fb0722`,
	`7f0e397bd097c36b0b6fc9210c8dc2`,
	`9778397bd097c36b0b6fc9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`,
	`97bcf7f1487f595b0b0bb0b6fb0722`,
	`7f0e397bd097c35b0b6fc920fb0722`,
	`9778397bd097c36b0b6fc9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`,
	`97bcf7f1487f595b0b0bb0b6fb0722`,
	`7f0e397bd097c35b0b6fc920fb0722`,
	`9778397bd097c36b0b6fc9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`,
	`97bcf7f1487f531b0b0bb0b6fb0722`,
	`7f0e397bd097c35b0b6fc920fb0722`,
	`9778397bd097c36b0b6fc9274c91aa`,
	`97b6b97bd19801ec9210c965cc920e`,
	`97bcf7f1487f531b0b0bb0b6fb0722`,
	`7f0e397bd07f595b0b6fc920fb0722`,
	`9778397bd097c36b0b6fc9274c91aa`,
	`97b6b97bd19801ec9210c9274c920e`,
	`97bcf7f0e47f531b0b0bb0b6fb0722`,
	`7f0e397bd07f595b0b0bc920fb0722`,
	`9778397bd097c36b0b6fc9210c91aa`,
	`97b6b97bd197c36c9210c9274c920e`,
	`97bcf7f0e47f531b0b0bb0b6fb0722`,
	`7f0e397bd07f595b0b0bc920fb0722`,
	`9778397bd097c36b0b6fc9210c8dc2`,
	`9778397bd097c36c9210c9274c920e`,
	`97b6b7f0e47f531b0723b0b6fb0722`,
	`7f0e37f5307f595b0b0bc920fb0722`,
	`7f0e397bd097c36b0b6fc9210c8dc2`,
	`9778397bd097c36b0b70c9274c91aa`,
	`97b6b7f0e47f531b0723b0b6fb0721`,
	`7f0e37f1487f595b0b0bb0b6fb0722`,
	`7f0e397bd097c35b0b6fc9210c8dc2`,
	`9778397bd097c36b0b6fc9274c91aa`,
	`97b6b7f0e47f531b0723b0b6fb0721`,
	`7f0e27f1487f595b0b0bb0b6fb0722`,
	`7f0e397bd097c35b0b6fc920fb0722`,
	`9778397bd097c36b0b6fc9274c91aa`,
	`97b6b7f0e47f531b0723b0b6fb0721`,
	`7f0e27f1487f531b0b0bb0b6fb0722`,
	`7f0e397bd097c35b0b6fc920fb0722`,
	`9778397bd097c36b0b6fc9274c91aa`,
	`97b6b7f0e47f531b0723b0b6fb0721`,
	`7f0e27f1487f531b0b0bb0b6fb0722`,
	`7f0e397bd097c35b0b6fc920fb0722`,
	`9778397bd097c36b0b6fc9274c91aa`,
	`97b6b7f0e47f531b0723b0b6fb0721`,
	`7f0e27f1487f531b0b0bb0b6fb0722`,
	`7f0e397bd07f595b0b0bc920fb0722`,
	`9778397bd097c36b0b6fc9274c91aa`,
	`97b6b7f0e47f531b0723b0787b0721`,
	`7f0e27f0e47f531b0b0bb0b6fb0722`,
	`7f0e397bd07f595b0b0bc920fb0722`,
	`9778397bd097c36b0b6fc9210c91aa`,
	`97b6b7f0e47f149b0723b0787b0721`,
	`7f0e27f0e47f531b0723b0b6fb0722`,
	`7f0e397bd07f595b0b0bc920fb0722`,
	`9778397bd097c36b0b6fc9210c8dc2`,
	`977837f0e37f149b0723b0787b0721`,
	`7f07e7f0e47f531b0723b0b6fb0722`,
	`7f0e37f5307f595b0b0bc920fb0722`,
	`7f0e397bd097c35b0b6fc9210c8dc2`,
	`977837f0e37f14998082b0787b0721`,
	`7f07e7f0e47f531b0723b0b6fb0721`,
	`7f0e37f1487f595b0b0bb0b6fb0722`,
	`7f0e397bd097c35b0b6fc9210c8dc2`,
	`977837f0e37f14998082b0787b06bd`,
	`7f07e7f0e47f531b0723b0b6fb0721`,
	`7f0e27f1487f531b0b0bb0b6fb0722`,
	`7f0e397bd097c35b0b6fc920fb0722`,
	`977837f0e37f14998082b0787b06bd`,
	`7f07e7f0e47f531b0723b0b6fb0721`,
	`7f0e27f1487f531b0b0bb0b6fb0722`,
	`7f0e397bd097c35b0b6fc920fb0722`,
	`977837f0e37f14998082b0787b06bd`,
	`7f07e7f0e47f531b0723b0b6fb0721`,
	`7f0e27f1487f531b0b0bb0b6fb0722`,
	`7f0e397bd07f595b0b0bc920fb0722`,
	`977837f0e37f14998082b0787b06bd`,
	`7f07e7f0e47f531b0723b0b6fb0721`,
	`7f0e27f1487f531b0b0bb0b6fb0722`,
	`7f0e397bd07f595b0b0bc920fb0722`,
	`977837f0e37f14998082b0787b06bd`,
	`7f07e7f0e47f149b0723b0787b0721`,
	`7f0e27f0e47f531b0b0bb0b6fb0722`,
	`7f0e397bd07f595b0b0bc920fb0722`,
	`977837f0e37f14998082b0723b06bd`,
	`7f07e7f0e37f149b0723b0787b0721`,
	`7f0e27f0e47f531b0723b0b6fb0722`,
	`7f0e397bd07f595b0b0bc920fb0722`,
	`977837f0e37f14898082b0723b02d5`,
	`7ec967f0e37f14998082b0787b0721`,
	`7f07e7f0e47f531b0723b0b6fb0722`,
	`7f0e37f1487f595b0b0bb0b6fb0722`,
	`7f0e37f0e37f14898082b0723b02d5`,
	`7ec967f0e37f14998082b0787b0721`,
	`7f07e7f0e47f531b0723b0b6fb0722`,
	`7f0e37f1487f531b0b0bb0b6fb0722`,
	`7f0e37f0e37f14898082b0723b02d5`,
	`7ec967f0e37f14998082b0787b06bd`,
	`7f07e7f0e47f531b0723b0b6fb0721`,
	`7f0e37f1487f531b0b0bb0b6fb0722`,
	`7f0e37f0e37f14898082b072297c35`,
	`7ec967f0e37f14998082b0787b06bd`,
	`7f07e7f0e47f531b0723b0b6fb0721`,
	`7f0e27f1487f531b0b0bb0b6fb0722`,
	`7f0e37f0e37f14898082b072297c35`,
	`7ec967f0e37f14998082b0787b06bd`,
	`7f07e7f0e47f531b0723b0b6fb0721`,
	`7f0e27f1487f531b0b0bb0b6fb0722`,
	`7f0e37f0e366aa89801eb072297c35`,
	`7ec967f0e37f14998082b0787b06bd`,
	`7f07e7f0e47f149b0723b0787b0721`,
	`7f0e27f1487f531b0b0bb0b6fb0722`,
	`7f0e37f0e366aa89801eb072297c35`,
	`7ec967f0e37f14998082b0723b06bd`,
	`7f07e7f0e47f149b0723b0787b0721`,
	`7f0e27f0e47f531b0723b0b6fb0722`,
	`7f0e37f0e366aa89801eb072297c35`,
	`7ec967f0e37f14998082b0723b06bd`,
	`7f07e7f0e37f14998083b0787b0721`,
	`7f0e27f0e47f531b0723b0b6fb0722`,
	`7f0e37f0e366aa89801eb072297c35`,
	`7ec967f0e37f14898082b0723b02d5`,
	`7f07e7f0e37f14998082b0787b0721`,
	`7f07e7f0e47f531b0723b0b6fb0722`,
	`7f0e36665b66aa89801e9808297c35`,
	`665f67f0e37f14898082b0723b02d5`,
	`7ec967f0e37f14998082b0787b0721`,
	`7f07e7f0e47f531b0723b0b6fb0722`,
	`7f0e36665b66a449801e9808297c35`,
	`665f67f0e37f14898082b0723b02d5`,
	`7ec967f0e37f14998082b0787b06bd`,
	`7f07e7f0e47f531b0723b0b6fb0721`,
	`7f0e36665b66a449801e9808297c35`,
	`665f67f0e37f14898082b072297c35`,
	`7ec967f0e37f14998082b0787b06bd`,
	`7f07e7f0e47f531b0723b0b6fb0721`,
	`7f0e26665b66a449801e9808297c35`,
	`665f67f0e37f1489801eb072297c35`,
	`7ec967f0e37f14998082b0787b06bd`,
	`7f07e7f0e47f531b0723b0b6fb0721`,
	`7f0e27f1487f531b0b0bb0b6fb0722`,
}

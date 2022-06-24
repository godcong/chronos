package chronos

import (
	"fmt"
	"strings"
	"time"
)

// Lunar ...
type Lunar struct {
	time.Time
	year      int
	month     int
	day       int
	hour      int
	leapMonth int
	leap      bool
	fixLiChun int //立春当天如果未到时辰：-1
}

var loc *time.Location

func init() {
	loc, _ = time.LoadLocation("Local")
}

func (lunar *Lunar) isLeap() bool {
	if lunar.leap && (lunar.month == lunar.leapMonth) {
		return true
	}
	return false
}

// Type ...
func (lunar *Lunar) Type() string {
	return "lunar"
}

func (lunar *Lunar) FixLiChun(fix int) {
	lunar.fixLiChun = fix
}

// Calendar ...
func (lunar *Lunar) Calendar() Calendar {
	t := time.Time{}
	t.AddDate(lunar.year, lunar.month, lunar.day)
	return New(t)
}

// EightCharacter ...
func (lunar *Lunar) EightCharacter() []string {
	rlt := lunar.nianZhu(lunar.fixLiChun) + lunar.yueZhu() + lunar.riZhu() + lunar.shiZhu()
	return strings.Split(rlt, "")
}

//shiZhu 时柱
func (lunar *Lunar) shiZhu() string {
	return StemBranchHour(lunar.Year(), int(lunar.Month()), lunar.Day(), lunar.Hour())
}

//riZhu 日柱
func (lunar *Lunar) riZhu() string {
	if lunar.Hour() >= 23 {
		return StemBranchDay(lunar.Year(), int(lunar.Month()), lunar.Day()+1)
	}
	return StemBranchDay(lunar.Year(), int(lunar.Month()), lunar.Day())
}

//yueZhu 月柱
func (lunar *Lunar) yueZhu() string {
	return StemBranchMonth(lunar.Year(), int(lunar.Month()), lunar.Day())
}

//nianZhu 年柱
func (lunar *Lunar) nianZhu(fix int) string {
	//log.Println("year", lunar.Year(), "nyear", lunar.year, "month", lunar.Month(), "day", lunar.Day(), "lichun", getLiChunDay(lunar.Year()))
	if lunar.Month() > 2 || (lunar.Month() == 2 && lunar.Day() >= getLiChunDay(lunar.Year())) {
		return StemBranchYear(lunar.Year() + fix)
	}
	return StemBranchYear(lunar.Year() - 1)
}

// GetZodiac ...
func GetZodiac(lunar *Lunar) string {
	s := string([]rune(lunar.nianZhu(lunar.fixLiChun))[1])
	for idx, v := range earthyBranch {
		if strings.Compare(v, s) == 0 {
			return zodiacs[idx]
		}
	}
	return ""
}

func yearDay(y int) int {
	i, sum := 348, 348
	for i = 0x8000; i > 0x8; i >>= 1 {
		if (GetLunarInfo(y) & i) != 0 {
			sum++
		}
	}
	return sum + leapDay(y)
}

func leapDay(y int) int {
	if leapMonth(y) != 0 {
		if (GetLunarInfo(y) & 0x10000) != 0 {
			return 30
		}
		return 29
	}
	return 0
}

func leapMonth(y int) int {
	return GetLunarInfo(y) & 0xf
}

func monthDays(y int, m int) int {
	//月份参数从1至12，参数错误返回-1
	if m > 12 || m < 1 {
		return -1
	}
	if GetLunarInfo(y)&(0x10000>>uint32(m)) != 0 {
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
	}
	return monthDay[idx]
}

//GetAstro 取得星座
func GetAstro(m, d int) string {
	arr := []int{20, 19, 21, 21, 21, 22, 23, 23, 23, 23, 22, 22}
	idx := d < arr[m-1]
	index := m * 2
	if idx {
		index = m*2 - 2
	}
	return constellation[index] + "座"
}

func lunarYear(offset int) (int, int) {
	day := 0
	i := 0
	//求当年农历年天数
	for i = yearMin; i <= yearMax; i++ {
		day = yearDay(i)
		if offset-day < 1 {
			break
		}
		offset -= day
	}
	return i, offset
}

func lunarStart() time.Time {
	loc, _ := time.LoadLocation("Local")
	start, err := time.ParseInLocation("2006/01/02", "1900/01/30", loc)
	if err != nil {
		fmt.Println(err.Error())
	}
	return start
}

func lunarInput(date string) time.Time {

	input, err := time.ParseInLocation(DefaultDateFormat, date, loc)
	if err != nil {
		fmt.Println(err.Error())
		return time.Time{}
	}
	//newInput, err := time.ParseInLocation(LunarDateFormat, input.Format(LunarDateFormat), loc)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return time.Time{}
	//}
	return input
}

// CalculateLunar ...
func CalculateLunar(date string) *Lunar {
	input := lunarInput(date)
	lunar := Lunar{
		Time: input,
		leap: false,
	}

	i, day := 0, 0
	isLeapYear := false

	start := lunarStart()
	offset := betweenDay(input, start)
	year, offset := lunarYear(offset)
	lunar.leapMonth = leapMonth(year) //计算该年闰哪个月

	//设定当年是否有闰月
	if lunar.leapMonth > 0 {
		isLeapYear = true
	}

	for i = 1; i <= 12; i++ {
		if i == lunar.leapMonth+1 && isLeapYear {
			day = leapDay(year)
			isLeapYear = false
			lunar.leap = true
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
	lunar.month = i
	lunar.day = offset
	lunar.year = year
	return &lunar

}

//betweenDay 计算两个时间差的天数
func betweenDay(d time.Time, s time.Time) int {
	newInput, err := time.ParseInLocation(LunarDateFormat, d.Format(LunarDateFormat), loc)
	if err != nil {
		return 0
	}

	subValue := float64(newInput.Unix()-s.Unix())/86400.0 + 0.5
	return int(subValue)
}

//Solar2Lunar 输入日历输出月历
func Solar2Lunar(time time.Time) string {
	lunar := CalculateLunar(time.Format(DefaultDateFormat))
	result := StemBranchYear(lunar.year) + "年"
	if lunar.leap && (lunar.month == lunar.leapMonth) {
		result += "闰"
	}
	result += getChineseMonth(lunar.month)
	result += getChineseDay(lunar.day)
	return result
}

// Date ...
func (lunar *Lunar) Date() string {
	result := getChineseYear(lunar.year)
	if lunar.isLeap() {
		result += "闰"
	}
	result += getChineseMonth(lunar.month)
	result += getChineseDay(lunar.day)
	return result
}

const yearMin = 1900
const yearMax = 2100

var solarTerms = []string{
	`小寒`, `大寒`, `立春`, `雨水`, `惊蛰`, `春分`, `清明`, `谷雨`, `立夏`, `小满`, `芒种`, `夏至`, `小暑`, `大暑`, `立秋`, `处暑`, `白露`, `秋分`, `寒露`, `霜降`, `立冬`, `小雪`, `大雪`, `冬至`,
}

//一月	二月	三月	四月	五月	六月	七月	八月	九月	十月	十一月	十二月 :年份
//var ChinesMonth = []string{
//	`丙寅`, `丁卯`, `戊辰`, `己巳`, `庚午`, `辛未`, `壬申`, `癸酉`, `甲戌`, `乙亥`, `丙子`, `丁丑`, //甲、己
//	`戊寅`, `己卯`, `庚辰`, `辛巳`, `壬午`, `癸未`, `甲申`, `乙酉`, `丙戌`, `丁亥`, `戊子`, `己丑`, //乙、庚
//	`庚寅`, `辛卯`, `壬辰`, `癸巳`, `甲午`, `乙未`, `丙申`, `丁酉`, `戊戌`, `己亥`, `庚子`, `辛丑`, //丙、辛
//	`壬寅`, `癸卯`, `甲辰`, `乙巳`, `丙午`, `丁未`, `戊申`, `己酉`, `庚戌`, `辛亥`, `壬子`, `癸丑`, //丁、壬
//	`甲寅`, `乙卯`, `丙辰`, `丁巳`, `戊午`, `己未`, `庚申`, `辛酉`, `壬戌`, `癸亥`, `甲子`, `乙丑`, //戊、癸
//}

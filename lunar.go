package lunar

import (
	"time"
	"fmt"
)

type Lunar struct {
	year      int
	month     int
	day       int
	hour      int
	leapMonth int
	isLeap    bool
}

func (lunar *Lunar) Type() string {
	return "lunar"
}

func (lunar *Lunar) Calendar() Calendar {
	return CalendarFromLunar(lunar.year, lunar.month, lunar.day)
}

func GetZodiac(time time.Time) string {
	return Zodiac[(time.Year()-4)%12]
}

//NewLunar 取得月历
// 默认返回当前时间月历
func NewLunar(calendar Calendar) *Lunar {
	t := time.Now()
	if calendar != nil {
		if calendar.Lunar() != nil {
			return calendar.Lunar()
		}
		if calendar.Solar() != nil {
			t = calendar.Solar().time
		}
	}
	return CalculateLunar(t.Format(DATE_FORMAT))
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
	} else {
		return monthDay[idx]
	}
}

//GetAstro 取得星座
func GetAstro(m, d int) string {
	arr := []int{20, 19, 21, 21, 21, 22, 23, 23, 23, 23, 22, 22}
	idx := d < arr[m-1]
	index := m * 2
	if idx {
		index = m*2 - 2
	}
	return Astro[index] + "座"
}

func lunarYear(offset int) (int, int) {
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
	return i, offset
}

func CalculateLunar(date string) *Lunar {
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
	lunar.month = i
	lunar.day = offset
	lunar.year = year
	return &lunar

}

//BetweenDay 计算两个时间差的天数
func BetweenDay(d time.Time, s time.Time) int {
	subValue := float64(d.Unix()-s.Unix())/86400.0 + 0.5
	return int(subValue)
}

//Solar2Lunar 输入日历输出月历
func Solar2Lunar(time time.Time) string {
	lunar := CalculateLunar(time.Format(DATE_FORMAT))
	result := StemBranchYear(lunar.year) + "年"
	if lunar.isLeap && (lunar.month == lunar.leapMonth) {
		result += "闰"
	}
	result += GetChineseMonth(lunar.month)
	result += GetChineseDay(lunar.day)
	return result
}
func (lunar *Lunar) Date() string {
	result := StemBranchYear(lunar.year) + "年"
	if lunar.isLeap && (lunar.month == lunar.leapMonth) {
		result += "闰"
	}
	result += GetChineseMonth(lunar.month)
	result += GetChineseDay(lunar.day)
	return result
}

const YEAR_MIN = 1900
const YEAR_MAX = 2100

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
//var ChinesMonth = []string{
//	`丙寅`, `丁卯`, `戊辰`, `己巳`, `庚午`, `辛未`, `壬申`, `癸酉`, `甲戌`, `乙亥`, `丙子`, `丁丑`, //甲、己
//	`戊寅`, `己卯`, `庚辰`, `辛巳`, `壬午`, `癸未`, `甲申`, `乙酉`, `丙戌`, `丁亥`, `戊子`, `己丑`, //乙、庚
//	`庚寅`, `辛卯`, `壬辰`, `癸巳`, `甲午`, `乙未`, `丙申`, `丁酉`, `戊戌`, `己亥`, `庚子`, `辛丑`, //丙、辛
//	`壬寅`, `癸卯`, `甲辰`, `乙巳`, `丙午`, `丁未`, `戊申`, `己酉`, `庚戌`, `辛亥`, `壬子`, `癸丑`, //丁、壬
//	`甲寅`, `乙卯`, `丙辰`, `丁巳`, `戊午`, `己未`, `庚申`, `辛酉`, `壬戌`, `癸亥`, `甲子`, `乙丑`, //戊、癸
//}

/**
 * 公历每个月份的天数普通表
 * @Array Of Property
 * @return Number
 */
var monthDay = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

package chronos

import (
	"fmt"
	"strings"
	"time"

	"github.com/godcong/chronos/v2/utils"
)

// minLunarYear 最小可转换年
const minLunarYear = 1900

// maxLunarYear 最大可转换年
const maxLunarYear = 3000

// lunar ...
type lunar struct {
	year      int
	month     int
	day       int
	hour      int
	leapMonth int
	leap      bool
}

func (l *lunar) LeapMonth() int {
	return yearLeapMonth(l.year)
}

func (l *lunar) Date() LunarDate {
	return LunarDate{
		Year:        l.year,
		Month:       l.month,
		Day:         l.day,
		Hour:        l.hour,
		IsLeapMonth: l.IsLeapMonth(),
		LeapMonth:   l.LeapMonth(),
	}
}

func (l *lunar) IsLeapMonth() bool {
	return l.month == l.leapMonth
}

func (l *lunar) Year() int {
	return l.year
}

func (l *lunar) Day() int {
	return l.day
}

func (l *lunar) Hour() int {
	return l.hour
}

//var loc *time.Location

func init() {
	//loc, _ = time.LoadLocation("Local")
}

func (l *lunar) isLeap() bool {
	if l.leap && (l.month == l.leapMonth) {
		return true
	}
	return false
}

func (l *lunar) Month() int {
	return l.month
}

// Calendar ...
func (l *lunar) Calendar() Calendar {
	t := time.Time{}
	t.AddDate(l.year, l.month, l.day)
	return NewSolarCalendar(t)
}

// EightCharacter ...
func (l *lunar) EightCharacter() []string {
	rlt := l.nianZhu() + l.yueZhu() + l.riZhu() + l.shiZhu()
	return strings.Split(rlt, "")
}

//shiZhu 时柱
func (l *lunar) shiZhu() string {
	return shiZhu(l.year, time.Month(l.month), l.Day(), l.Hour()).Chinese()
}

//riZhu 日柱
func (l *lunar) riZhu() string {
	return riZhu(l.year, time.Month(l.month), l.Day()).Chinese()
}

//yueZhu 月柱
func (l *lunar) yueZhu() string {
	return yueZhu(l.Year(), time.Month(l.Month()), l.Day()).Chinese()
}

//nianZhu 年柱
func (l *lunar) nianZhu() string {
	if l.Month() > 2 || (l.Month() == 2 && l.Day() >= yearLiChunDay(l.Year())) {
		return nianZhuChinese(l.Year())
	}
	return nianZhuChinese(l.Year() - 1)
}

func yearDay(y int) int {
	i, sum := 348, 348
	for i = 0x8000; i > 0x8; i >>= 1 {
		if (getLunarInfo(y) & i) != 0 {
			sum++
		}
	}
	//offset := yearOffset(y)
	return sum + leapDay(y)
}

func yearDayOld(y int) int {
	current := getYearSolarTermTime(y, SolarTermLiChun)
	next := getYearSolarTermTime(y+1, SolarTermLiChun)
	fmt.Println("day", current.Format(DateFormatYMDHMS), "next", next.Format(DateFormatYMDHMS))
	return utils.BetweenDay(next, current)
}

func leapDay(y int) int {
	if yearLeapMonth(y) != 0 {
		if yearLeapMonthBS(y) == LeapMonthBig {
			return 30
		}
		return 29
	}
	return 0
}

func monthDays(y int, m int, leapMonth int, isleap bool) int {
	days := utils.CalcYearMonthDays(y)

	if (isleap && m == leapMonth) || leapMonth > 0 && m > leapMonth {
		return days[m]
	}
	return days[m-1]
}

func solarDays(y, m int) int {
	if m == 2 && ((y%4 == 0) && (y%100 != 0) || (y%400 == 0)) { //2月份的闰平规律测算后确认返回28或29
		return monthDay[m-1] + 1
	}
	return monthDay[m-1]
}

func lunarYear(offset int) (int, int) {
	day := 0
	i := 0
	//求当年农历年天数
	for i = minLunarYear; i <= maxLunarYear; i++ {
		day = yearDay(i)
		if offset-day < 1 {
			break
		}
		offset -= day
	}
	return i, offset
}

func lunarByTime(t time.Time) *lunar {
	//todo: use lunar time instead of solar time
	return &lunar{
		year:      t.Year(),
		month:     int(t.Month()),
		day:       t.Day(),
		hour:      t.Hour(),
		leapMonth: yearLeapMonth(t.Year()),
		leap:      int(t.Month()) == yearLeapMonth(t.Year()),
	}
}

// ParseLunarString
// @param string
// @return time.Time
// @return error
func lunarByString(date string) (time.Time, error) {
	t, err := time.ParseInLocation(DateFormatYMDHMS, date, time.Local)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// calculateLunar ...
func calculateLunar(t time.Time) *lunar {
	//input, _ := lunarByString(date)
	lunar := &lunar{
		leap: false,
	}

	i, day := 0, 0
	isLeapYear := false

	start := lunarStartTime
	offset := utils.BetweenDay(t, start)
	year, offset := lunarYear(offset)
	lunar.leapMonth = yearLeapMonth(year) //计算该年闰哪个月

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
			day = monthDays(year, i, lunar.leapMonth, false)
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
	return lunar

}

//solar2Lunar 输入日历输出月历
func solar2Lunar(t time.Time) string {
	lunar := calculateLunar(t)
	result := nianZhuChinese(lunar.year) + "年"
	if lunar.leap && (lunar.month == lunar.leapMonth) {
		result += "闰"
	}
	result += getChineseMonth(lunar.month)
	result += getChineseDay(lunar.day)
	return result
}

// Date ...
func (l *lunar) Date2() string {
	result := getChineseYear(l.year)
	if l.isLeap() {
		result += "闰"
	}
	result += getChineseMonth(l.month)
	result += getChineseDay(l.day)
	return result
}

//一月	二月	三月	四月	五月	六月	七月	八月	九月	十月	十一月	十二月 :年份
//var ChinesMonth = []string{
//	`丙寅`, `丁卯`, `戊辰`, `己巳`, `庚午`, `辛未`, `壬申`, `癸酉`, `甲戌`, `乙亥`, `丙子`, `丁丑`, //甲、己
//	`戊寅`, `己卯`, `庚辰`, `辛巳`, `壬午`, `癸未`, `甲申`, `乙酉`, `丙戌`, `丁亥`, `戊子`, `己丑`, //乙、庚
//	`庚寅`, `辛卯`, `壬辰`, `癸巳`, `甲午`, `乙未`, `丙申`, `丁酉`, `戊戌`, `己亥`, `庚子`, `辛丑`, //丙、辛
//	`壬寅`, `癸卯`, `甲辰`, `乙巳`, `丙午`, `丁未`, `戊申`, `己酉`, `庚戌`, `辛亥`, `壬子`, `癸丑`, //丁、壬
//	`甲寅`, `乙卯`, `丙辰`, `丁巳`, `戊午`, `己未`, `庚申`, `辛酉`, `壬戌`, `癸亥`, `甲子`, `乙丑`, //戊、癸
//}

var _ Lunar = &lunar{}

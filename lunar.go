package chronos

import (
	"fmt"
	"strings"
	"time"

	lc "github.com/6tail/lunar-go/calendar"

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
	minute    int
	second    int
	leapMonth int
	leap      bool
	date      time.Time
}

func (l *lunar) LeapMonth() int {
	return l.leapMonth
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

// shiZhu 时柱
func (l *lunar) shiZhu() string {
	return shiZhu(l.year, time.Month(l.month), l.Day(), l.Hour()).Chinese()
}

// riZhu 日柱
func (l *lunar) riZhu() string {
	return riZhu(l.year, time.Month(l.month), l.Day()).Chinese()
}

// yueZhu 月柱
func (l *lunar) yueZhu() string {
	return yueZhu(l.Year(), time.Month(l.Month()), l.Day()).Chinese()
}

// nianZhu 年柱
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
	days := utils.YearLunarMonthDays(y)
	days = days[2:]
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

func calcLunarYear(offset int) (int, int) {
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

func ParseLunarTime(date time.Time) *lunar {
	// 节气(中午12点)，长度27
	//jq := make([]float64, 27)
	// 合朔，即每月初一(中午12点)，长度16
	//hs := make([]float64, 16)
	// 每月天数，长度15
	currentYear := date.Year()
	currentMonth := date.Month()
	currentDay := date.Day()
	lunarYear := 0
	lunarMonth := 0
	lunarDay := 0
	//year := currentYear - 2000

	// 从上年的大雪到下年的立春
	//j := len(JIE_QI_IN_USE)
	//lunarYear.jieQiJulianDays = make([]float64, j)
	jq := utils.YearSolarTermList(currentYear)
	hs := utils.YearHeshuoList(currentYear)
	dayCounts := utils.YearLunarMonthDays(currentYear)

	currentYearLeap, exists := utils.LeapList[currentYear]
	if !exists {
		currentYearLeap = -1
		if hs[13] <= jq[24] {
			i := 1
			for {
				if hs[i+1] <= jq[2*i] {
					break
				}
				if i >= 13 {
					break
				}
				i++
			}
			currentYearLeap = i
		}
	}

	prevYear := currentYear - 1
	prevYearLeap, exists := utils.LeapList[prevYear]
	if !exists {
		prevYearLeap = -1
	} else {
		prevYearLeap -= 12
	}
	y := prevYear
	m := 11
	for i := 0; i < 15; i++ {
		cm := m
		isNextLeap := false
		if y == currentYear && i == currentYearLeap {
			cm = -cm
		} else if y == prevYear && i == prevYearLeap {
			cm = -cm
		}
		if y == currentYear && i+1 == currentYearLeap {
			isNextLeap = true
		} else if y == prevYear && i+1 == prevYearLeap {
			isNextLeap = true
		}

		lm := newLunarMonth(cm, dayCounts[i], hs[i]+J2000)
		firstDay := utils.JulianDayTime(hs[i] + J2000)
		days := utils.BetweenDay(TimeFromYmd(firstDay.Year(), firstDay.Month(), firstDay.Day()), TimeFromYmd(currentYear, currentMonth, currentDay))
		if days < lm.DayCount() {
			lunarYear = y
			lunarMonth = lm.Month()
			lunarDay = days + 1
			break
		}
		if !isNextLeap {
			m++
		}
		if m == 13 {
			m = 1
			y++
		}
	}

	return &lunar{
		date:      date,
		year:      lunarYear,
		month:     lunarMonth,
		day:       lunarDay,
		hour:      date.Hour(),
		minute:    date.Minute(),
		second:    date.Second(),
		leapMonth: yearLeapMonth(date.Year()),
		leap:      int(date.Month()) == yearLeapMonth(date.Year()),
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

// lunarFromSolar ...
func lunarFromSolar(t time.Time) *lunar {
	date := lc.NewLunarFromDate(t)

	//input, _ := lunarByString(date)
	month := date.GetMonth()
	isLeap := false
	if month < 0 {
		month = -month
		isLeap = true
	}
	lunar := &lunar{
		year:      date.GetYear(),
		month:     month,
		day:       date.GetDay(),
		hour:      date.GetHour(),
		minute:    date.GetMinute(),
		second:    date.GetSecond(),
		leapMonth: yearLeapMonth(date.GetYear()),
		leap:      isLeap,
		date:      time.Time{},
	}
	return lunar

}

// solar2Lunar 输入日历输出月历
func solar2Lunar(t time.Time) string {
	lunar := lunarFromSolar(t)
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

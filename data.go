package chronos

var chineseNumbers = [...]string{"一", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "十二"}

var chineseTenPrefix = [...]string{"初", "十", "廿", "卅"}

var chineseMonthNames = [...]string{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "冬", "腊"}

func getChineseYear(year int) string {
	return nianZhuChinese(year) + "年"
}

func getChineseMonth(m int) string {
	if m > 12 || m < 1 {
		return "?月"
	}
	return chineseMonthNames[m-1] + "月"
}

func getChineseDay(d int) string {
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
		s = chineseTenPrefix[d/10] + chineseNumbers[n]
	}
	return s + "日"
}

package chronos

import (
	"github.com/godcong/chronos/v2/runes"
)

var number = runes.Runes(`一二三四五六七八九十十一十二`)
var ten = runes.Runes(`初十廿卅`)

var chineseNumber = runes.Runes(`正二三四五六七八九十冬腊`)

func getChineseYear(year int) string {
	return nianZhuChinese(year) + "年"
}

func getChineseMonth(m int) string {
	if m > 12 || m < 1 {
		return "?月"
	}
	return chineseNumber.MustReadString(m-1, 1) + "月"
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
		s = ten.MustReadString(d/10, 1) + number.MustReadString(n, 1)
	}
	return s + "日"
}

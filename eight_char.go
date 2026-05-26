package chronos

import (
	"container/list"

	"github.com/6tail/lunar-go/calendar"
)

type eightChar struct {
	*calendar.EightChar
}

type EightCharIndex int

const (
	EightCharYear EightCharIndex = iota
	EightCharMonth
	EightCharDay
	EightCharTime
	eightCharMax
)

var _ EightChar = (*eightChar)(nil)

func (e *eightChar) GetNaYin() [4]string {
	return [4]string{
		e.GetYearNaYin(),
		e.GetMonthNaYin(),
		e.GetDayNaYin(),
		e.GetTimeNaYin(),
	}
}

func (e *eightChar) GetSiZhu() [4]string {
	return [4]string{
		e.GetYear(),
		e.GetMonth(),
		e.GetDay(),
		e.GetTime(),
	}
}

func (e *eightChar) GetWuXing() [4]string {
	return [4]string{
		e.GetYearWuXing(),
		e.GetMonthWuXing(),
		e.GetDayWuXing(),
		e.GetTimeWuXing(),
	}
}

func (e *eightChar) GetCangGan() [4][]string {
	return [4][]string{
		diZhiCangGan[e.GetYearZhi()],
		diZhiCangGan[e.GetMonthZhi()],
		diZhiCangGan[e.GetDayZhi()],
		diZhiCangGan[e.GetTimeZhi()],
	}
}

func (e *eightChar) GetShiShenGan() [4]string {
	return [4]string{
		e.GetYearShiShenGan(),
		e.GetMonthShiShenGan(),
		e.GetDayShiShenGan(),
		e.GetTimeShiShenGan(),
	}
}

func (e *eightChar) GetShiShenZhi() [4][]string {
	return [4][]string{
		getShiShenZhiString(e.GetYearShiShenZhi()),
		getShiShenZhiString(e.GetMonthShiShenZhi()),
		getShiShenZhiString(e.GetDayShiShenZhi()),
		getShiShenZhiString(e.GetTimeShiShenZhi()),
	}
}

func (e *eightChar) GetDaYun(sex int) []int {
	dayun := e.GetYun(sex).GetDaYunBy(11)
	if len(dayun) <= 0 {
		return nil
	}
	var result []int
	for i := 1; i < len(dayun); i++ {
		result = append(result, dayun[i].GetStartYear())
	}
	return result
}

func getShiShenZhiString(l *list.List) []string {
	var result []string
	for e := l.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(string))
	}
	return result
}

package chronos

import (
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

func (e *eightChar) GetShiShenGan() [4]string {
	return [4]string{
		e.GetYearShiShenGan(),
		e.GetMonthShiShenGan(),
		e.GetDayShiShenGan(),
		e.GetTimeShiShenGan(),
	}
}

func (e *eightChar) GetShiShenZhi() [4]string {
	return [4]string{
		e.GetYearShiShenZhi().Front().Value.(string),
		e.GetMonthShiShenZhi().Front().Value.(string),
		e.GetDayShiShenZhi().Front().Value.(string),
		e.GetTimeShiShenZhi().Front().Value.(string),
	}
}

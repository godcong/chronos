package chronos

import (
	"container/list"

	"github.com/6tail/lunar-go/calendar"

	"github.com/godcong/chronos/runes"
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

// 1、子中藏‘癸’，为子与癸具有比肩之义。
// 2、丑中藏‘己辛癸’（1）丑与己有比肩之义；（2）辛是己和丑的食神；（3）癸为己和丑的偏财。
// 3、寅中藏‘甲丙戊’（1）寅与甲具有比肩之义；（2）丙是甲和寅的食神；（3）戊是甲和寅的偏财。
// 4、卯中藏‘乙’，卯与乙具有比肩之义。
// 5、辰中藏‘乙戊癸’。（1）辰与戊具有比肩之义；（2）乙是戊和辰的正官；（3）癸是戊和辰的正财。
// 6、巳中藏‘庚丙戊’。（1）巳与丙具有比肩之义；（2）庚是丙和巳的偏财；（3）戊是丙和巳的食神。
// 7、午中藏‘己丁’。（1）戊与丁具有比肩之义；（2）己是戊和丁的食神。
// 8、未中藏‘乙己丁’。（1）未与己具有比肩之义；（2）乙是己和未的偏官；（3）丁是己和未的偏印。
// 9、申中藏‘戊庚壬’。（1）申与庚具有比肩之义；（2）戊是庚和申的偏印；（3）壬是庚和申的食神。
// 10、酉中藏‘辛’。酉与辛具有比肩之义。
// 11、戌中藏‘辛丁戊’。（1）戌和戊具有比肩之义；（2）丁是戌和戊的正印；（3）辛是戌和戊的伤官。
// 12、亥中藏‘甲壬’。（1）亥和壬具有比肩之义；（2）甲是壬和亥的食神。
var _cangGan = [12]runes.Runes{
	runes.Runes("癸"),
	runes.Runes("己辛癸"),
	runes.Runes("甲丙戊"),
	runes.Runes("乙"),
	runes.Runes("乙戊癸"),
	runes.Runes("庚丙戊"),
	runes.Runes("丁己"),
	runes.Runes("乙丁己"),
	runes.Runes("戊庚壬"),
	runes.Runes("辛"),
	runes.Runes("辛丁戊"),
	runes.Runes("甲壬"),
}

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
		_cangGan[_DiZhiTable.FindString(e.GetYearZhi())].StringArray(),
		_cangGan[_DiZhiTable.FindString(e.GetMonthZhi())].StringArray(),
		_cangGan[_DiZhiTable.FindString(e.GetDayZhi())].StringArray(),
		_cangGan[_DiZhiTable.FindString(e.GetTimeZhi())].StringArray(),
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

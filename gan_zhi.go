package chronos

import (
	"errors"

	"github.com/godcong/chronos/v2/runes"
)

const defaultTianGan = "天干"
const defaultDiZhi = "地支"

//TianGan
//ENUM(Jia , Yi , Bing , Ding , Wu , Ji , Geng , Xin , Ren , Gui , Max)
type TianGan uint32

//DiZhi
//ENUM(Zi , Chou , Yin , Mao , Chen , Si , Wu , Wei , Shen , You , Xu , Hai , Max)
type DiZhi uint32

var _TianGanTable = runes.Runes(`甲乙丙丁戊己庚辛壬癸`)

// ErrWrongTianGanTypes returns an error
var ErrWrongTianGanTypes = errors.New("[chronos] wrong tiangan types")

var _DiZhiTable = runes.Runes(`子丑寅卯辰巳午未申酉戌亥`)

// ErrWrongDiZhiTypes returns an error
var ErrWrongDiZhiTypes = errors.New("[chronos] wrong dizhi types")

var _GanZhiTable = []string{
	`甲子`, `乙丑`, `丙寅`, `丁卯`, `戊辰`, `己巳`, `庚午`, `辛未`, `壬申`, `癸酉`, `甲戌`, `乙亥`, //甲或己日
	`丙子`, `丁丑`, `戊寅`, `己卯`, `庚辰`, `辛巳`, `壬午`, `癸未`, `甲申`, `乙酉`, `丙戌`, `丁亥`, //乙或庚日
	`戊子`, `己丑`, `庚寅`, `辛卯`, `壬辰`, `癸巳`, `甲午`, `乙未`, `丙申`, `丁酉`, `戊戌`, `己亥`, //丙或辛日
	`庚子`, `辛丑`, `壬寅`, `癸卯`, `甲辰`, `乙巳`, `丙午`, `丁未`, `戊申`, `己酉`, `庚戌`, `辛亥`, //丁或壬日
	`壬子`, `癸丑`, `甲寅`, `乙卯`, `丙辰`, `丁巳`, `戊午`, `己未`, `庚申`, `辛酉`, `壬戌`, `癸亥`, //戊或癸日
}

// PillarHour is an alias name of ShiZhu
var PillarHour = ShiZhu
var PillarDay = RiZhu
var PillarMonty = YueZhu
var PillarYear = NianZhuChineseV2

// YearGanZhiChineseV2 returns the year of ganzhi string
// @param int
// @return string
func YearGanZhiChineseV2(y int) string {
	return TianGanChineseV2(TianGan(y%10)) + DiZhiChineseV2(DiZhi(y%12))
}

// YearGanZhiChinese returns the year of the chinese ganzhi string
// @param int
// @return string
// @return error
func YearGanZhiChinese(y int) (string, error) {
	tgc, err := TianGanChinese(TianGan(y % 10))
	if err != nil {
		return "", err
	}
	dzc, err := DiZhiChinese(DiZhi(y % 12))
	if err != nil {
		return "", err
	}
	return tgc + dzc, nil
}

// TianGanChineseV2 returns the chinese tiangan string
// @param TianGan
// @return string
func TianGanChineseV2(tiangan TianGan) string {
	return _TianGanTable.MustReadString(int(tiangan), 1)
}

// TianGanChinese returns the chinese tiangan string
// @param TianGan
// @return string
// @return error
func TianGanChinese(tiangan TianGan) (string, error) {
	readString, err := _TianGanTable.ReadString(int(tiangan), 1)
	if err != nil {
		return "", ErrWrongTianGanTypes
	}
	return readString, nil
}

// DiZhiChineseV2 returns the chinese dizhi string
// @param DiZhi
// @return string
func DiZhiChineseV2(dizhi DiZhi) string {
	return _DiZhiTable.MustReadString(int(dizhi), 1)
}

// DiZhiChinese returns the chinese dizhi string
// @param DiZhi
// @return string
// @return error
func DiZhiChinese(tiangan DiZhi) (string, error) {
	readString, err := _DiZhiTable.ReadString(int(tiangan), 1)
	if err != nil {
		return "", ErrWrongDiZhiTypes
	}
	return readString, nil
}

//ShiZhu 获取时柱
//　	子 　　丑 　　寅 　　卯 　　辰 　　己
//　　　23-01：01-03：03-05 :05-07：07-09：09-11
//　　　午 　　未 　　申 　　酉 　　戊 　　亥
//　　　11-13：13-15：15-17：17-19：19-21：21-23
//`甲子`, `乙丑`, `丙寅`, `丁卯`, `戊辰`, `己巳`, `庚午`, `辛未`, `壬申`, `癸酉`, `甲戌`, `乙亥`, //甲或己日
//`丙子`, `丁丑`, `戊寅`, `己卯`, `庚辰`, `辛巳`, `壬午`, `癸未`, `甲申`, `乙酉`, `丙戌`, `丁亥`, //乙或庚日
//`戊子`, `己丑`, `庚寅`, `辛卯`, `壬辰`, `癸巳`, `甲午`, `乙未`, `丙申`, `丁酉`, `戊戌`, `己亥`, //丙或辛日
//`庚子`, `辛丑`, `壬寅`, `癸卯`, `甲辰`, `乙巳`, `丙午`, `丁未`, `戊申`, `己酉`, `庚戌`, `辛亥`, //丁或壬日
//`壬子`, `癸丑`, `甲寅`, `乙卯`, `丙辰`, `丁巳`, `戊午`, `己未`, `庚申`, `辛酉`, `壬戌`, `癸亥`, //戊或癸日
func ShiZhu(y, m, d, h int) string {
	i := stemBranchIndex(y, m, d) % 5 * 12
	idx := (h + 1) / 2 % 12
	return _GanZhiTable[fixDayNext(i, idx, h)]
}

// RiZhu 获取日柱
func RiZhu(y, m, d int) string {
	return _GanZhiTable[stemBranchIndex(y, m, d)]
}

//YueZhu 获取月柱
func YueZhu(y, m, d int) string {
	//月柱 1900年1月小寒以前为 丙子月(60进制12)
	fir := GetTermInfo(y, m*2-1) //返回当月「节」为几日开始

	//依据12节气修正干支月
	var sb = YearGanZhiChineseV2(fixSuffix(y)*12 + m + 11)
	if d >= fir {
		sb = YearGanZhiChineseV2(fixSuffix(y)*12 + m + 12)
	}
	return sb
}

// NianZhuChineseV2 returns the chinese nianzhu string
// @param int
// @return string
func NianZhuChineseV2(y int) string {
	return YearGanZhiChineseV2(y - 4)
}

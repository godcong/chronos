package chronos

import (
	"errors"
	"time"

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

// GanZhi
//ENUM(JiaZi,YiChou,BingYin,DingMao,WuChen,JiSi,GengWu,XinWei,RenShen,GuiYou,JiaXu,YiHai,
//BingZi,DingChou,WuYin,JiMao,GengChen,XinSi,RenWu,GuiWei,JiaShen,YiYou,BingXu,DingHai,
//WuZi,JiChou,GengYin,XinMao,RenChen,GuiSi,JiaWu,YiWei,BingShen,DingYou,WuXu,JiHai,
//GengZi,XinChou,RenYin,GuiMao,JiaChen,YiSi,BingWu,DingWei,WuShen,JiYou,GengXu,XinHai,
//RenZi,GuiChou,JiaYin,YiMao,BingChen,DingSi,WuWu,JiWei,GengShen,XinYou,RenXu,GuiHai,Max)
type GanZhi uint32

//StemBranch is an alias name for GanZhi
type StemBranch = GanZhi

var _TianGanTable = runes.Runes(`甲乙丙丁戊己庚辛壬癸`)

// ErrWrongTianGanTypes returns an error
var ErrWrongTianGanTypes = errors.New("[chronos] wrong tiangan types")

var _DiZhiTable = runes.Runes(`子丑寅卯辰巳午未申酉戌亥`)

// ErrWrongDiZhiTypes returns an error
var ErrWrongDiZhiTypes = errors.New("[chronos] wrong dizhi types")

var _GanZhiTable = runes.Runes(
	"甲子乙丑丙寅丁卯戊辰己巳庚午辛未壬申癸酉甲戌乙亥" +
		"丙子丁丑戊寅己卯庚辰辛巳壬午癸未甲申乙酉丙戌丁亥" +
		"戊子己丑庚寅辛卯壬辰癸巳甲午乙未丙申丁酉戊戌己亥" +
		"庚子辛丑壬寅癸卯甲辰乙巳丙午丁未戊申己酉庚戌辛亥" +
		"壬子癸丑甲寅乙卯丙辰丁巳戊午己未庚申辛酉壬戌癸亥")

// ErrWrongGanZhiTypes returns an error
var ErrWrongGanZhiTypes = errors.New("[chronos] wrong ganzhi types")

// PillarHour is an alias name of ShiZhu
var PillarHour = ShiZhu
var PillarDay = RiZhu
var PillarMonty = YueZhuChineseV2
var PillarYear = NianZhuChineseV2

func (x TianGan) Chinese() string {
	return TianGanChineseV2(x)
}

func (x DiZhi) Chinese() string {
	return DiZhiChineseV2(x)
}

func (x GanZhi) Chinese() string {
	return GanZhiChineseV2(x)
}

func (x GanZhi) index() int {
	return int(x * 2)
}

func nianZhu(year int) GanZhi {
	return GanZhi((year - 4) % 60)
}

func nianZhuChinese(year int) string {
	return nianZhu(year).Chinese()
	//return ganZhiChinese(year - 4)
}

func ganZhiChinese(gz int) string {
	return getTianGan(gz).Chinese() + getDiZhi(gz).Chinese()
}

func getGanZhi(v int) GanZhi {
	return GanZhi(v)
}

func getGanAndZhi(year int) (TianGan, DiZhi) {
	return getTianGan(year), getDiZhi(year)
}

func getTianGan(v int) TianGan {
	return TianGan(v % 10)
}

func getDiZhi(v int) DiZhi {
	return DiZhi(v % 12)
}

func splitGanZhi(gz GanZhi) (TianGan, DiZhi) {
	return TianGan(gz % 10), DiZhi(gz % 12)
}

// parseGanZhiV2
// @param TianGan
// @param DiZhi
// @return GanZhi
func parseGanZhiV2(tiangan TianGan, dizhi DiZhi) GanZhi {
	if tiangan >= TianGanMax || dizhi >= DiZhiMax {
		return GanZhiMax
	}
	return _TianGanDiZhiGanZhiTable[tiangan][dizhi]
}

// parseGanZhi
// @param TianGan
// @param DiZhi
// @return GanZhi
// decrypted use parseGanZhiV2
func parseGanZhi(tiangan TianGan, dizhi DiZhi) GanZhi {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if v := i*10 + int(tiangan); v == j*12+int(dizhi) {
				return GanZhi(v)
			}
		}
	}
	return GanZhiMax
}

// YearGanZhiChinese returns the year of the chinese GanZhi string
// @param int
// @return string
// @return error
func YearGanZhiChinese(t time.Time) (string, error) {
	tgc, err := TianGanChinese(TianGan(t.Year() % 10))
	if err != nil {
		return "", err
	}
	dzc, err := DiZhiChinese(DiZhi(t.Year() % 12))
	if err != nil {
		return "", err
	}
	return tgc + dzc, nil
}

// TianGanChineseV2 returns the chinese TianGan string
// @param TianGan
// @return string
func TianGanChineseV2(tiangan TianGan) string {
	return _TianGanTable.MustReadString(int(tiangan), 1)
}

// TianGanChinese returns the chinese TianGan string
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

// DiZhiChineseV2 returns the chinese DiZhi string
// @param DiZhi
// @return string
func DiZhiChineseV2(dizhi DiZhi) string {
	return _DiZhiTable.MustReadString(int(dizhi), 1)
}

// DiZhiChinese returns the chinese DiZhi string
// @param DiZhi
// @return string
// @return error
func DiZhiChinese(dizhi DiZhi) (string, error) {
	readString, err := _DiZhiTable.ReadString(int(dizhi), 1)
	if err != nil {
		return "", ErrWrongDiZhiTypes
	}
	return readString, nil
}

// GanZhiChineseV2 returns the chinese GanZhi string
// @param GanZhi
// @return string
func GanZhiChineseV2(ganzhi GanZhi) string {
	return _GanZhiTable.MustReadString(ganzhi.index(), 2)
}

// GanZhiChinese returns the chinese GanZhi string
// @param GanZhi
// @return string
// @return error
func GanZhiChinese(ganzhi GanZhi) (string, error) {
	readString, err := _GanZhiTable.ReadString(ganzhi.index(), 2)
	if err != nil {
		return "", ErrWrongGanZhiTypes
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
	return GanZhi(fixDayNext(i, idx, h)).Chinese()
}

// RiZhu 获取日柱
func RiZhu(y, m, d int) string {
	return GanZhi(stemBranchIndex(y, m, d)).Chinese()
}

// YueZhuChineseV2 returns the chinese YueZhuChineseV2 string
// @param time.Time
// @return string
func YueZhuChineseV2(t time.Time) string {
	//月柱 1900年1月小寒以前为 丙子月(60进制12)
	return yueZhu(t.Date()).Chinese()
}

func YueZhu(t time.Time) GanZhi {
	return yueZhu(t.Date())
}

func yueZhu(y int, m time.Month, d int) GanZhi {
	_, min := getSolarTermDay(y, m)
	gz := yearOffset(y+1)*12 + int(m)
	if d < min {
		gz -= 1
	}
	gz %= 60
	return GanZhi(gz)
}

// NianZhu returns a GanZhi of year
// @param time.Time
// @return string
func NianZhu(t time.Time) GanZhi {
	return nianZhu(t.Year())
}

// NianZhuChineseV2 returns the chinese NianZhu string
// @param time.Time
// @return string
func NianZhuChineseV2(t time.Time) string {
	return nianZhuChinese(t.Year())
}

func yearOffset(y int) int {
	return y - 1900
}

var _ ChineseSupport = TianGan(0)
var _ ChineseSupport = DiZhi(0)
var _ ChineseSupport = GanZhi(0)

package chronos

import (
	"time"

	"github.com/godcong/chronos/runes"
	"github.com/godcong/chronos/utils"
)

const defaultTianGan = "天干"
const defaultDiZhi = "地支"

// TianGan
// ENUM(Jia , Yi , Bing , Ding , Wu , Ji , Geng , Xin , Ren , Gui , Max)
type TianGan uint32

// DiZhi
// ENUM(Zi , Chou , Yin , Mao , Chen , Si , Wu , Wei , Shen , You , Xu , Hai , Max)
type DiZhi uint32

// GanZhi returns the GanZhi enum
// ENUM(JiaZi,YiChou,BingYin,DingMao,WuChen,JiSi,GengWu,XinWei,RenShen,GuiYou,JiaXu,YiHai,
// BingZi,DingChou,WuYin,JiMao,GengChen,XinSi,RenWu,GuiWei,JiaShen,YiYou,BingXu,DingHai,
// WuZi,JiChou,GengYin,XinMao,RenChen,GuiSi,JiaWu,YiWei,BingShen,DingYou,WuXu,JiHai,
// GengZi,XinChou,RenYin,GuiMao,JiaChen,YiSi,BingWu,DingWei,WuShen,JiYou,GengXu,XinHai,
// RenZi,GuiChou,JiaYin,YiMao,BingChen,DingSi,WuWu,JiWei,GengShen,XinYou,RenXu,GuiHai,Max)
type GanZhi uint32

// StemBranch is an alias name for GanZhi
type StemBranch = GanZhi

var (
	_TianGanTable  = runes.Runes(`甲乙丙丁戊己庚辛壬癸`)
	_DiZhiTable    = runes.Runes(`子丑寅卯辰巳午未申酉戌亥`)
	_TianGanWuXing = runes.Runes("木木火火土土金金水水")
	_DiZhiWuXing   = runes.Runes("水土木木土火火土金金土水")
	_GanZhiTable   = runes.Runes(
		"甲子乙丑丙寅丁卯戊辰己巳庚午辛未壬申癸酉甲戌乙亥" +
			"丙子丁丑戊寅己卯庚辰辛巳壬午癸未甲申乙酉丙戌丁亥" +
			"戊子己丑庚寅辛卯壬辰癸巳甲午乙未丙申丁酉戊戌己亥" +
			"庚子辛丑壬寅癸卯甲辰乙巳丙午丁未戊申己酉庚戌辛亥" +
			"壬子癸丑甲寅乙卯丙辰丁巳戊午己未庚申辛酉壬戌癸亥")
)

// 天干强度表
var _TianGan = [][]int{
	{1200, 1200, 1000, 1000, 1000, 1000, 1000, 1000, 1200, 1200},
	{1060, 1060, 1000, 1000, 1100, 1100, 1140, 1140, 1100, 1100},
	{1140, 1140, 1200, 1200, 1060, 1060, 1000, 1000, 1000, 1000},
	{1200, 1200, 1200, 1200, 1000, 1000, 1000, 1000, 1000, 1000},
	{1100, 1100, 1060, 1060, 1100, 1100, 1100, 1100, 1040, 1040},
	{1000, 1000, 1140, 1140, 1140, 1140, 1060, 1060, 1060, 1060},
	{1000, 1000, 1200, 1200, 1200, 1200, 1000, 1000, 1000, 1000},
	{1040, 1040, 1100, 1100, 1160, 1160, 1100, 1100, 1000, 1000},
	{1060, 1060, 1000, 1000, 1000, 1000, 1140, 1140, 1200, 1200},
	{1000, 1000, 1000, 1000, 1000, 1000, 1200, 1200, 1200, 1200},
	{1000, 1000, 1040, 1040, 1140, 1140, 1160, 1160, 1060, 1060},
	{1200, 1200, 1000, 1000, 1000, 1000, 1000, 1000, 1140, 1140},
}

// 地支强度表
var _DiZhi = []map[string][]int{
	{
		"癸": {1200, 1100, 1000, 1000, 1040, 1060, 1000, 1000, 1200, 1200, 1060, 1140},
	}, {
		"癸": {360, 330, 300, 300, 312, 318, 300, 300, 360, 360, 318, 342},
		"辛": {200, 228, 200, 200, 230, 212, 200, 220, 228, 248, 232, 200},
		"己": {500, 550, 530, 500, 550, 570, 600, 580, 500, 500, 570, 500},
	}, {
		"丙": {300, 300, 360, 360, 318, 342, 360, 330, 300, 300, 342, 318},
		"甲": {840, 742, 798, 840, 770, 700, 700, 728, 742, 700, 700, 840},
	}, {
		"乙": {1200, 1060, 1140, 1200, 1100, 1000, 1000, 1040, 1060, 1000, 1000, 1200},
	}, {
		"乙": {360, 318, 342, 360, 330, 300, 300, 312, 318, 300, 300, 360},
		"癸": {240, 220, 200, 200, 208, 200, 200, 200, 240, 240, 212, 228},
		"戊": {500, 550, 530, 500, 550, 600, 600, 580, 500, 500, 570, 500},
	}, {
		"庚": {300, 342, 300, 300, 330, 300, 300, 330, 342, 360, 348, 300},
		"丙": {700, 700, 840, 840, 742, 840, 840, 798, 700, 700, 728, 742},
	}, {
		"丁": {1000, 1000, 1200, 1200, 1060, 1140, 1200, 1100, 1000, 1000, 1040, 1060},
	}, {
		"丁": {300, 300, 360, 360, 318, 342, 360, 330, 300, 300, 312, 318},
		"乙": {240, 212, 228, 240, 220, 200, 200, 208, 212, 200, 200, 240},
		"己": {500, 550, 530, 500, 550, 570, 600, 580, 500, 500, 570, 500},
	}, {
		"壬": {360, 330, 300, 300, 312, 318, 300, 300, 360, 360, 318, 342},
		"庚": {700, 798, 700, 700, 770, 742, 700, 770, 798, 840, 812, 700},
	}, {
		"辛": {1000, 1140, 1000, 1000, 1100, 1060, 1000, 1100, 1140, 1200, 1160, 1000},
	}, {
		"辛": {300, 342, 300, 300, 330, 318, 300, 330, 342, 360, 348, 300},
		"丁": {200, 200, 240, 240, 212, 228, 240, 220, 200, 200, 208, 212},
		"戊": {500, 550, 530, 500, 550, 570, 600, 580, 500, 500, 570, 500},
	}, {
		"甲": {360, 318, 342, 360, 330, 300, 300, 312, 318, 300, 300, 360},
		"壬": {840, 770, 700, 700, 728, 742, 700, 700, 840, 840, 724, 798},
	},
}

var (
	// PillarHour is an alias name of ShiZhu
	PillarHour = ShiZhu
	// PillarDay is an alias name of RiZhu
	PillarDay = RiZhu
	// PillarMonty is an alias name of YueZhu
	PillarMonty = YueZhu
	// PillarYear is an alias name of NianZhu
	PillarYear = NianZhu
)

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
	gz := int(tiangan)*6 - int(dizhi)*5
	if gz < 0 {
		gz += 60
	}
	gz %= 60
	return GanZhi(gz)
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

// ShiZhu returns a GanZhi of hour
// @param time.Time
// @return GanZhi
// @descriptions
// 子 　　丑 　　寅 　　卯 　　辰 　　己
// 23-01：01-03：03-05 :05-07：07-09：09-11
// 午 　　未 　　申 　　酉 　　戊 　　亥
// 11-13：13-15：15-17：17-19：19-21：21-23
// `甲子`, `乙丑`, `丙寅`, `丁卯`, `戊辰`, `己巳`, `庚午`, `辛未`, `壬申`, `癸酉`, `甲戌`, `乙亥`, //甲或己日
// `丙子`, `丁丑`, `戊寅`, `己卯`, `庚辰`, `辛巳`, `壬午`, `癸未`, `甲申`, `乙酉`, `丙戌`, `丁亥`, //乙或庚日
// `戊子`, `己丑`, `庚寅`, `辛卯`, `壬辰`, `癸巳`, `甲午`, `乙未`, `丙申`, `丁酉`, `戊戌`, `己亥`, //丙或辛日
// `庚子`, `辛丑`, `壬寅`, `癸卯`, `甲辰`, `乙巳`, `丙午`, `丁未`, `戊申`, `己酉`, `庚戌`, `辛亥`, //丁或壬日
// `壬子`, `癸丑`, `甲寅`, `乙卯`, `丙辰`, `丁巳`, `戊午`, `己未`, `庚申`, `辛酉`, `壬戌`, `癸亥`, //戊或癸日
func ShiZhu(t time.Time) GanZhi {
	return shiZhu(t.Year(), t.Month(), t.Day(), t.Hour())
}

func shiZhu(y int, m time.Month, d int, h int) GanZhi {
	days := utils.BetweenDay(TimeFromYmd(y, m, 1), startTime) + d + 9
	zhi := ((h + 1) / 2) % 12
	if h >= 23 {
		days += 1
	}
	gan := (days%10%5)*2 + zhi
	return parseGanZhi(getTianGan(gan), getDiZhi(zhi))
}

// RiZhu returns a GanZhi of day
// @param time.Time
// @return GanZhi
func RiZhu(t time.Time) GanZhi {
	return riZhu(t.Date())
}

func riZhu(y int, m time.Month, d int) GanZhi {
	days := utils.BetweenDay(TimeFromYmd(y, m, 1), startTime) + d + 9
	return parseGanZhi(getTianGan(days), getDiZhi(days))
}

// YueZhuChineseV2 returns the chinese YueZhuChineseV2 string
// @param time.Time
// @return string
func YueZhuChineseV2(t time.Time) string {
	//月柱 1900年1月小寒以前为 丙子月(60进制12)
	return yueZhu(t.Date()).Chinese()
}

// YueZhu returns a GanZhi of month
// @param time.Time
// @return GanZhi
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

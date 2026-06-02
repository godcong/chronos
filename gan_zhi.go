package chronos

import (
	"time"

	"github.com/godcong/chronos/v2/utils"
)

// TianGan represents one of the ten Heavenly Stems (天干): 甲乙丙丁戊己庚辛壬癸.
type TianGan uint32

// DiZhi represents one of the twelve Earthly Branches (地支): 子丑寅卯辰巳午未申酉戌亥.
type DiZhi uint32

// GanZhi represents one of the sixty Stem-Branch combinations (干支) in the
// sexagenary cycle.
type GanZhi uint32
// StemBranch is an alias for GanZhi, representing a Stem-Branch combination.
type StemBranch = GanZhi

var tianGanChinese = [...]string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}

var diZhiChinese = [...]string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

var ganZhiChinese = [...]string{
	"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉",
	"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未",
	"甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳",
	"甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯",
	"甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑",
	"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥",
}

// Removed: _TianGan and _DiZhi variables have been migrated to
// fate/internal/bazi/bazi.go (tiangan, dizhi) for xi-yong-shen calculations.
// These were only used by the fate module's Five Elements scoring algorithm.
// See: SCOPE.md Phase 1 - Dead code cleanup.

var (
	PillarHour  = ShiZhu
	PillarDay   = RiZhu
	PillarMonth = YueZhu
	PillarYear  = NianZhu
)

func (x TianGan) Chinese() string {
	if x >= TianGanMax {
		return ""
	}
	return tianGanChinese[x]
}

func (x DiZhi) Chinese() string {
	if x >= DiZhiMax {
		return ""
	}
	return diZhiChinese[x]
}

func (x GanZhi) Chinese() string {
	if x >= GanZhiMax {
		return ""
	}
	return ganZhiChinese[x]
}

func nianZhu(year int) GanZhi {
	return GanZhi((year - 4) % 60)
}

func nianZhuChinese(year int) string {
	return nianZhu(year).Chinese()
}

func ganZhiStr(gz int) string {
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

// Removed: parseGanZhiV2 used lookup table (_TianGanDiZhiGanZhiTable) from gan_zhi_data.go.
// Replaced by parseGanZhi which uses a math formula producing identical results.
// Tests updated to call parseGanZhi instead.

func parseGanZhi(tiangan TianGan, dizhi DiZhi) GanZhi {
	// Invalid combination: tiangan and dizhi must share the same parity
	// (both even or both odd) to form a valid stem-branch pair.
	if tiangan >= TianGanMax || dizhi >= DiZhiMax {
		return GanZhiMax
	}
	if int(tiangan)%2 != int(dizhi)%2 {
		return GanZhiMax
	}
	gz := int(tiangan)*6 - int(dizhi)*5
	if gz < 0 {
		gz += 60
	}
	gz %= 60
	return GanZhi(gz)
}

// ShiZhu returns the Hour Pillar (时柱) for the given time.
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

// RiZhu returns the Day Pillar (日柱) for the given time.
func RiZhu(t time.Time) GanZhi {
	return riZhu(t.Date())
}

func riZhu(y int, m time.Month, d int) GanZhi {
	days := utils.BetweenDay(TimeFromYmd(y, m, 1), startTime) + d + 9
	return parseGanZhi(getTianGan(days), getDiZhi(days))
}

// YueZhu returns the Month Pillar (月柱) for the given time.
func YueZhu(t time.Time) GanZhi {
	return yueZhu(t.Date())
}

func yueZhu(y int, m time.Month, d int) GanZhi {
	min, _ := getSolarTermDay(y, m)
	gz := yearOffset(y+1)*12 + int(m)
	if d < min {
		gz -= 1
	}
	gz %= 60
	return GanZhi(gz)
}

// NianZhu returns the Year Pillar (年柱) for the given time.
func NianZhu(t time.Time) GanZhi {
	return nianZhu(t.Year())
}

func yearOffset(y int) int {
	return y - 1900
}

var _ ChineseSupport = TianGan(0)
var _ ChineseSupport = DiZhi(0)
var _ ChineseSupport = GanZhi(0)

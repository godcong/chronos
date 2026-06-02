package chronos

import (
	"testing"
	"time"
)

func TestNianZhuChinese(t *testing.T) {
	tests := []struct {
		y    time.Time
		want string
	}{
		{TimeFromY(1900), "庚子"},
		{TimeFromY(1899), "己亥"},
		{TimeFromY(2099), "己未"},
		{TimeFromY(2100), "庚申"},
	}
	for _, tt := range tests {
		got := nianZhu(tt.y.Year()).Chinese()
		if got != tt.want {
			t.Errorf("nianZhu(%d).Chinese() = %v, want %v", tt.y.Year(), got, tt.want)
		}
	}
}

func TestGanZhiChinese(t *testing.T) {
	tests := []struct {
		gz       GanZhi
		want     string
		wantEmpty bool
	}{
		{0, "甲子", false},
		{59, "癸亥", false},
		{60, "", true},
	}
	for _, tt := range tests {
		got := tt.gz.Chinese()
		if (got == "") != tt.wantEmpty {
			t.Errorf("GanZhi(%d).Chinese() = %q, wantEmpty %v", tt.gz, got, tt.wantEmpty)
		}
		if got != tt.want {
			t.Errorf("GanZhi(%d).Chinese() = %v, want %v", tt.gz, got, tt.want)
		}
	}
}

func TestYueZhu(t *testing.T) {
	tests := []struct {
		t    time.Time
		want string
	}{
		{TimeFromYmd(1900, 1, 1), "丙子"},
	}
	for _, tt := range tests {
		got := YueZhu(tt.t).Chinese()
		if got != tt.want {
			t.Errorf("YueZhu() = %v, want %v", got, tt.want)
		}
	}
}

func Test_monthGanZhi(t *testing.T) {
	tests := []struct {
		year  int
		month time.Month
		day   int
		want  string
	}{
		{1900, 1, 1, "丙子"},
		{1900, 11, 3, "丙戌"},
		{2099, 11, 3, "甲戌"},
	}
	for _, tt := range tests {
		got := yueZhu(tt.year, tt.month, tt.day).Chinese()
		if got != tt.want {
			t.Errorf("yueZhu() = %v, want %v", got, tt.want)
		}
	}
}

func Test_splitGanZhi(t *testing.T) {
	got, got1 := splitGanZhi(GanZhiJiaWu)
	if got != TianGanJia || got1 != DiZhiWu {
		t.Errorf("splitGanZhi() = (%v, %v), want (%v, %v)", got, got1, TianGanJia, DiZhiWu)
	}
}

func Test_parseGanZhi(t *testing.T) {
	tests := []struct {
		tiangan TianGan
		dizhi   DiZhi
		want    GanZhi
	}{
		{TianGanJia, 0, 0},
		{TianGanJia, DiZhiWu, GanZhiJiaWu},
		{TianGanJia, DiZhiChou, GanZhiMax},
		{TianGanYi, DiZhiHai, GanZhiYiHai},
		{TianGanGeng, DiZhiHai, GanZhiMax},
	}
	for _, tt := range tests {
		got := parseGanZhi(tt.tiangan, tt.dizhi)
		if got != tt.want {
			t.Errorf("parseGanZhi() got = %v, want %v", got, tt.want)
		}
	}
}

func TestShiZhu(t *testing.T) {
	tests := []struct {
		t    time.Time
		want string
	}{
		{time.Date(2022, 9, 5, 14, 42, 30, 0, time.Local), "乙未"},
		{time.Date(2022, 9, 5, 23, 42, 30, 0, time.Local), "庚子"},
		{time.Date(2022, 9, 6, 0, 42, 30, 0, time.Local), "庚子"},
		{time.Date(2022, 10, 1, 0, 42, 30, 0, time.Local), "庚子"},
	}
	for _, tt := range tests {
		got := ShiZhu(tt.t).Chinese()
		if got != tt.want {
			t.Errorf("ShiZhu() = %v, want %v", got, tt.want)
		}
	}
}

func TestNianZhu(t *testing.T) {
	tests := []struct {
		year int
		want string
	}{
		{2024, "甲辰"},
		{1990, "庚午"},
		{2000, "庚辰"},
		{1900, "庚子"},
	}
	for _, tt := range tests {
		got := NianZhu(time.Date(tt.year, 1, 1, 0, 0, 0, 0, loc))
		if got.Chinese() != tt.want {
			t.Errorf("NianZhu(%d) = %s, want %s", tt.year, got.Chinese(), tt.want)
		}
	}
}

func TestYueZhuGanZhi(t *testing.T) {
	tests := []struct {
		year  int
		month time.Month
		day   int
		want  string
	}{
		{2024, 2, 5, "丙寅"},
		{2024, 1, 5, "甲子"},
	}
	for _, tt := range tests {
		got := YueZhu(time.Date(tt.year, tt.month, tt.day, 12, 0, 0, 0, loc))
		if got.Chinese() != tt.want {
			t.Errorf("YueZhu(%d-%d-%d) = %s, want %s", tt.year, tt.month, tt.day, got.Chinese(), tt.want)
		}
	}
}

func TestRiZhuGanZhi(t *testing.T) {
	got := RiZhu(time.Date(2024, 2, 5, 12, 0, 0, 0, loc))
	if got.Chinese() == "" {
		t.Error("RiZhu should not be empty")
	}
}

func TestShiZhuGanZhi(t *testing.T) {
	got := ShiZhu(time.Date(2024, 2, 5, 12, 0, 0, 0, loc))
	if got.Chinese() == "" {
		t.Error("ShiZhu should not be empty")
	}
}

func TestGetTianGan(t *testing.T) {
	tests := []struct {
		v    int
		want string
	}{
		{0, "甲"}, {1, "乙"}, {4, "戊"}, {9, "癸"},
	}
	for _, tt := range tests {
		got := getTianGan(tt.v)
		if got.Chinese() != tt.want {
			t.Errorf("getTianGan(%d) = %s, want %s", tt.v, got.Chinese(), tt.want)
		}
	}
}

func TestGetDiZhi(t *testing.T) {
	tests := []struct {
		v    int
		want string
	}{
		{0, "子"}, {1, "丑"}, {11, "亥"},
	}
	for _, tt := range tests {
		got := getDiZhi(tt.v)
		if got.Chinese() != tt.want {
			t.Errorf("getDiZhi(%d) = %s, want %s", tt.v, got.Chinese(), tt.want)
		}
	}
}

func TestParseGanZhi(t *testing.T) {
	gz := parseGanZhi(TianGan(0), DiZhi(0))
	if gz != GanZhi(0) {
		t.Errorf("parseGanZhi(甲,子) = %d, want 0", gz)
	}
}

func TestSplitGanZhi(t *testing.T) {
	tg, dz := splitGanZhi(GanZhi(0))
	if tg != TianGan(0) || dz != DiZhi(0) {
		t.Errorf("splitGanZhi(甲子) = (%d, %d), want (0, 0)", tg, dz)
	}
}

func TestRiZhuDate(t *testing.T) {
	tests := []struct {
		t    time.Time
		want string
	}{
		{TimeFromYmd(2022, 9, 5), "辛酉"},
		{TimeFromYmd(2099, 11, 4), "乙巳"},
	}
	for _, tt := range tests {
		got := RiZhu(tt.t).Chinese()
		if got != tt.want {
			t.Errorf("RiZhu() = %v, want %v", got, tt.want)
		}
	}
}

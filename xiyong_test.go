package chronos

import (
	"testing"
)

func TestGetWuxingOfTianGan(t *testing.T) {
	tests := []struct {
		gan  string
		want string
	}{
		{"甲", "木"}, {"乙", "木"}, {"丙", "火"}, {"丁", "火"},
		{"戊", "土"}, {"己", "土"}, {"庚", "金"}, {"辛", "金"},
		{"壬", "水"}, {"癸", "水"}, {"?", ""},
	}
	for _, tt := range tests {
		got := getWuxingOfTianGan(tt.gan)
		if got != tt.want {
			t.Errorf("getWuxingOfTianGan(%s) = %s, want %s", tt.gan, got, tt.want)
		}
	}
}

func TestGetTongleiWuxing(t *testing.T) {
	tests := []struct {
		wx   string
		want []string
	}{
		{"木", []string{"木", "水"}},
		{"火", []string{"火", "木"}},
		{"土", []string{"土", "火"}},
		{"金", []string{"金", "土"}},
		{"水", []string{"水", "金"}},
	}
	for _, tt := range tests {
		got := getTongleiWuxing(tt.wx)
		if len(got) != len(tt.want) {
			t.Errorf("getTongleiWuxing(%s) = %v, want %v", tt.wx, got, tt.want)
			continue
		}
		for i, v := range got {
			if v != tt.want[i] {
				t.Errorf("getTongleiWuxing(%s)[%d] = %s, want %s", tt.wx, i, v, tt.want[i])
			}
		}
	}
}

func TestFindXianWuxing(t *testing.T) {
	xyj := &XiYongJiChou{
		YongWuxing: "金",
		XiWuxing:   []string{"水"},
		JiWuxing:   []string{"木"},
		ChouWuxing: []string{"火"},
	}
	xian := findXianWuxing(xyj)
	if len(xian) != 1 || xian[0] != "土" {
		t.Errorf("findXianWuxing = %v, want [土]", xian)
	}
}

func TestDetermineGeJu(t *testing.T) {
	geJu := determineGeJu("甲", "寅", [4]string{"甲子", "丙寅", "甲子", "甲子"})
	if geJu == nil {
		t.Fatal("expected non-nil GeJu")
	}
	if geJu.Type == GeJuSpecial && geJu.ShiShen == "" {
		t.Error("expected non-special geju for valid input")
	}
}

func TestGetShiShen(t *testing.T) {
	tests := []struct {
		dayGan    string
		targetGan string
		want      string
	}{
		{"甲", "甲", "比肩"},
		{"甲", "庚", "七杀"},
		{"甲", "己", "正财"},
		{"丙", "壬", "七杀"},
	}
	for _, tt := range tests {
		got := getShiShen(tt.dayGan, tt.targetGan)
		if got != tt.want {
			t.Errorf("getShiShen(%s, %s) = %s, want %s", tt.dayGan, tt.targetGan, got, tt.want)
		}
	}
}

func TestGeJuType_String(t *testing.T) {
	tests := []struct {
		gt   GeJuType
		want string
	}{
		{GeJuZhengGuan, "正官格"},
		{GeJuQiSha, "七杀格"},
		{GeJuZhengYin, "正印格"},
		{GeJuShangGuan, "伤官格"},
	}
	for _, tt := range tests {
		got := tt.gt.String()
		if got != tt.want {
			t.Errorf("GeJuType(%d).String() = %s, want %s", tt.gt, got, tt.want)
		}
	}
}

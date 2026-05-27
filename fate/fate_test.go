package fate

import (
	"testing"

	"github.com/godcong/chronos/v2"
)

func TestGetFateData(t *testing.T) {
	birthDate := chronos.ParseSolarTime(chronos.TimeFromYmdHms(1990, 6, 15, 12, 0, 0))

	t.Run("nil input returns error", func(t *testing.T) {
		_, err := GetFateData(FateInput{Calendar: nil})
		if err == nil {
			t.Error("expected error for nil calendar")
		}
	})

	t.Run("valid input returns fate data with balance method", func(t *testing.T) {
		input := FateInput{
			Calendar:     birthDate,
			Gender:       1,
			XiYongMethod: XiYongMethodBalance,
		}
		data, err := GetFateData(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(data.BaziInfo.SiZhu) != 4 {
			t.Errorf("expected 4 pillars, got %d", len(data.BaziInfo.SiZhu))
		}
		if data.WuxingXijiInfo.Xi == "" {
			t.Error("expected non-empty Xi")
		}
	})

	t.Run("valid input returns fate data with geju method", func(t *testing.T) {
		input := FateInput{
			Calendar:     birthDate,
			Gender:       1,
			XiYongMethod: XiYongMethodGeJu,
		}
		data, err := GetFateData(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if data.GeJuInfo == nil {
			t.Error("expected non-nil GeJuInfo for geju method")
		}
	})

	t.Run("default method is balance", func(t *testing.T) {
		input := FateInput{
			Calendar: birthDate,
			Gender:   1,
		}
		data, err := GetFateData(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if data.XiYongJiChou.Yong == "" {
			t.Error("expected non-empty Yong")
		}
	})
}

func TestWuxingOfTianGan(t *testing.T) {
	tests := []struct {
		gan  string
		want string
	}{
		{"甲", "木"}, {"乙", "木"}, {"丙", "火"}, {"丁", "火"},
		{"戊", "土"}, {"己", "土"}, {"庚", "金"}, {"辛", "金"},
		{"壬", "水"}, {"癸", "水"}, {"?", ""},
	}
	for _, tt := range tests {
		got := wuxingOfTianGan(tt.gan)
		if got != tt.want {
			t.Errorf("wuxingOfTianGan(%s) = %s, want %s", tt.gan, got, tt.want)
		}
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

func TestGeJuTypeName(t *testing.T) {
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
		got := geJuTypeName(tt.gt)
		if got != tt.want {
			t.Errorf("geJuTypeName(%d) = %s, want %s", tt.gt, got, tt.want)
		}
	}
}

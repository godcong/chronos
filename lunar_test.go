package chronos

import (
	"reflect"
	"testing"
)

func Test_lunar_GetZodiac(t *testing.T) {
	tests := []struct {
		name   string
		lunar  Lunar
		want   Zodiac
	}{
		{"", NewSolarCalendar(TimeFromYmdHms(2023, 2, 4, 10, 42, 21)).Lunar(), ZodiacRabbit},
		{"", NewSolarCalendar(TimeFromYmdHms(2023, 2, 4, 0, 0, 0)).Lunar(), ZodiacRabbit},
		{"", NewSolarCalendar(TimeFromYmdHms(2023, 2, 3, 10, 42, 20)).Lunar(), ZodiacTiger},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lunar.GetZodiac(); got != tt.want {
				t.Errorf("GetZodiac() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lunar_GetSolarTerm(t *testing.T) {
	tests := []struct {
		name   string
		lunar  Lunar
		want   SolarTerm
	}{
		{"", NewSolarCalendar(TimeFromYmd(2023, 2, 4)).Lunar(), 2},
		{"", NewSolarCalendar(TimeFromYmd(2023, 2, 5)).Lunar(), 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lunar.GetSolarTerm(); got != tt.want {
				t.Errorf("GetSolarTerm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lunar_GetEightChar(t *testing.T) {
	tests := []struct {
		name            string
		lunar           Lunar
		want            [4]string
		wantWuXing      [4]string
		wantNayin       [4]string
		wantShiShengGan [4]string
		wantShiShengZhi [4][]string
		wantCangGan     [4][]string
		wantDaYun       []int
	}{
		{
			name:            "",
			lunar:           NewSolarCalendar(TimeFromYmdHms(2023, 2, 5, 12, 0, 0)).Lunar(),
			want:            [4]string{"癸卯", "甲寅", "甲午", "庚午"},
			wantWuXing:      [4]string{"水木", "木木", "木火", "金火"},
			wantNayin:       [4]string{"金箔金", "大溪水", "沙中金", "路旁土"},
			wantShiShengGan: [4]string{"正印", "比肩", "日主", "七杀"},
			wantShiShengZhi: [4][]string{{"劫财"}, {"比肩", "食神", "偏财"}, {"伤官", "正财"}, {"伤官", "正财"}},
			wantCangGan:     [4][]string{{"乙"}, {"甲", "丙", "戊"}, {"丁", "己"}, {"丁", "己"}},
			wantDaYun:       []int{2023, 2033, 2043, 2053, 2063, 2073, 2083, 2093, 2103, 2113},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := tt.lunar.GetEightChar()
			if got := ec.FourPillars(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FourPillars() = %v, want %v", got, tt.want)
			}
			if got := ec.NaYin(); !reflect.DeepEqual(got, tt.wantNayin) {
				t.Errorf("NaYin() = %v, want %v", got, tt.wantNayin)
			}
			if got := ec.FiveElements(); !reflect.DeepEqual(got, tt.wantWuXing) {
				t.Errorf("FiveElements() = %v, want %v", got, tt.wantWuXing)
			}
			if got := ec.TenGodsStems(); !reflect.DeepEqual(got, tt.wantShiShengGan) {
				t.Errorf("TenGodsStems() = %v, want %v", got, tt.wantShiShengGan)
			}
			if got := ec.TenGodsBranches(); !reflect.DeepEqual(got, tt.wantShiShengZhi) {
				t.Errorf("TenGodsBranches() = %v, want %v", got, tt.wantShiShengZhi)
			}
			if got := ec.HiddenStems(); !reflect.DeepEqual(got, tt.wantCangGan) {
				t.Errorf("HiddenStems() = %v, want %v", got, tt.wantCangGan)
			}
			if got := ec.DaYun(1); !reflect.DeepEqual(got, tt.wantDaYun) {
				t.Errorf("DaYun() = %v, want %v", got, tt.wantDaYun)
			}
		})
	}
}

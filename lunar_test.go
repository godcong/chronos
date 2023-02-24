package chronos

import (
	"reflect"
	"testing"
)

func Test_lunar_GetZodiac(t *testing.T) {
	type fields struct {
		Lunar Lunar
	}
	tests := []struct {
		name   string
		fields fields
		want   Zodiac
	}{
		{
			name: "",
			fields: fields{
				Lunar: NewSolarCalendar(TimeFromYmdHms(2023, 2, 4, 10, 42, 21)).Lunar(),
			},
			want: ZodiacRabbit,
		},
		{
			name: "",
			fields: fields{
				Lunar: NewSolarCalendar(TimeFromYmdHms(2023, 2, 4, 10, 42, 20)).Lunar(),
			},
			want: ZodiacTiger,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.fields.Lunar
			if got := l.GetZodiac(); got != tt.want {
				t.Errorf("GetZodiac() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lunar_GetSolarTerm(t *testing.T) {
	type fields struct {
		Lunar Lunar
	}
	tests := []struct {
		name   string
		fields fields
		want   SolarTerm
	}{
		{
			name: "",
			fields: fields{
				Lunar: NewSolarCalendar(TimeFromYmd(2023, 2, 4)).Lunar(),
			},
			want: 2,
		},
		{
			name: "",
			fields: fields{
				Lunar: NewSolarCalendar(TimeFromYmd(2023, 2, 5)).Lunar(),
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.fields.Lunar
			if got := l.GetSolarTerm(); got != tt.want {
				t.Errorf("GetSolarTerm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lunar_GetEightChar(t *testing.T) {
	type fields struct {
		Lunar Lunar
	}
	tests := []struct {
		name            string
		fields          fields
		want            [4]string
		wantWuXing      [4]string
		wantNayin       [4]string
		wantShiShengGan [4]string
		wantShiShengZhi [4][]string
		wantCangGan     [4][]string
		wantDaYun       []int
	}{
		{
			name: "",
			fields: fields{
				Lunar: NewSolarCalendar(TimeFromYmdHms(2023, 2, 5, 12, 0, 0)).Lunar(),
			},
			want:            [4]string{"癸卯", "甲寅", "甲午", "庚午"},
			wantWuXing:      [4]string{"水木", "木木", "木火", "金火"},
			wantNayin:       [4]string{"金箔金", "大溪水", "沙中金", "路旁土"},
			wantShiShengGan: [4]string{"正印", "比肩", "日主", "七杀"}, //偏官(七杀)
			wantShiShengZhi: [4][]string{{"劫财"}, {"比肩", "食神", "偏财"}, {"伤官", "正财"}, {"伤官", "正财"}},
			wantCangGan:     [4][]string{{"乙"}, {"甲", "丙", "戊"}, {"丁", "己"}, {"丁", "己"}},
			wantDaYun:       []int{2023, 2033, 2043, 2053, 2063, 2073, 2083, 2093, 2103, 2113},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.fields.Lunar
			if got := l.GetEightChar(); !reflect.DeepEqual(got.GetSiZhu(), tt.want) {
				t.Errorf("GetSiZhu() = %v, want %v", got.GetSiZhu(), tt.want)
			}
			if got := l.GetEightChar(); !reflect.DeepEqual(got.GetNaYin(), tt.wantNayin) {
				t.Errorf("GetNaYin() = %v, want %v", got.GetNaYin(), tt.wantNayin)
			}
			if got := l.GetEightChar(); !reflect.DeepEqual(got.GetWuXing(), tt.wantWuXing) {
				t.Errorf("GetWuXing() = %v, want %v", got.GetWuXing(), tt.wantWuXing)
			}
			if got := l.GetEightChar(); !reflect.DeepEqual(got.GetShiShenGan(), tt.wantShiShengGan) {
				t.Errorf("GetShiShenGan() = %v, want %v", got.GetShiShenGan(), tt.wantShiShengGan)
			}
			if got := l.GetEightChar(); !reflect.DeepEqual(got.GetShiShenZhi(), tt.wantShiShengZhi) {
				t.Errorf("GetShiShenZhi() = %v, want %v", got.GetShiShenZhi(), tt.wantShiShengZhi)
			}
			if got := l.GetEightChar(); !reflect.DeepEqual(got.GetCangGan(), tt.wantCangGan) {
				t.Errorf("GetCangGan() = %v, want %v", got.GetCangGan(), tt.wantCangGan)
			}
			if got := l.GetEightChar(); !reflect.DeepEqual(got.GetDaYun(1), tt.wantDaYun) {
				t.Errorf("GetDaYun() = %v, want %v", got.GetDaYun(1), tt.wantDaYun)
			}
		})
	}
}

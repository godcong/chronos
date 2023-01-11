package chronos

import (
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

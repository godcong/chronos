package chronos

import (
	"testing"
)

func Test_solar_GetConstellation(t *testing.T) {
	type fields struct {
		Solar Solar
	}
	tests := []struct {
		name   string
		fields fields
		want   Constellation
	}{
		{
			name: "",
			fields: fields{
				Solar: NewSolarCalendar(TimeFromYmd(2023, 1, 19)).Solar(),
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				Solar: NewSolarCalendar(TimeFromYmd(2023, 1, 20)).Solar(),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields.Solar
			if got := s.GetConstellation(); got != tt.want {
				t.Errorf("Constellation() = %v, want %v", got, tt.want)
			}
		})
	}
}

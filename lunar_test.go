package chronos

import (
	"testing"
	"time"
)

func Test_betweenDay(t *testing.T) {
	type args struct {
		d time.Time
		s time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				d: yearMonthDayDate(1900, 1, 3),
				s: yearMonthDayDate(1900, 1, 2),
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				d: yearMonthDayDate(1900, 1, 3).Add(24 * time.Hour),
				s: yearMonthDayDate(1900, 1, 2),
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				d: yearMonthDayDate(1900, 1, 3).Add(23*time.Hour + 59*time.Minute + 59*time.Second),
				s: yearMonthDayDate(1900, 1, 2),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := betweenDay(tt.args.d, tt.args.s); got != tt.want {
				t.Errorf("betweenDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

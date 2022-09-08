package chronos

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/godcong/chronos/v2/utils"
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
				d: TimeFromYmd(1900, 1, 3),
				s: TimeFromYmd(1900, 1, 2),
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				d: TimeFromYmd(1900, 1, 3).Add(24 * time.Hour),
				s: TimeFromYmd(1900, 1, 2),
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				d: TimeFromYmd(1900, 1, 3).Add(23*time.Hour + 59*time.Minute + 59*time.Second),
				s: TimeFromYmd(1900, 1, 2),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.BetweenDay(tt.args.d, tt.args.s); got != tt.want {
				t.Errorf("betweenDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lunarYear(t *testing.T) {
	type args struct {
		offset int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "",
			args: args{
				offset: 1900,
			},
			want:  0,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//got, got1 := lunarYear(tt.args.offset)
			//if got != tt.want {
			//	date.Errorf("lunarYear() got = %v, want %v", got, tt.want)
			//}
			//if got1 != tt.want1 {
			//	date.Errorf("lunarYear() got1 = %v, want %v", got1, tt.want1)
			//}
		})
	}
}

func Test_yearDay(t *testing.T) {
	type args struct {
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				y: 1900,
			},
			want: 384,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := yearDay(tt.args.y)
			if got != tt.want {
				t.Errorf("yearDay() = %v, want %v", got, tt.want)
			}
			//got2 := yearDayOld(tt.args.y)
			//if got != got2 {
			//	date.Errorf("yearDay() = %v, want %v", got2, tt.want)
			//}
		})
	}
}

func TestParseLunarTime(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				date: TimeFromYmd(1900, 2, 4),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseLunarTime(tt.args.date)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLunarTime() = %v, want %v", got, tt.want)
			}
			fmt.Printf("date:%+v\n", got.Date())
		})
	}
}

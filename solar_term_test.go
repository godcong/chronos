package chronos

import (
	"reflect"
	"testing"
	"time"
)

func TestSolarTermChinese(t *testing.T) {
	type args struct {
		i SolarTerm
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				i: 0,
			},
			want:    "小寒",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				i: 1,
			},
			want:    "大寒",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				i: 2,
			},
			want:    "立春",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				i: 3,
			},
			want:    "雨水",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				i: 4,
			},
			want:    "惊蛰",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				i: 5,
			},
			want:    "春分",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				i: 6,
			},
			want:    "清明",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				i: 7,
			},
			want:    "谷雨",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				i: 8,
			},
			want:    "立夏",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				i: 23,
			},
			want:    "冬至",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				i: 24,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolarTermChinese(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolarTermChinese() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolarTermChinese() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSolarTermTime(t *testing.T) {
	type args struct {
		year int
		st   SolarTerm
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				year: 1900,
				st:   23,
			},
			want: "1900-01-20 19:32:25",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getYearSolarTermTime(tt.args.year, tt.args.st); !reflect.DeepEqual(got.Format("2006-01-02 15:04:05"), tt.want) {
				t.Errorf("getYearSolarTermTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYearSolarTermDetail(t *testing.T) {
	type args struct {
		year time.Time
		st   SolarTerm
	}
	tests := []struct {
		name string
		args args
		want SolarTermDetail
	}{
		{
			name: "",
			args: args{
				year: yearDate(1900),
				st:   0,
			},
			want: SolarTermDetail{
				Index:       0,
				SolarTerm:   0,
				Time:        "1900/02/04 13:51:31",
				SanHou:      "",
				Explanation: "",
			},
		},
		{
			name: "",
			args: args{
				year: yearDate(1900),
				st:   23,
			},
			want: SolarTermDetail{
				Index:       0,
				SolarTerm:   0,
				Time:        "1900/01/20 19:32:25",
				SanHou:      "",
				Explanation: "",
			},
		},
		{
			name: "",
			args: args{
				year: yearDate(1900),
				st:   24,
			},
			want: SolarTermDetail{
				Index:       0,
				SolarTerm:   0,
				Time:        "",
				SanHou:      "",
				Explanation: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := YearSolarTermDetail(tt.args.year, tt.args.st); !reflect.DeepEqual(got.Time, tt.want.Time) {
				t.Errorf("YearSolarTermDetail() = %v, want %v", got.Time, tt.want.Time)
			}
		})
	}
}

func TestIsSolarTermDetailDay(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				t: time.Date(1900, 01, 20, 19, 32, 25, 0, time.UTC),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSolarTermDay(tt.args.t); got != tt.want {
				t.Errorf("IsSolarTermDetailDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yearLiChunDay(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name    string
		args    args
		wantDay int
	}{
		{
			name: "",
			args: args{
				year: 1900,
			},
			wantDay: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDay := yearLiChunDay(tt.args.year); gotDay != tt.wantDay {
				t.Errorf("yearLiChunDay() = %v, want %v", gotDay, tt.wantDay)
			}
		})
	}
}

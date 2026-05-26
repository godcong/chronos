package chronos

import (
	"testing"
	"time"
)

func TestConstellationChinese(t *testing.T) {
	type args struct {
		c Constellation
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				c: 0,
			},
			want:    "摩羯",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConstellationChinese(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConstellationChinese() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConstellationChinese() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getConstellation2(t *testing.T) {
	type args struct {
		month time.Month
		day   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				month: 1,
				day:   19,
			},
			want: "摩羯",
		},
		{
			name: "",
			args: args{
				month: 1,
				day:   20,
			},
			want: "水瓶",
		},
		{
			name: "",
			args: args{
				month: 2,
				day:   18,
			},
			want: "水瓶",
		},
		{
			name: "",
			args: args{
				month: 2,
				day:   19,
			},
			want: "双鱼",
		},
		{
			name: "",
			args: args{
				month: 3,
				day:   20,
			},
			want: "双鱼",
		},
		{
			name: "",
			args: args{
				month: 3,
				day:   21,
			},
			want: "白羊",
		},
		{
			name: "",
			args: args{
				month: 4,
				day:   19,
			},
			want: "白羊",
		},
		{
			name: "",
			args: args{
				month: 4,
				day:   20,
			},
			want: "金牛",
		},
		{
			name: "",
			args: args{
				month: 5,
				day:   19,
			},
			want: "金牛",
		},
		{
			name: "",
			args: args{
				month: 5,
				day:   21,
			},
			want: "双子",
		},
		{
			name: "",
			args: args{
				month: 10,
				day:   23,
			},
			want: "天秤",
		},
		{
			name: "",
			args: args{
				month: 10,
				day:   24,
			},
			want: "天蝎",
		},
		{
			name: "",
			args: args{
				month: 11,
				day:   22,
			},
			want: "天蝎",
		},
		{
			name: "",
			args: args{
				month: 11,
				day:   23,
			},
			want: "射手",
		},
		{
			name: "",
			args: args{
				month: 12,
				day:   21,
			},
			want: "射手",
		},
		{
			name: "",
			args: args{
				month: 12,
				day:   22,
			},
			want: "摩羯",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConstellation(0, tt.args.month, tt.args.day); got.Chinese() != tt.want {
				t.Errorf("getConstellation2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetConstellation(t *testing.T) {
	tests := []struct {
		month time.Month
		day   int
		want  string
	}{
		{1, 15, "摩羯"},
		{2, 15, "水瓶"},
		{3, 15, "双鱼"},
		{4, 15, "白羊"},
		{5, 15, "金牛"},
		{6, 15, "双子"},
		{7, 15, "巨蟹"},
		{8, 15, "狮子"},
		{9, 15, "处女"},
		{10, 15, "天秤"},
		{11, 15, "天蝎"},
		{12, 15, "射手"},
	}
	for _, tt := range tests {
		got := GetConstellation(time.Date(2024, tt.month, tt.day, 0, 0, 0, 0, loc))
		if got.Chinese() != tt.want {
			t.Errorf("GetConstellation(%d-%d) = %s, want %s", tt.month, tt.day, got.Chinese(), tt.want)
		}
	}
}

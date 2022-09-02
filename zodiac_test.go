package chronos

import (
	"testing"
	"time"
)

func TestZodiacChinese(t *testing.T) {
	type args struct {
		zodiac Zodiac
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				zodiac: 0,
			},
			want: "鼠",
		},
		{
			name: "",
			args: args{
				zodiac: 1,
			},
			want: "牛",
		},
		{
			name: "",
			args: args{
				zodiac: 2,
			},
			want: "虎",
		},
		{
			name: "",
			args: args{
				zodiac: 3,
			},
			want: "兔",
		},
		{
			name: "",
			args: args{
				zodiac: 4,
			},
			want: "龙",
		},
		{
			name: "",
			args: args{
				zodiac: 5,
			},
			want: "蛇",
		},
		{
			name: "",
			args: args{
				zodiac: 6,
			},
			want: "马",
		},
		{
			name: "",
			args: args{
				zodiac: 7,
			},
			want: "羊",
		},
		{
			name: "",
			args: args{
				zodiac: 8,
			},
			want: "猴",
		},
		{
			name: "",
			args: args{
				zodiac: 9,
			},
			want: "鸡",
		},
		{
			name: "",
			args: args{
				zodiac: 10,
			},
			want: "狗",
		},
		{
			name: "",
			args: args{
				zodiac: 11,
			},
			want: "猪",
		},
		{
			name: "",
			args: args{
				zodiac: 12,
			},
			want: "猫",
		},
		{
			name: "",
			args: args{
				zodiac: 13,
			},
			want: "猫",
		},
		{
			name: "",
			args: args{
				zodiac: 14,
			},
			want: "猫",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZodiacChineseV2(tt.args.zodiac); got != tt.want {
				t.Errorf("ZodiacChinese() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getZodiac(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want Zodiac
	}{
		{
			name: "",
			args: args{
				year: 1899,
			},
			want: 11,
		},
		{
			name: "",
			args: args{
				year: 1900,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				year: 1901,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getZodiac(tt.args.year); got != tt.want {
				t.Errorf("getZodiac() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetYearZodiac(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    Zodiac
		wantErr bool
	}{
		{
			name: "",
			args: args{
				t: time.Date(),
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetYearZodiac(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetYearZodiac() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetYearZodiac() got = %v, want %v", got, tt.want)
			}
		})
	}
}

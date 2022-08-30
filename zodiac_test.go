package chronos

import (
	"testing"
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
				zodiac: ZodiacCat,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZodiacChinese(tt.args.zodiac); got != tt.want {
				t.Errorf("ZodiacChinese() = %v, want %v", got, tt.want)
			}
		})
	}
}

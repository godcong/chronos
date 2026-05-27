package chronos

import (
	"testing"
	"time"
)

func TestConstellationChinese(t *testing.T) {
	tests := []struct {
		c    Constellation
		want string
	}{
		{0, "摩羯"},
		{1, "水瓶"},
		{11, "射手"},
		{12, ""},
	}
	for _, tt := range tests {
		got := tt.c.Chinese()
		if got != tt.want {
			t.Errorf("Constellation(%d).Chinese() = %s, want %s", tt.c, got, tt.want)
		}
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
		{"", args{1, 19}, "摩羯"},
		{"", args{1, 20}, "水瓶"},
		{"", args{2, 18}, "水瓶"},
		{"", args{2, 19}, "双鱼"},
		{"", args{3, 20}, "双鱼"},
		{"", args{3, 21}, "白羊"},
		{"", args{4, 19}, "白羊"},
		{"", args{4, 20}, "金牛"},
		{"", args{5, 19}, "金牛"},
		{"", args{5, 21}, "双子"},
		{"", args{10, 23}, "天秤"},
		{"", args{10, 24}, "天蝎"},
		{"", args{11, 22}, "天蝎"},
		{"", args{11, 23}, "射手"},
		{"", args{12, 21}, "射手"},
		{"", args{12, 22}, "摩羯"},
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

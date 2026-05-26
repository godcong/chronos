package chronos

import (
	"testing"
)

func TestGetChineseMonth(t *testing.T) {
	tests := []struct {
		m    int
		want string
	}{
		{1, "正月"}, {2, "二月"}, {12, "腊月"},
	}
	for _, tt := range tests {
		got := getChineseMonth(tt.m)
		if got != tt.want {
			t.Errorf("getChineseMonth(%d) = %s, want %s", tt.m, got, tt.want)
		}
	}
}

func TestGetChineseDay(t *testing.T) {
	tests := []struct {
		d    int
		want string
	}{
		{1, "初一日"}, {10, "初十日"}, {15, "十五日"}, {20, "二十日"}, {30, "三十日"},
	}
	for _, tt := range tests {
		got := getChineseDay(tt.d)
		if got != tt.want {
			t.Errorf("getChineseDay(%d) = %s, want %s", tt.d, got, tt.want)
		}
	}
}

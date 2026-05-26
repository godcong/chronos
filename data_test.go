package chronos

import (
	"testing"
	"time"
)

func Test_yearLeapMonth(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				year: 1900,
			},
			want: 8,
		},
		{
			name: "",
			args: args{
				year: 2997,
			},
			want: 8,
		},
		{
			name: "",
			args: args{
				year: 2883,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := yearLeapMonth(tt.args.year); got != tt.want {
				t.Errorf("yearLeapMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeapMonth(t *testing.T) {
	lm, err := LeapMonth(time.Date(2023, 1, 1, 0, 0, 0, 0, loc))
	if err != nil {
		t.Fatal(err)
	}
	if lm != 2 {
		t.Errorf("2023 leap month = %d, want 2", lm)
	}
}

func TestLeapMonth_NoLeap(t *testing.T) {
	_, err := LeapMonth(time.Date(2024, 1, 1, 0, 0, 0, 0, loc))
	if err == nil {
		t.Error("2024 should not have leap month")
	}
}

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

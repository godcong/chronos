package chronos

import (
	"testing"
)

func TestYearZodiac(t *testing.T) {
	tests := []struct {
		year int
		want string
	}{
		{2024, "龙"}, {2023, "兔"}, {1990, "马"}, {2000, "龙"},
	}
	for _, tt := range tests {
		t.Run(string(rune(tt.year)), func(t *testing.T) {
			cal := ParseSolarDate(tt.year, 6, 15, 12, 0, 0)
			got := cal.Lunar().GetZodiac()
			if got.Chinese() != tt.want {
				t.Errorf("YearZodiac(%d) = %s, want %s", tt.year, got.Chinese(), tt.want)
			}
		})
	}
}

func TestYearZodiacNoFix(t *testing.T) {
	tests := []struct {
		year int
		want string
	}{
		{2024, "龙"}, {2023, "兔"}, {1990, "马"},
	}
	for _, tt := range tests {
		got := YearZodiacNoFix(tt.year)
		if got.Chinese() != tt.want {
			t.Errorf("YearZodiacNoFix(%d) = %s, want %s", tt.year, got.Chinese(), tt.want)
		}
	}
}

func TestZodiacChinese(t *testing.T) {
	got, err := ZodiacChinese(ZodiacRat)
	if err != nil {
		t.Fatal(err)
	}
	if got != "鼠" {
		t.Errorf("ZodiacChinese(Rat) = %s, want 鼠", got)
	}
}

func TestParseZodiac(t *testing.T) {
	got, err := ParseZodiac("Rat")
	if err != nil {
		t.Fatal(err)
	}
	if got != ZodiacRat {
		t.Errorf("ParseZodiac(Rat) = %d, want %d", got, ZodiacRat)
	}
}

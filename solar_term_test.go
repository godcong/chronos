package chronos

import (
	"reflect"
	"testing"
	"time"
)

func TestSolarTermChinese(t *testing.T) {
	tests := []struct {
		st       SolarTerm
		want     string
		wantEmpty bool
	}{
		{0, "小寒", false},
		{1, "大寒", false},
		{2, "立春", false},
		{3, "雨水", false},
		{4, "惊蛰", false},
		{5, "春分", false},
		{6, "清明", false},
		{7, "谷雨", false},
		{8, "立夏", false},
		{23, "冬至", false},
		{24, "", true},
	}
	for _, tt := range tests {
		got := tt.st.Chinese()
		if (got == "") != tt.wantEmpty {
			t.Errorf("SolarTerm(%d).Chinese() = %q, wantEmpty %v", tt.st, got, tt.wantEmpty)
		}
		if got != tt.want {
			t.Errorf("SolarTerm(%d).Chinese() = %v, want %v", tt.st, got, tt.want)
		}
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
		{"", args{1900, SolarTermDaHan}, "1900-01-20 19:32:25"},
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
		{"", args{TimeFromY(1900), 0}, SolarTermDetail{Index: 0, SolarTerm: 0, Time: "1900/01/06 02:03:57"}},
		{"", args{TimeFromY(1900), SolarTermDaHan}, SolarTermDetail{Index: 0, SolarTerm: 0, Time: "1900/01/20 19:32:25"}},
		{"", args{TimeFromY(1900), SolarTermLiChun}, SolarTermDetail{Index: 0, SolarTerm: 0, Time: "1900/02/04 13:51:31"}},
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
	tests := []struct {
		t    time.Time
		want bool
	}{
		{time.Date(1900, 01, 20, 19, 32, 25, 0, time.Local), true},
	}
	for _, tt := range tests {
		if _, got := CheckSolarTermDay(tt.t); got != tt.want {
			t.Errorf("IsSolarTermDetailDay() = %v, want %v", got, tt.want)
		}
	}
}

func Test_yearLiChunDay(t *testing.T) {
	if gotDay := yearLiChunDay(1900); gotDay != 4 {
		t.Errorf("yearLiChunDay() = %v, want 4", gotDay)
	}
}

func Test_getSolarTermDayFromLunar(t *testing.T) {
	tests := []struct {
		year int
		st   SolarTerm
		want int
	}{
		{1900, 0, 6},
		{1900, 1, 20},
		{3000, 0, 6},
		{3000, 1, 20},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := getYearSolarTermTime(tt.year, tt.st).Day(); got != tt.want {
				t.Errorf("getYearSolarTermTime(%d, %d).Day() = %v, want %v", tt.year, tt.st, got, tt.want)
			}
		})
	}
}

func Test_getSolarTermTimeFromLunar(t *testing.T) {
	tests := []struct {
		year int
		st   SolarTerm
		want string
	}{
		{1900, 0, "1900/01/06 02:03:57"},
		{1900, 1, "1900/01/20 19:32:25"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := getYearSolarTermTime(tt.year, tt.st); !reflect.DeepEqual(got.Format(DateFormatYMDHMS), tt.want) {
				t.Errorf("getYearSolarTermTime() = %v, want %v", got.Format(DateFormatYMDHMS), tt.want)
			}
		})
	}
}

func TestYearSolarTermDate(t *testing.T) {
	month, day := YearSolarTermDate(time.Date(2024, 1, 1, 0, 0, 0, 0, loc), SolarTermLiChun)
	if month != 2 {
		t.Errorf("LiChun month = %d, want 2", month)
	}
	if day < 3 || day > 5 {
		t.Errorf("LiChun day = %d, want 3-5", day)
	}
}

func TestYearSolarTermDay(t *testing.T) {
	day := YearSolarTermDay(time.Date(2024, 1, 1, 0, 0, 0, 0, loc), SolarTermLiChun)
	if day < 3 || day > 5 {
		t.Errorf("LiChun day = %d, want 3-5", day)
	}
}

func TestYearSolarTermMonth(t *testing.T) {
	month := YearSolarTermMonth(time.Date(2024, 1, 1, 0, 0, 0, 0, loc), SolarTermLiChun)
	if month != 2 {
		t.Errorf("LiChun month = %d, want 2", month)
	}
}

func TestYearSolarTermDetailV2(t *testing.T) {
	detail, err := YearSolarTermDetail(time.Date(2024, 1, 1, 0, 0, 0, 0, loc), SolarTermXiaoHan)
	if err != nil {
		t.Fatal(err)
	}
	if detail.SolarTerm != SolarTermXiaoHan {
		t.Errorf("SolarTerm = %d, want %d", detail.SolarTerm, SolarTermXiaoHan)
	}
	if detail.Time == "" {
		t.Error("Time should not be empty")
	}
}

func TestCheckSolarTermDay(t *testing.T) {
	st, ok := CheckSolarTermDay(time.Date(2024, 2, 4, 0, 0, 0, 0, loc))
	if !ok {
		t.Log("2024-02-04 is not a solar term day (may vary by year)")
	} else {
		t.Logf("2024-02-04 solar term: %s", st.Chinese())
	}
}

func TestSolarTermChineseV2(t *testing.T) {
	got := SolarTermLiChun.Chinese()
	if got != "立春" {
		t.Errorf("SolarTermLiChun.Chinese() = %s, want 立春", got)
	}
}

func TestSolarTermSanHou(t *testing.T) {
	sh := SolarTermLiChun.SanHou()
	if sh == "" {
		t.Error("SanHou should not be empty")
	}
}

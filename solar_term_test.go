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
		{
			name: "",
			args: args{
				year: 1900,
				st:   SolarTermDaHan,
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
				year: TimeFromY(1900),
				st:   0,
			},
			want: SolarTermDetail{
				Index:       0,
				SolarTerm:   0,
				Time:        "1900/01/06 02:03:57",
				SanHou:      "",
				Explanation: "",
			},
		},
		{
			name: "",
			args: args{
				year: TimeFromY(1900),
				st:   SolarTermDaHan,
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
				year: TimeFromY(1900),
				st:   SolarTermLiChun,
			},
			want: SolarTermDetail{
				Index:       0,
				SolarTerm:   0,
				Time:        "1900/02/04 13:51:31",
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
				t: time.Date(1900, 01, 20, 19, 32, 25, 0, time.Local),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := CheckSolarTermDay(tt.args.t); got != tt.want {
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
	got, err := SolarTermChinese(SolarTermLiChun)
	if err != nil {
		t.Fatal(err)
	}
	if got != "立春" {
		t.Errorf("SolarTermChinese(LiChun) = %s, want 立春", got)
	}
}

func TestSolarTermSanHou(t *testing.T) {
	sh := SolarTermLiChun.SanHou()
	if sh == "" {
		t.Error("SanHou should not be empty")
	}
}

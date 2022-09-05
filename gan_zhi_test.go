package chronos

import (
	"fmt"
	"testing"
	"time"

	"github.com/godcong/chronos/v2/utils"
)

func TestNianZhuChineseV2(t *testing.T) {
	type args struct {
		y time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				y: yearDate(1900),
			},
			want: "庚子",
		},
		{
			name: "",
			args: args{
				y: yearDate(1899),
			},
			want: "己亥",
		},
		{
			name: "",
			args: args{
				y: yearDate(2099),
			},
			want: "己未",
		},
		{
			name: "",
			args: args{
				y: yearDate(2100),
			},
			want: "庚申",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NianZhuChineseV2(tt.args.y); got != tt.want {
				t.Errorf("NianZhuChineseV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGanZhiChinese(t *testing.T) {
	type args struct {
		ganzhi GanZhi
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
				ganzhi: 0,
			},
			want:    "甲子",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				ganzhi: 59,
			},
			want:    "癸亥",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				ganzhi: 60,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GanZhiChinese(tt.args.ganzhi)
			if (err != nil) != tt.wantErr {
				t.Errorf("GanZhiChinese() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GanZhiChinese() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYueZhu(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				t: yearMonthDayDate(1900, 1, 1),
			},
			want: "丙子",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YueZhuChineseV2(tt.args.t); got != tt.want {
				t.Errorf("YueZhuChineseV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monthGanZhi(t *testing.T) {
	type args struct {
		year  int
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
				year:  1900,
				month: 1,
				day:   1,
			},
			want: "丙子",
		},
		{
			name: "",
			args: args{
				year:  1900,
				month: 11,
				day:   3,
			},
			want: "丙戌",
		},
		{
			name: "",
			args: args{
				year:  2099,
				month: 11,
				day:   3,
			},
			want: "甲戌",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := yueZhu(tt.args.year, tt.args.month, tt.args.day); got.Chinese() != tt.want {
				t.Errorf("yueZhu() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitGanZhi(t *testing.T) {
	type args struct {
		gz GanZhi
	}
	tests := []struct {
		name  string
		args  args
		want  TianGan
		want1 DiZhi
	}{
		{
			name: "",
			args: args{
				gz: GanZhiJiaWu,
			},
			want:  TianGanJia,
			want1: DiZhiWu,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitGanZhi(tt.args.gz)
			if got != tt.want {
				t.Errorf("splitGanZhi() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("splitGanZhi() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseGanZhi(t *testing.T) {
	type args struct {
		tiangan TianGan
		dizhi   DiZhi
	}
	tests := []struct {
		name    string
		args    args
		want    GanZhi
		wantErr bool
	}{

		{
			name: "",
			args: args{
				tiangan: TianGanJia,
				dizhi:   0,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				tiangan: TianGanJia,
				dizhi:   DiZhiWu,
			},
			want:    GanZhiJiaWu,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				tiangan: TianGanJia,
				dizhi:   DiZhiChou,
			},
			want: GanZhiMax,
		},
		{
			name: "",
			args: args{
				tiangan: TianGanYi,
				dizhi:   DiZhiHai,
			},
			want: GanZhiYiHai,
		},
		{
			name: "",
			args: args{
				tiangan: TianGanGeng,
				dizhi:   DiZhiHai,
			},
			want: GanZhiMax,
		},
		{
			name: "",
			args: args{
				tiangan: TianGanGeng,
				dizhi:   DiZhiHai,
			},
			want: GanZhiMax,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseGanZhiV2(tt.args.tiangan, tt.args.dizhi)
			if got != tt.want {
				t.Errorf("parseGanZhiV2() got = %v, want %v", got, tt.want)
			}
			got = parseGanZhi(tt.args.tiangan, tt.args.dizhi)
			if got != tt.want {
				t.Errorf("parseGanZhiV2() got = %v, want %v", got, tt.want)
			}
		})
	}
	tm, err := time.Parse(DefaultDateFormat, "2000/01/01 00:00:00")
	fmt.Printf("time:%X,%v\n", uint64(tm.UTC().Unix()), err)

	fmt.Println("diff day:", utils.DateDiffDay(tm, startTime))
}

func TestShiZhu(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				t: time.Date(2022, 9, 5, 14, 42, 30, 0, time.UTC),
			},
			want: "乙未",
		},
		{
			name: "",
			args: args{
				t: time.Date(2022, 9, 5, 23, 42, 30, 0, time.UTC),
			},
			want: "庚子",
		},
		{
			name: "",
			args: args{
				t: time.Date(2022, 9, 6, 0, 42, 30, 0, time.UTC),
			},
			want: "庚子",
		},
		{
			name: "",
			args: args{
				t: time.Date(2022, 10, 1, 0, 42, 30, 0, time.UTC),
			},
			want: "庚子",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShiZhu(tt.args.t); got.Chinese() != tt.want {
				t.Errorf("ShiZhu() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRiZhu(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				t: yearMonthDayDate(2022, 9, 5),
			},
			want: "辛酉",
		},
		{
			name: "",
			args: args{
				t: yearMonthDayDate(2099, 11, 4),
			},
			want: "乙巳",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RiZhu(tt.args.t); got.Chinese() != tt.want {
				t.Errorf("RiZhu() = %v, want %v", got, tt.want)
			}
		})
	}
}

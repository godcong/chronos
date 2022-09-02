package chronos

import (
	"testing"
	"time"
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
			wantErr: true,
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
			wantErr: true,
		},
		{
			name: "",
			args: args{
				tiangan: TianGanGeng,
				dizhi:   DiZhiHai,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseGanZhi(tt.args.tiangan, tt.args.dizhi)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseGanZhi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseGanZhi() got = %v, want %v", got, tt.want)
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
			if got := YueZhu(tt.args.t); got != tt.want {
				t.Errorf("YueZhu() = %v, want %v", got, tt.want)
			}
		})
	}
}

package chronos

import (
	"fmt"
	"testing"
	"time"
)

func TestConstellationChinese(t *testing.T) {
	type args struct {
		c Constellation
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
				c: 0,
			},
			want:    "魔羯",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConstellationChinese(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConstellationChinese() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConstellationChinese() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetConstellation(t *testing.T) {
	type args struct {
		month time.Month
		day   int
	}
	tests := []struct {
		name    string
		args    args
		want    Constellation
		wantErr bool
	}{
		{
			name: "",
			args: args{
				month: 0,
				day:   0,
			},
			want:    -1,
			wantErr: true,
		},
		{
			name: "",
			args: args{
				month: 1,
				day:   1,
			},
			want:    0,
			wantErr: false,
		},
	}
	tm, err := time.Parse(DefaultDateFormat, "2006/01/01 01:01")
	if err != nil {
		return
	}
	//m := time.Month(0)

	fmt.Println("month", tm.Month(), "day", tm.Day())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConstellation(tt.args.month, tt.args.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConstellation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetConstellation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

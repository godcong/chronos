package chronos

import (
	"testing"
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
		// TODO: Add test cases.
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
			want:    "谷 雨",
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

package chronos

import (
	"testing"
)

func TestNianZhuChineseV2(t *testing.T) {
	type args struct {
		y int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				y: 1900,
			},
			want: "庚子",
		},
		{
			name: "",
			args: args{
				y: 1899,
			},
			want: "己亥",
		},
		{
			name: "",
			args: args{
				y: 2099,
			},
			want: "己未",
		},
		{
			name: "",
			args: args{
				y: 2100,
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

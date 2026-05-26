package chronos

import (
	"testing"
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

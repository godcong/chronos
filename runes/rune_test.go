package runes

import (
	"testing"
)

func TestRunes_FindString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		r    Runes
		args args
		want int
	}{
		{
			name: "",
			r:    []rune("你好世界Hello Worldこんにちわ!"),
			args: args{
				s: "世界",
			},
			want: 2,
		},
		{
			name: "",
			r:    []rune("你好世界Hello Worldこんにちわ!"),
			args: args{
				s: "世界",
			},
			want: 4,
		},
		{
			name: "",
			r:    []rune("你好世界Hello Worldこんにちわ!"),
			args: args{
				s: "こん",
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.FindString(tt.args.s); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

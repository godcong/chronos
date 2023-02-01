package runes

import (
	"testing"
)

func TestLastIndex(t *testing.T) {
	type args struct {
		s   []rune
		sep []rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s:   []rune("hello world"),
				sep: []rune("hello world"),
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				s:   []rune("hello world"),
				sep: []rune("world"),
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				s:   []rune("hello world"),
				sep: []rune("world!"),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastIndex(tt.args.s, tt.args.sep); got != tt.want {
				t.Errorf("LastIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	type args struct {
		s   []rune
		sep []rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s:   []rune("hello world"),
				sep: []rune("hello world"),
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				s:   []rune("hello world"),
				sep: []rune("world"),
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				s:   []rune("hello world"),
				sep: []rune("world!"),
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				s:   []rune("你好世界Hello Worldこんにちわ!"),
				sep: []rune("好"),
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				s:   []rune("你好世界Hello Worldこんにちわ!"),
				sep: []rune("こん"),
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Index(tt.args.s, tt.args.sep); got != tt.want {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

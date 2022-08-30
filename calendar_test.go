package chronos

import (
	"reflect"
	"testing"
	"time"
)

func TestNewSolarCalendar(t *testing.T) {
	type args struct {
		v []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				v: nil,
			},
			want: time.Now().Format(DefaultDateFormat),
		},
		{
			name: "",
			args: args{
				v: []any{
					"2022/08/24 12:34",
				},
			},
			want: "2022/08/24 12:34",
		},
		{
			name: "",
			args: args{
				v: []any{
					"2022-08-24 12:34",
					"2006-01-02 15:04",
				},
			},
			want: "2022/08/24 12:34",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSolarCalendar(tt.args.v...); !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

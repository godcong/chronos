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
					"2022/08/24 12:34:00",
				},
			},
			want: "2022/08/24 12:34:00",
		},
		{
			name: "",
			args: args{
				v: []any{
					"2022-08-24 12:34",
					"2006-01-02 15:04",
				},
			},
			want: "2022/08/24 12:34:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSolarCalendar(tt.args.v...); !reflect.DeepEqual(got.FormatTime(), tt.want) {
				t.Errorf("New() = %v, want %v", got.FormatTime(), tt.want)
			}
		})
	}
}

func Test_isToday(t *testing.T) {
	type args struct {
		t   time.Time
		now time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				t:   yearDate(1900),
				now: time.Now(),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isToday(tt.args.t, tt.args.now); got != tt.want {
				t.Errorf("isToday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSolarNow(t *testing.T) {
	tests := []struct {
		name string
		want Calendar
	}{
		{
			name: "",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseSolarNow()
			if got == nil {
				t.Errorf("ParseSolarNow() = %v, want %v", got, tt.want)
			}
			date := got.Date()
			t.Log("date:", date)
		})
	}
}

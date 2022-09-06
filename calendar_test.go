package chronos

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/godcong/chronos/v2/utils"
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
			want: time.Now().Format(DateFormatYMDHMS),
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
			got := NewSolarCalendar(tt.args.v...)
			if !reflect.DeepEqual(got.FormatTime(), tt.want) {
				t.Errorf("New() = %v, want %v", got.FormatTime(), tt.want)
			}
		})
	}
}

func TestMonthDayCheck(t *testing.T) {
	//t.Log(got.Date())
	for idx := 1900; idx < 2100; idx++ {
		t := yearDate(idx)
		days := utils.CalcYearMonthDays(t.Year())
		var mdays []int
		//if idx == 2034 {
		//	lm, err1 := LeapMonth(t.AddDate(-1, 0, 0))
		//	s, err2 := LeapMonthBS(t.AddDate(-1, 0, 0))
		//	if 11 == lm && err1 == nil && err2 == nil {
		//		if s == LeapMonthBig {
		//			mdays = append(mdays, 30)
		//		} else {
		//			mdays = append(mdays, 29)
		//		}
		//	}
		//}
		lm, _ := LeapMonth(t)
		for m := 1; m <= 12; m++ {
			mdays = append(mdays, monthDays(t.Year(), m, lm, false))
			//lm, err1 := LeapMonth(t)
			//s, err2 := LeapMonthBS(t)
			if m == lm {
				mdays = append(mdays, monthDays(t.Year(), m, lm, true))
			}
			//	if s == LeapMonthBig {
			//		mdays = append(mdays, 30)
			//	} else {
			//		mdays = append(mdays, 29)
			//	}
			//}

			//fmt.Printf("wrong year: %d,month: %d,day: (%d,%d)\n", t.Year(), t.Month(), days[m-1], mday)
		}
		//idx, err1 := LeapMonth(t)
		//s, err2 := LeapMonthBS(t)
		//if err1 == nil && err2 == nil {
		//	tmp := make([]int, len(mdays)+1)
		//	//copy(tmp, mdays[0:idx])
		//	fmt.Println("tmp", tmp)
		//	if s == LeapMonthBig {
		//		//mdays = append(mdays[0:idx], []int{120,mdays[idx:]...}, mdays[idx:])
		//	} else {
		//		//tmp[idx] = 129
		//	}
		//
		//	//mdays = append(tmp, mdays[idx:]...)
		//}
		for i, mday := range mdays {
			if days[i] != mday {
				fmt.Printf("wrong year: %d,month: %d,day: (%d,%d)\n", t.Year(), i+1, days[i], mday)
			}
		}
		fmt.Printf("month days: %+v,%+v\n", days, mdays)
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

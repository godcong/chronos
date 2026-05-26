package chronos

import (
	"time"
)

func parseFirstArg(v []any) any {
	if len(v) > 0 {
		return v[0]
	}
	return nil
}

func parseArgs(v []any) []any {
	if len(v) > 1 {
		return v[1:]
	}
	return nil
}

func parseStringFormat(f string, v ...any) string {
	if len(v) == 0 {
		return f
	}
	df, ok := (v[0]).(string)
	if ok {
		f = df
	}
	return f
}

func parseStringTime(v any) string {
	t := ""
	df, ok := (v).(string)
	if ok {
		t = df
	}
	return t
}

func parseStringDateFormat(t string, vv ...string) *calendarTime {
	c := &calendarTime{
		time: localTime(),
	}
	if t == "" {
		return c
	}
	f := DateFormatYMDHMS
	if len(vv) > 0 {
		f = vv[0]
	}
	tt, err := time.ParseInLocation(f, t, loc)
	if err == nil {
		c.time = tt
	}
	return c
}

func parseStringDate(t string, vv ...any) *calendarTime {
	c := &calendarTime{
		time: localTime(),
	}
	if t == "" {
		return c
	}

	f := parseStringFormat(DateFormatYMDHMS, vv...)
	tt, err := time.ParseInLocation(f, t, loc)
	if err == nil {
		c.time = tt
	}
	return c
}

func parseTime(t time.Time) *calendarTime {
	c := &calendarTime{
		time: t,
	}
	return c
}

// ParseTime parse time.Time to Calendar
// @param time.Time
// @param *time.Location
// @return Calendar
func ParseTime(t time.Time) Calendar {
	return parseTime(t).initialize()
}

func parseIntDate(vv int, args ...any) *calendarTime {
	var c calendarTime
	switch len(args) {
	case 0:
		c.time = TimeFromY(vv)
	case 1:
		c.time = TimeFromYm(vv, args[0].(time.Month))
	case 2:
		c.time = TimeFromYmd(vv, args[0].(time.Month), args[1].(int))
	case 3:
	case 4:
		c.time = TimeFromYmd(vv, args[0].(time.Month), args[1].(int))
	//case 5:
	//	c.time = TimeFromYmdHms(vv, args[0].(time.Month), args[1].(int), args[2].(int), args[3].(int), args[4].(int))
	default:
		c.time = TimeFromYmdHms(vv, args[0].(time.Month), args[1].(int), args[2].(int), args[3].(int), args[4].(int))
	}
	return &c
}

func localTime() time.Time {
	return time.Now().In(loc)
}

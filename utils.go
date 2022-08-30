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

func parseStringDate(t string, vv ...any) *calendar {
	c := &calendar{
		loc:  time.Local,
		time: time.Now(),
	}
	if t == "" {
		return c
	}

	f := parseStringFormat(DefaultDateFormat, vv...)
	tt, err := time.Parse(f, t)
	if err == nil {
		c.time = tt
	}
	return c
}

// ParseTime parse time.Time to Calendar
// @param time.Time
// @param *time.Location
// @return Calendar
func ParseTime(t time.Time, local *time.Location) Calendar {
	return &calendar{
		loc:  local,
		time: t,
	}
}

package chronos

import (
	"time"

	"github.com/6tail/lunar-go/calendar"
)

// solar ...
type solar struct {
	*calendar.Solar
}

func ParseSolarByTime(date time.Time) Solar {
	return &solar{Solar: calendar.NewSolarFromDate(date)}
}

var _ Solar = &solar{}

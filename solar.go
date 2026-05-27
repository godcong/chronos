package chronos

import (
	"time"

	"github.com/6tail/lunar-go/calendar"
)

type solar struct {
	*calendar.Solar
}

func (s solar) GetConstellation() Constellation {
	return constellationFromChinese(s.GetXingZuo())
}

func (s solar) GetFestivals() []string {
	return listToStrings(s.Solar.GetFestivals())
}

func (s solar) GetOtherFestivals() []string {
	return listToStrings(s.Solar.GetOtherFestivals())
}

// ParseSolarByTime creates a Solar from a time.Time value.
func ParseSolarByTime(date time.Time) Solar {
	return &solar{Solar: calendar.NewSolarFromDate(date)}
}

var _ Solar = &solar{}

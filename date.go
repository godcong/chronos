package chronos

import (
	"time"
)

type SolarDate struct {
	Year    int          `json:"year"`
	Month   time.Month   `json:"month"`
	Day     int          `json:"day"`
	Hour    int          `json:"hour"`
	Minute  int          `json:"minute"`
	Second  int          `json:"second"`
	WeekDay time.Weekday `json:"week_day"`
}

type LunarDate struct {
	Year        int  `json:"year"`
	Month       int  `json:"month"`
	Day         int  `json:"day"`
	Hour        int  `json:"hour"`
	IsLeapMonth bool `json:"is_leap_month"`
	LeapMonth   int  `json:"leap_month"`
}

type EightCharacter struct {
	NianZhu GanZhi `json:"nian_zhu"`
	YueZhu  GanZhi `json:"yue_zhu"`
	Rizhu   GanZhi `json:"ri_zhu"`
	ShiZhu  GanZhi `json:"shi_zhu"`
}

type CalendarDate struct {
	IsToday        bool           `json:"is_today"`
	Solar          SolarDate      `json:"solar"`
	Lunar          LunarDate      `json:"lunar"`
	EightCharacter EightCharacter `json:"eight_character"`
	Zodiac         Zodiac         `json:"zodiac"`
	Constellation  Constellation  `json:"constellation"`
	IsSolarTermDay bool           `json:"is_solar_day"`
	SolarTerm      SolarTerm      `json:"solar_term"`
}

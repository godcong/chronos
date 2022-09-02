package chronos

import (
	"time"
)

type CalendarDate struct {
	IsToday bool `json:"is_today"`
	Solar   struct {
		Year    int          `json:"year"`
		Month   time.Month   `json:"month"`
		Day     int          `json:"day"`
		Hour    int          `json:"hour"`
		WeekDay time.Weekday `json:"week_day"`
	} `json:"solar"`
	Lunar struct {
		Year        int  `json:"year"`
		Month       int  `json:"month"`
		Day         int  `json:"day"`
		Hour        int  `json:"hour"`
		IsLeapMonth bool `json:"is_leap_month"`
		IsLeapYear  bool `json:"is_leap_year"`
	} `json:"lunar"`
	Ganzhi struct {
		NianZhu string `json:"nian_zhu"`
		YueZhu  string `json:"yue_zhu"`
		Rizhu   string `json:"ri_zhu"`
		ShiZhu  string `json:"shi_zhu"`
	}
	Zodiac         Zodiac        `json:"zodiac"`
	Constellation  Constellation `json:"constellation"`
	IsSolarTermDay bool          `json:"is_solar_day"`
	SolarTerm      SolarTerm     `json:"solar_term"`
}

package chronos

type View struct {
	IsToday string `json:"is_today"`
	WeekDay string `json:"week_day"`
	Solar   struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
		Hour  int `json:"hour"`
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
		Year  string `json:"year"`
		Month string `json:"month"`
		Day   string `json:"day"`
		Hour  string `json:"hour"`
	}
	Zodiac        string `json:"zodiac"`
	SolarTerm     string `json:"solar_term"`
	Constellation string `json:"constellation"`
}

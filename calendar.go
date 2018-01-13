package lunar

import "time"

type Calendar struct {
	lunar *Lunar
	solar *Solar
}

func CalendarFromLunar(y,m,d int) Calendar{
	return Calendar{
		lunar:&Lunar{
			year:y,
			month:m,
			day:d,
		},
	}
}

func CalendarFromSolar(time time.Time)Calendar  {
	return Calendar{
		solar:&Solar{
			time:time,
		},
	}
}

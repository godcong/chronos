package chronos

// solar ...
type solar struct {
	year  int
	month int
	day   int
	hour  int
}

func (s *solar) Year() int {
	return s.year
}

func (s *solar) Month() int {
	return s.month
}

func (s *solar) Day() int {
	return s.day
}

func (s *solar) Hour() int {
	return s.hour
}

var _ Solar = &solar{}

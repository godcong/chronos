package chronos

import "time"

// Solar ...
type Solar struct {
	time time.Time
}

// NewSolar ...
func NewSolar(calendar Calendar) *Solar {
	return calendar.Solar()
}

func (s *Solar) Time() time.Time {
	return s.time
}

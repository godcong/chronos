package chronos

import (
	"time"
)

// Calendar ...
type Calendar interface {
	Lunar() *Lunar
	Solar() *Solar
	LocalTime() time.Time
	ViewData() View
}

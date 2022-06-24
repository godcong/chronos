package chronos

import (
	"time"
)

// Calendar ...
type Calendar interface {
	Lunar() *Lunar
	Solar() *Solar
	Time() time.Time
}

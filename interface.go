package chronos

import (
	"time"
)

// Calendar ...
type Calendar interface {
	Lunar() *Lunar
	Solar() *Solar
	Time() time.Time
	//Bind(v any) error
	//JSON() []byte
	String() string
}

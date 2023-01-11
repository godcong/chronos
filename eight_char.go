package chronos

import (
	"github.com/6tail/lunar-go/calendar"
)

type eightChar struct {
	*calendar.EightChar
}

var _ EightChar = (*eightChar)(nil)

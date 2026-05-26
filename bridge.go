package chronos

import "time"

// Bridge provides compatibility with the old fate/chronos bridge interface
type Bridge struct {
	calendar Calendar
}

// NewBridge creates a new Bridge from a time
func NewBridge(t time.Time) *Bridge {
	return &Bridge{
		calendar: ParseSolarTime(t),
	}
}

// Calendar returns the underlying Calendar
func (b *Bridge) Calendar() Calendar {
	return b.calendar
}

// Lunar returns the Lunar interface
func (b *Bridge) Lunar() Lunar {
	return b.calendar.Lunar()
}

// Solar returns the Solar interface
func (b *Bridge) Solar() Solar {
	return b.calendar.Solar()
}

// EightChar returns the EightChar interface
func (b *Bridge) EightChar() EightChar {
	return b.calendar.Lunar().GetEightChar()
}

// SiZhu returns the four pillars
func (b *Bridge) SiZhu() [4]string {
	return b.EightChar().GetSiZhu()
}

// WuXing returns the five elements of the four pillars
func (b *Bridge) WuXing() [4]string {
	return b.EightChar().GetWuXing()
}

// Zodiac returns the zodiac animal
func (b *Bridge) Zodiac() string {
	return b.calendar.Lunar().GetZodiac().String()
}

// Constellation returns the constellation
func (b *Bridge) Constellation() string {
	return b.calendar.Solar().GetConstellation().String()
}

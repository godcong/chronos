package chronos

import (
	"testing"
	"time"
)

func TestBridge_NewBridge(t *testing.T) {
	birthDate := time.Date(2023, 2, 5, 12, 0, 0, 0, time.Local)
	b := NewBridge(birthDate)

	if b == nil {
		t.Fatal("expected non-nil Bridge")
	}
}

func TestBridge_Calendar(t *testing.T) {
	birthDate := time.Date(2023, 2, 5, 12, 0, 0, 0, time.Local)
	b := NewBridge(birthDate)

	c := b.Calendar()
	if c == nil {
		t.Error("expected non-nil Calendar")
	}
}

func TestBridge_Lunar(t *testing.T) {
	birthDate := time.Date(2023, 2, 5, 12, 0, 0, 0, time.Local)
	b := NewBridge(birthDate)

	l := b.Lunar()
	if l == nil {
		t.Error("expected non-nil Lunar")
	}
}

func TestBridge_Solar(t *testing.T) {
	birthDate := time.Date(2023, 2, 5, 12, 0, 0, 0, time.Local)
	b := NewBridge(birthDate)

	s := b.Solar()
	if s == nil {
		t.Error("expected non-nil Solar")
	}
}

func TestBridge_EightChar(t *testing.T) {
	birthDate := time.Date(2023, 2, 5, 12, 0, 0, 0, time.Local)
	b := NewBridge(birthDate)

	ec := b.EightChar()
	if ec == nil {
		t.Error("expected non-nil EightChar")
	}
}

func TestBridge_SiZhu(t *testing.T) {
	birthDate := time.Date(2023, 2, 5, 12, 0, 0, 0, time.Local)
	b := NewBridge(birthDate)

	sz := b.SiZhu()
	if len(sz) != 4 {
		t.Errorf("expected 4 pillars, got %d", len(sz))
	}
}

func TestBridge_WuXing(t *testing.T) {
	birthDate := time.Date(2023, 2, 5, 12, 0, 0, 0, time.Local)
	b := NewBridge(birthDate)

	wx := b.WuXing()
	if len(wx) != 4 {
		t.Errorf("expected 4 elements, got %d", len(wx))
	}
}

func TestBridge_Zodiac(t *testing.T) {
	birthDate := time.Date(2023, 2, 5, 12, 0, 0, 0, time.Local)
	b := NewBridge(birthDate)

	z := b.Zodiac()
	if z == "" {
		t.Error("expected non-empty zodiac")
	}
}

func TestBridge_Constellation(t *testing.T) {
	birthDate := time.Date(2023, 2, 5, 12, 0, 0, 0, time.Local)
	b := NewBridge(birthDate)

	c := b.Constellation()
	if c == "" {
		t.Error("expected non-empty constellation")
	}
}

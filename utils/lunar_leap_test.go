package utils

import (
	"testing"
)

func TestInitLeap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitLeap()
		})
	}
}

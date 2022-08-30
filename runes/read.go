package runes

import (
	"errors"
	"fmt"
)

type Runes []rune

var ErrTooLarge = errors.New("runes.Runes: too large")

func (r Runes) ReadString(offset int, limit int) (string, error) {
	fmt.Println(len(r), offset+limit, offset, limit)
	if len(r) < offset+limit {
		return "", ErrTooLarge
	}
	return string(r[offset : offset+limit]), nil
}

func ReadString(runs []rune, offset int, limit int) (string, error) {
	return Runes(runs).ReadString(offset, limit)
}

package runes

import (
	"errors"
)

//const runeSize = len(rune)/len(byte)

type Runes []rune

var ErrTooLarge = errors.New("runes.Runes: number is too large than length")

func (r Runes) ReadString(offset int, limit int) (string, error) {
	//fmt.Println(len(r), offset+limit, offset, limit)
	if len(r) < offset+limit {
		return "", ErrTooLarge
	}
	return string(r[offset : offset+limit]), nil
}

func (r Runes) MustReadString(offset int, limit int) string {
	readString, err := r.ReadString(offset, limit)
	if err != nil {
		panic(err)
	}
	return readString
}

func (r Runes) Index(sub []rune) int {
	return Index(r, sub)
}

func (r Runes) FindString(s string) int {
	return r.Index([]rune(s))
}

func ReadString(runs []rune, offset int, limit int) (string, error) {
	return Runes(runs).ReadString(offset, limit)
}

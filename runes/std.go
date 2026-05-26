package runes

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// explode splits s into a slice of UTF-8 []runes,
// one []rune per Unicode character up to a maximum of n (n < 0 means no limit).
// Invalid UTF-8 sequences become correct encodings of U+FFFD.
func explode(s []rune, n int) [][]rune {
	l := len(s)
	if n < 0 || n > l {
		n = l
	}
	a := make([][]rune, n)
	for i := 0; i < n-1; i++ {
		a[i] = s[:1]
		s = s[1:]
	}
	if n > 0 {
		a[n-1] = s
	}
	return a
}

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty []rune, Count returns 1 + the number of Unicode code points in s.
func Count(s, substr []rune) int {
	return strings.Count(string(s), string(substr))
}

// Contains reports whether substr is within s.
func Contains(s, substr []rune) bool {
	return strings.Index(string(s), string(substr)) >= 0
}

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny(s []rune, chars string) bool {
	return IndexAny(s, chars) >= 0
}

// ContainsRune reports whether the Unicode code point r is within s.
func ContainsRune(s []rune, r rune) bool {
	return IndexRune(s, r) >= 0
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex(s, sep []rune) int {
	n := len(sep)
	switch {
	case n == 0:
		return len(s)
	case n == 1:
		return LastIndexRune(s, sep[0])
	case n == len(s):
		if Equal(s, sep) {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	}
	last := len(s) - n
	if Equal(sep, s[last:]) {
		return last
	}
	for i := last - 1; i >= 0; i-- {
		if Equal(sep, s[i:i+n]) {
			return i
		}
	}
	return -1
}

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func IndexByte(s []rune, c byte) int {
	return IndexRune(s, rune(c))
}

// IndexRune returns the index of the first instance of the runes point
// r, or -1 if rune is not present in s.
func IndexRune(s []rune, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return -1
}

// IndexAny returns the index of the first instance of any Unicode code point
// from chars in s, or -1 if no Unicode code point from chars is present in s.
func IndexAny(s []rune, chars string) int {
	//if chars == "" {
	//	// Avoid scanning all of s.
	//	return -1
	//}
	//if len(s) == 1 {
	//	r := s[0]
	//	if r >= utf8.RuneSelf {
	//		// search utf8.RuneError.
	//		for _, r = range chars {
	//			if r == utf8.RuneError {
	//				return 0
	//			}
	//		}
	//		return -1
	//	}
	//	if bytealg.IndexByteString(chars, s[0]) >= 0 {
	//		return 0
	//	}
	//	return -1
	//}
	//if len(chars) == 1 {
	//	r := rune(chars[0])
	//	if r >= utf8.RuneSelf {
	//		r = utf8.RuneError
	//	}
	//	return IndexRune(s, r)
	//}
	//if len(s) > 8 {
	//	if as, isASCII := makeASCIISet(chars); isASCII {
	//		for i, c := range s {
	//			if as.contains(c) {
	//				return i
	//			}
	//		}
	//		return -1
	//	}
	//}
	//var width int
	//for i := 0; i < len(s); i += width {
	//	r := rune(s[i])
	//	if r < utf8.RuneSelf {
	//		if bytealg.IndexByteString(chars, s[i]) >= 0 {
	//			return i
	//		}
	//		width = 1
	//		continue
	//	}
	//	r, width = utf8.DecodeRune(s[i:])
	//	if r != utf8.RuneError {
	//		// r is 2 to 4 bytes
	//		if len(chars) == width {
	//			if chars == string(r) {
	//				return i
	//			}
	//			continue
	//		}
	//		// Use bytealg.IndexString for performance if available.
	//		if bytealg.MaxLen >= width {
	//			if bytealg.IndexString(chars, string(r)) >= 0 {
	//				return i
	//			}
	//			continue
	//		}
	//	}
	//	for _, ch := range chars {
	//		if r == ch {
	//			return i
	//		}
	//	}
	//}
	return -1

}

// LastIndexAny returns the index of the last instance of any Unicode code
// point from chars in s, or -1 if no Unicode code point from chars is
// present in s.
func LastIndexAny(s, chars []rune) int {
	//if chars == "" {
	//	// Avoid scanning all of s.
	//	return -1
	//}
	//if len(s) > 8 {
	//	if as, isASCII := makeASCIISet(chars); isASCII {
	//		for i := len(s) - 1; i >= 0; i-- {
	//			if as.contains(s[i]) {
	//				return i
	//			}
	//		}
	//		return -1
	//	}
	//}
	//if len(s) == 1 {
	//	r := rune(s[0])
	//	if r >= utf8.RuneSelf {
	//		for _, r = range chars {
	//			if r == utf8.RuneError {
	//				return 0
	//			}
	//		}
	//		return -1
	//	}
	//	if bytealg.IndexByteString(chars, s[0]) >= 0 {
	//		return 0
	//	}
	//	return -1
	//}
	//if len(chars) == 1 {
	//	cr := chars[0]
	//	for i := len(s); i > 0; {
	//		r, size := s[i-1], 1
	//		i -= size
	//		if r == cr {
	//			return i
	//		}
	//	}
	//	return -1
	//}
	//for i := len(s); i > 0; {
	//	r := s[i-1]
	//
	//	r, size := s[i-1], 1
	//	i -= size
	//	if r != utf8.RuneError {
	//		// r is 2 to 4 bytes
	//		if len(chars) == size {
	//			if chars == string(r) {
	//				return i
	//			}
	//			continue
	//		}
	//		// Use bytealg.IndexString for performance if available.
	//		if bytealg.MaxLen >= size {
	//			if bytealg.IndexString(chars, string(r)) >= 0 {
	//				return i
	//			}
	//			continue
	//		}
	//	}
	//	for _, ch := range chars {
	//		if r == ch {
	//			return i
	//		}
	//	}
	//}
	return -1
}

// LastIndexRune returns the index of the last instance of c in s, or -1 if c is not present in s.
func LastIndexRune(s []rune, c rune) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}

// Generic split: splits after each instance of sep,
// including sepSave bytes of sep in the subarrays.
func genSplit(s, sep []rune, sepSave, n int) [][]rune {
	if n == 0 {
		return nil
	}
	if len(sep) == 0 {
		return explode(s, n)
	}
	if n < 0 {
		n = Count(s, sep) + 1
	}
	if n > len(s)+1 {
		n = len(s) + 1
	}

	a := make([][]rune, n)
	n--
	i := 0
	for i < n {
		m := Index(s, sep)
		if m < 0 {
			break
		}
		a[i] = s[: m+sepSave : m+sepSave]
		s = s[m+len(sep):]
		i++
	}
	a[i] = s
	return a[:i+1]
}

// SplitN slices s into sub[]runes separated by sep and returns a slice of
// the sub[]runes between those separators.
//
// The count determines the number of sub[]runes to return:
//
//	n > 0: at most n sub[]runes; the last sub[]rune will be the unsplit remainder.
//	n == 0: the result is nil (zero sub[]runes)
//	n < 0: all sub[]runes
//
// Edge cases for s and sep (for example, empty []runes) are handled
// as described in the documentation for Split.
//
// To split around the first instance of a separator, see Cut.
func SplitN(s, sep []rune, n int) [][]rune { return genSplit(s, sep, 0, n) }

// SplitAfterN slices s into sub[]runes after each instance of sep and
// returns a slice of those sub[]runes.
//
// The count determines the number of sub[]runes to return:
//
//	n > 0: at most n sub[]runes; the last sub[]rune will be the unsplit remainder.
//	n == 0: the result is nil (zero sub[]runes)
//	n < 0: all sub[]runes
//
// Edge cases for s and sep (for example, empty []runes) are handled
// as described in the documentation for SplitAfter.
func SplitAfterN(s, sep []rune, n int) [][]rune {
	return genSplit(s, sep, len(sep), n)
}

// Split slices s into all sub[]runes separated by sep and returns a slice of
// the sub[]runes between those separators.
//
// If s does not contain sep and sep is not empty, Split returns a
// slice of length 1 whose only element is s.
//
// If sep is empty, Split splits after each UTF-8 sequence. If both s
// and sep are empty, Split returns an empty slice.
//
// It is equivalent to SplitN with a count of -1.
//
// To split around the first instance of a separator, see Cut.
func Split(s, sep []rune) [][]rune { return genSplit(s, sep, 0, -1) }

// SplitAfter slices s into all sub[]runes after each instance of sep and
// returns a slice of those sub[]runes.
//
// If s does not contain sep and sep is not empty, SplitAfter returns
// a slice of length 1 whose only element is s.
//
// If sep is empty, SplitAfter splits after each UTF-8 sequence. If
// both s and sep are empty, SplitAfter returns an empty slice.
//
// It is equivalent to SplitAfterN with a count of -1.
func SplitAfter(s, sep []rune) [][]rune {
	return genSplit(s, sep, len(sep), -1)
}

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// Fields splits the []rune s around each instance of one or more consecutive white space
// characters, as defined by unicode.IsSpace, returning a slice of sub[]runes of s or an
// empty slice if s contains only white space.
func Fields(s []rune) [][]rune {
	// First count the fields.
	// This is an exact count if s is ASCII, otherwise it is an approximation.
	n := 0
	wasSpace := 1
	// setBits is used to track which bits are set in the bytes of s.
	setBits := int32(0)
	for i := 0; i < len(s); i++ {
		r := s[i]
		setBits |= r
		isSpace := int(asciiSpace[r])
		n += wasSpace & ^isSpace
		wasSpace = isSpace
	}

	if setBits >= utf8.RuneSelf {
		// Some runes in the input []rune are not ASCII.
		return FieldsFunc(s, unicode.IsSpace)
	}
	// ASCII fast path
	a := make([][]rune, n)
	na := 0
	fieldStart := 0
	i := 0
	// Skip spaces in the front of the input.
	for i < len(s) && asciiSpace[s[i]] != 0 {
		i++
	}
	fieldStart = i
	for i < len(s) {
		if asciiSpace[s[i]] == 0 {
			i++
			continue
		}
		a[na] = s[fieldStart:i]
		na++
		i++
		// Skip spaces in between fields.
		for i < len(s) && asciiSpace[s[i]] != 0 {
			i++
		}
		fieldStart = i
	}
	if fieldStart < len(s) { // Last field might end at EOF.
		a[na] = s[fieldStart:]
	}
	return a
}

// FieldsFunc splits the []rune s at each run of Unicode code points c satisfying f(c)
// and returns an array of slices of s. If all code points in s satisfy f(c) or the
// []rune is empty, an empty slice is returned.
//
// FieldsFunc makes no guarantees about the order in which it calls f(c)
// and assumes that f always returns the same value for a given c.
func FieldsFunc(s []rune, f func(rune) bool) [][]rune {
	// A span is used to record a slice of s of the form s[start:end].
	// The start index is inclusive and the end index is exclusive.
	type span struct {
		start int
		end   int
	}
	spans := make([]span, 0, 32)

	// Find the field start and end indices.
	// Doing this in a separate pass (rather than slicing the []rune s
	// and collecting the result sub[]runes right away) is significantly
	// more efficient, possibly due to cache effects.
	start := -1 // valid span start if >= 0
	for end, rune := range s {
		if f(rune) {
			if start >= 0 {
				spans = append(spans, span{start, end})
				// Set start to a negative value.
				// Note: using -1 here consistently and reproducibly
				// slows down this code by a several percent on amd64.
				start = ^start
			}
		} else {
			if start < 0 {
				start = end
			}
		}
	}

	// Last field might end at EOF.
	if start >= 0 {
		spans = append(spans, span{start, len(s)})
	}

	// Create []runes from recorded field indices.
	a := make([][]rune, len(spans))
	for i, span := range spans {
		a[i] = s[span.start:span.end]
	}

	return a
}

// Join concatenates the elements of its first argument to create a single []rune. The separator
// []rune sep is placed between elements in the resulting []rune.
func Join(s [][]rune, sep []rune) []rune {
	if len(s) == 0 {
		return []rune{}
	}
	if len(s) == 1 {
		// Just return a copy.
		return append([]rune(nil), s[0]...)
	}
	n := len(sep) * (len(s) - 1)
	for _, v := range s {
		n += len(v)
	}

	b := make([]rune, n)
	bp := copy(b, s[0])
	for _, v := range s[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], v)
	}
	return b
}

// HasPrefix tests whether the []rune s begins with prefix.
func HasPrefix(s, prefix []rune) bool {
	return len(s) >= len(prefix) && Equal(s[0:len(prefix)], prefix)
}

// HasSuffix tests whether the []rune s ends with suffix.
func HasSuffix(s, suffix []rune) bool {
	return len(s) >= len(suffix) && Equal(s[len(s)-len(suffix):], suffix)
}

// Map returns a copy of the []rune s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the []rune with no replacement.
func Map(mapping func(rune) rune, s []rune) []rune {
	// In the worst case, the slice can grow when mapped, making
	// things unpleasant. But it's so rare we barge in assuming it's
	// fine. It could also shrink but that falls out naturally.
	maxrunes := len(s) // length of b
	nrunes := 0        // number of bytes encoded in b
	b := make([]rune, maxrunes)
	for i := 0; i < len(s); {
		wid := 1
		r := s[i]
		r = mapping(r)
		if r >= 0 {
			rl := utf8.RuneLen(r)
			if rl < 0 {
				rl = len(string(utf8.RuneError))
			}
			if nrunes+rl > maxrunes {
				// Grow the buffer.
				maxrunes = maxrunes*2 + utf8.UTFMax
				nb := make([]rune, maxrunes)
				copy(nb, b[0:nrunes])
				b = nb
			}
			nrunes++
		}
		i += wid
	}
	return b[0:nrunes]
}

// Repeat returns a new []rune consisting of count copies of the []rune s.
//
// It panics if count is negative or if
// the result of (len(s) * count) overflows.
func Repeat(b []rune, count int) []rune {
	if count == 0 {
		return []rune{}
	}
	// Since we cannot return an error on overflow,
	// we should panic if the repeat will generate
	// an overflow.
	// See Issue golang.org/issue/16237.
	if count < 0 {
		panic("bytes: negative Repeat count")
	} else if len(b)*count/count != len(b) {
		panic("bytes: Repeat count causes overflow")
	}

	nb := make([]rune, len(b)*count)
	bp := copy(nb, b)
	for bp < len(nb) {
		copy(nb[bp:], nb[:bp])
		bp *= 2
	}
	return nb
}

// ToUpper returns s with all Unicode letters mapped to their upper case.
func ToUpper(s []rune) []rune {
	hasLower := true
	for i := 0; i < len(s); i++ {
		c := s[i]
		//if c >= utf8.RuneSelf {
		//	isASCII = false
		//	break
		//}
		hasLower = hasLower || ('a' <= c && c <= 'z')
	}

	if !hasLower {
		// Just return a copy.
		return append([]rune(""), s...)
	}
	b := make([]rune, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'a' <= c && c <= 'z' {
			c -= 'a' - 'A'
		}
		b[i] = c
	}
	return b
}

// ToLower returns s with all Unicode letters mapped to their lower case.
func ToLower(s []rune) []rune {
	hasUpper := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		hasUpper = hasUpper || ('A' <= c && c <= 'Z')
	}

	if !hasUpper {
		return append([]rune(""), s...)
	}
	b := make([]rune, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
		}
		b[i] = c
	}
	return b

}

// ToTitle returns a copy of the []rune s with all Unicode letters mapped to
// their Unicode title case.
func ToTitle(s []rune) []rune { return Map(unicode.ToTitle, s) }

// ToUpperSpecial returns a copy of the []rune s with all Unicode letters mapped to their
// upper case using the case mapping specified by c.
func ToUpperSpecial(c unicode.SpecialCase, s []rune) []rune {
	return Map(c.ToUpper, s)
}

// ToLowerSpecial returns a copy of the []rune s with all Unicode letters mapped to their
// lower case using the case mapping specified by c.
func ToLowerSpecial(c unicode.SpecialCase, s []rune) []rune {
	return Map(c.ToLower, s)
}

// ToTitleSpecial returns a copy of the []rune s with all Unicode letters mapped to their
// Unicode title case, giving priority to the special casing rules.
func ToTitleSpecial(c unicode.SpecialCase, s []rune) []rune {
	return Map(c.ToTitle, s)
}

// ToValidUTF8 returns a copy of the []rune s with each run of invalid UTF-8 byte sequences
// replaced by the replacement []rune, which may be empty.
func ToValidUTF8(s, replacement []rune) []rune {
	b := make([]rune, 0, len(s)+len(replacement))
	invalid := false // previous byte was from an invalid UTF-8 sequence
	for i := 0; i < len(s); {
		c := s[i]
		if c < utf8.RuneSelf {
			i++
			invalid = false
			b = append(b, c)
			continue
		}
		i++
		if !invalid {
			invalid = true
			b = append(b, replacement...)
		}
		continue
	}
	return b
}

// isSeparator reports whether the rune could mark a word boundary.
// TODO: update when package unicode captures more of the properties.
func isSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}

// Title returns a copy of the []rune s with all Unicode letters that begin words
// mapped to their Unicode title case.
//
// Deprecated: The rule Title uses for word boundaries does not handle Unicode
// punctuation properly. Use golang.org/x/text/cases instead.
func Title(s []rune) []rune {
	// Use a closure here to remember state.
	// Hackish but effective. Depends on Map scanning in order and calling
	// the closure once per rune.
	prev := ' '
	return Map(
		func(r rune) rune {
			if isSeparator(prev) {
				prev = r
				return unicode.ToTitle(r)
			}
			prev = r
			return r
		},
		s)
}

// TrimLeftFunc returns a slice of the []rune s with all leading
// Unicode code points c satisfying f(c) removed.
func TrimLeftFunc(s []rune, f func(rune) bool) []rune {
	i := indexFunc(s, f, false)
	if i == -1 {
		return nil
	}
	return s[i:]
}

// TrimRightFunc returns a slice of the []rune s with all trailing
// Unicode code points c satisfying f(c) removed.
func TrimRightFunc(s []rune, f func(rune) bool) []rune {
	i := lastIndexFunc(s, f, false)
	if i >= 0 && s[i] >= utf8.RuneSelf {
		i++
	} else {
		i++
	}
	return s[0:i]
}

// TrimFunc returns a slice of the []rune s with all leading
// and trailing Unicode code points c satisfying f(c) removed.
func TrimFunc(s []rune, f func(rune) bool) []rune {
	return TrimRightFunc(TrimLeftFunc(s, f), f)
}

// IndexFunc returns the index into s of the first Unicode
// code point satisfying f(c), or -1 if none do.
func IndexFunc(s []rune, f func(rune) bool) int {
	return indexFunc(s, f, true)
}

// LastIndexFunc returns the index into s of the last
// Unicode code point satisfying f(c), or -1 if none do.
func LastIndexFunc(s []rune, f func(rune) bool) int {
	return lastIndexFunc(s, f, true)
}

// indexFunc is the same as IndexFunc except that if
// truth==false, the sense of the predicate function is
// inverted.
func indexFunc(s []rune, f func(rune) bool, truth bool) int {
	for i, r := range s {
		if f(r) == truth {
			return i
		}
	}
	return -1
}

// lastIndexFunc is the same as LastIndexFunc except that if
// truth==false, the sense of the predicate function is
// inverted.
func lastIndexFunc(s []rune, f func(rune) bool, truth bool) int {
	for i := len(s); i > 0; {
		r := s[i-1]
		i--
		if f(r) == truth {
			return i
		}
	}
	return -1
}

// asciiSet is a 32-byte value, where each bit represents the presence of a
// given ASCII character in the set. The 128-bits of the lower 16 bytes,
// starting with the least-significant bit of the lowest word to the
// most-significant bit of the highest word, map to the full range of all
// 128 ASCII characters. The 128-bits of the upper 16 bytes will be zeroed,
// ensuring that any non-ASCII character will be reported as not in the set.
// This allocates a total of 32 bytes even though the upper half
// is unused to avoid bounds checks in asciiSet.contains.
type asciiSet [8]uint32

// makeASCIISet creates a set of ASCII characters and reports whether all
// characters in chars are ASCII.
func makeASCIISet(chars string) (as asciiSet, ok bool) {
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		if c >= utf8.RuneSelf {
			return as, false
		}
		as[c/32] |= 1 << (c % 32)
	}
	return as, true
}

// contains reports whether c is inside the set.
func (as *asciiSet) contains(c rune) bool {
	return (as[c/32] & (1 << (c % 32))) != 0
}

// containsRune is a simplified version of strings.ContainsRune
// to avoid importing the strings package.
// We avoid bytes.ContainsRune to avoid allocating a temporary copy of s.
func containsRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

// Trim returns a subslice of s by slicing off all leading and
// trailing UTF-8-encoded code points contained in cutset.
func Trim(s []rune, cutset string) []rune {
	if len(s) == 0 {
		// This is what we've historically done.
		return nil
	}
	if cutset == "" {
		return s
	}
	//if len(cutset) == 1 && cutset[0] < utf8.RuneSelf {
	//	return trimLeftRune(trimRightRune(s, cutset[0]), cutset[0])
	//}
	//if as, ok := makeASCIISet(cutset); ok {
	//	return trimLeftASCII(trimRightASCII(s, &as), &as)
	//}
	//return trimLeftUnicode(trimRightUnicode(s, cutset), cutset)
	panic("unreachable")
}

// TrimLeft returns a subslice of s by slicing off all leading
// UTF-8-encoded code points contained in cutset.
func TrimLeft(s []rune, cutset string) []rune {
	if len(s) == 0 {
		// This is what we've historically done.
		return nil
	}
	if cutset == "" {
		return s
	}
	//if len(cutset) == 1 && cutset[0] < utf8.RuneSelf {
	//	return trimLeftByte(s, cutset[0])
	//}
	//if as, ok := makeASCIISet(cutset); ok {
	//	return trimLeftASCII(s, &as)
	//}
	//return trimLeftUnicode(s, cutset)
	panic("unreachable")
}

func trimLefRune(s []rune, c rune) []rune {
	for len(s) > 0 && s[0] == c {
		s = s[1:]
	}
	if len(s) == 0 {
		// This is what we've historically done.
		return nil
	}
	return s
}

func trimLeftUnicode(s []rune, cutset string) []rune {
	for len(s) > 0 {
		r, n := s[0], 1
		if !containsRune(cutset, r) {
			break
		}
		s = s[n:]
	}
	if len(s) == 0 {
		// This is what we've historically done.
		return nil
	}
	return s
}

// TrimRight returns a slice of the []rune s, with all trailing
// Unicode code points contained in cutset removed.
//
// To remove a suffix, use TrimSuffix instead.
func TrimRight(s []rune, cutset string) []rune {
	if len(s) == 0 || cutset == "" {
		return s
	}
	//if len(cutset) == 1 && cutset[0] < utf8.RuneSelf {
	//	return trimRightRune(s, []rune(cutset)[0])
	//}
	//if as, ok := makeASCIISet(cutset); ok {
	//	return trimRightASCII(s, &as)
	//}
	//return trimRightUnicode(s, cutset)
	panic("unreachable")
}

func trimRightRune(s []rune, c rune) []rune {
	for len(s) > 0 && s[len(s)-1] == c {
		s = s[:len(s)-1]
	}
	return s
}

func trimRightASCII(s []rune, as *asciiSet) []rune {
	for len(s) > 0 {
		if !as.contains(s[len(s)-1]) {
			break
		}
		s = s[:len(s)-1]
	}
	return s
}

func trimRightUnicode(s, cutset []rune) []rune {
	for len(s) > 0 {
		r, n := s[len(s)-1], 1
		if !ContainsRune(cutset, r) {
			break
		}
		s = s[:len(s)-n]
	}
	return s
}

// TrimSpace returns a slice of the []rune s, with all leading
// and trailing white space removed, as defined by Unicode.
func TrimSpace(s []rune) []rune {
	// Fast path for ASCII: look for the first ASCII non-space byte
	start := 0
	for ; start < len(s); start++ {
		c := s[start]
		if c >= utf8.RuneSelf {
			// If we run into a non-ASCII byte, fall back to the
			// slower unicode-aware method on the remaining bytes
			return TrimFunc(s[start:], unicode.IsSpace)
		}
		if asciiSpace[c] == 0 {
			break
		}
	}

	// Now look for the first ASCII non-space byte from the end
	stop := len(s)
	for ; stop > start; stop-- {
		c := s[stop-1]
		if c >= utf8.RuneSelf {
			// start has been already trimmed above, should trim end only
			return TrimRightFunc(s[start:stop], unicode.IsSpace)
		}
		if asciiSpace[c] == 0 {
			break
		}
	}

	// At this point s[start:stop] starts and ends with an ASCII
	// non-space bytes, so we're done. Non-ASCII cases have already
	// been handled above.
	return s[start:stop]
}

// TrimPrefix returns s without the provided leading prefix []rune.
// If s doesn't start with prefix, s is returned unchanged.
func TrimPrefix(s, prefix []rune) []rune {
	if HasPrefix(s, prefix) {
		return s[len(prefix):]
	}
	return s
}

// TrimSuffix returns s without the provided trailing suffix []rune.
// If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix(s, suffix []rune) []rune {
	if HasSuffix(s, suffix) {
		return s[:len(s)-len(suffix)]
	}
	return s
}

// Replace returns a copy of the []rune s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the []rune
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune []rune.
// If n < 0, there is no limit on the number of replacements.
func Replace(s, old, new []rune, n int) []rune {
	m := 0
	if n != 0 {
		// Compute number of replacements.
		m = Count(s, old)
	}
	if m == 0 {
		// Just return a copy.
		return append([]rune(nil), s...)
	}
	if n < 0 || m < n {
		n = m
	}

	// Apply replacements to buffer.
	t := make([]rune, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				wid := len(s[start:])
				j += wid
			}
		} else {
			j += Index(s[start:], old)
		}
		w += copy(t[w:], s[start:j])
		w += copy(t[w:], new)
		start = j + len(old)
	}
	w += copy(t[w:], s[start:])
	return t[0:w]
}

// ReplaceAll returns a copy of the []rune s with all
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the []rune
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune []rune.
func ReplaceAll(s, old, new []rune) []rune {
	return Replace(s, old, new, -1)
}

// EqualFold reports whether s and t, interpreted as UTF-8 []runes,
// are equal under simple Unicode case-folding, which is a more general
// form of case-insensitivity.
func EqualFold(s, t []rune) bool {
	for len(s) != 0 && len(t) != 0 {
		// Extract first rune from each []rune.
		var sr, tr rune
		sr, s = s[0], s[1:]
		tr, t = t[0], t[1:]
		// If they match, keep going; if not, return false.

		// Easy case.
		if tr == sr {
			continue
		}

		// Make sr < tr to simplify what follows.
		if tr < sr {
			tr, sr = sr, tr
		}
		// Fast check for ASCII.
		if tr < utf8.RuneSelf {
			// ASCII only, sr/tr must be upper/lower case
			if 'A' <= sr && sr <= 'Z' && tr == sr+'a'-'A' {
				continue
			}
			return false
		}

		// General case. SimpleFold(x) returns the next equivalent rune > x
		// or wraps around to smaller values.
		r := unicode.SimpleFold(sr)
		for r != sr && r < tr {
			r = unicode.SimpleFold(r)
		}
		if r == tr {
			continue
		}
		return false
	}

	// One []rune is empty. Are both?
	return len(s) == len(t)
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s, sep []rune) int {
	n := len(sep)
	switch {
	case n == 0:
		return 0
	case n == 1:
		return IndexRune(s, sep[0])
	case n == len(s):
		if Equal(sep, s) {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	}
	c0 := sep[0]
	c1 := sep[1]
	i := 0
	t := len(s) - n + 1
	for i < t {
		if s[i] != c0 {
			o := IndexRune(s[i+1:t], c0)
			if o < 0 {
				return -1
			}
			i += o + 1
		}
		if s[i+1] == c1 && string(s[i:i+n]) == string(sep) {
			return i
		}
		i++
	}
	return -1
}

// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
func Cut(s, sep []rune) (before, after []rune, found bool) {
	if i := Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, nil, false
}

// Equal reports whether a and b
// are the same length and contain the same runes.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []rune) bool {
	return string(a) == string(b)
}

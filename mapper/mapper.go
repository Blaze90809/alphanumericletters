package mapper

import (
	"fmt"
	"strings"
	"unicode"
)

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

type SkipString struct {
	Skip  int
	Value string
	SkipIndex int // Use this to skip non-alphanumeric characters.
}

func MapString(i Interface) {
   for pos := range i.GetValueAsRuneSlice() {
      i.TransformRune(pos)
   }
}

func (s SkipString) GetValueAsRuneSlice() []rune {
	// If the skip is 0, return before runtime error: integer divide by zero.
	if s.Skip == 0 {
		return []rune{}
	}

	return []rune(s.Value)
}

func (s *SkipString) TransformRune(pos int) {
	if pos == 0 {
		s.ResetSkipIndex()
	}
	// Need to account for non-alphanumeric characters.
	const alpha = "abcdefghijklmnopqrstuvwxyz0123456789"

	if !strings.Contains(alpha, strings.ToLower(string([]rune(s.Value)[pos]))) {
		s.IncrementSkipIndex()
		return
	}

	if (pos + 1 - s.SkipIndex) % s.Skip == 0 {
		i := s.GetValueAsRuneSlice()
		i[pos] = unicode.ToUpper(rune(i[pos]))
		s.Value = string(i)
	} else {
		i := s.GetValueAsRuneSlice()
		i[pos] = unicode.ToLower(rune(i[pos]))
		s.Value = string(i)
	}
}

func NewSkipString(n int, s string) SkipString {
	return SkipString{Skip: n, Value: s, SkipIndex: 0}
}

func (d SkipString) String() string {
	return fmt.Sprintf("Every %v letter(s) should be capitalized: %v", d.Skip, d.Value)
}

func (d *SkipString) IncrementSkipIndex() {
	d.SkipIndex = d.SkipIndex + 1
}

func (d *SkipString) ResetSkipIndex() {
	d.SkipIndex = 0
}
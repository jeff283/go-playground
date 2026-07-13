// Package textutil provides helpers for exploring UTF-8 strings and runes.
//
// # Overview
//
// This package exists as a learning ground for how Go's doc-comment
// conventions map onto rendered documentation in godoc and pkgsite. It
// covers indexing quirks with multi-byte characters, rune-safe string
// manipulation, and a couple of small utility types.
//
// # Usage
//
//	rev := textutil.Reverse("Résumé")
//	fmt.Println(rev)
//
// See the Examples on this package's pkg.go.dev page for runnable
// demonstrations of each function.
package textutil

import (
	"errors"
	"strings"
)

// DefaultPadChar is the rune used by Pad when no explicit padding
// character is supplied.
const DefaultPadChar = ' '

// Case represents a target letter-casing style for NormalizeCase.
type Case int

// Supported Case values.
const (
	// Lower normalizes text to lowercase.
	Lower Case = iota
	// Upper normalizes text to uppercase.
	Upper
	// Title normalizes text to Title Case.
	Title
)

// ErrEmptyInput is returned by functions in this package when given an
// empty string where a non-empty value was required.
var ErrEmptyInput = errors.New("textutil: empty input")

// Reverse returns the input string with its runes in reverse order.
// It correctly handles multi-byte UTF-8 characters, unlike naive
// byte-reversal.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Pad pads s on both sides with DefaultPadChar until it reaches width
// runes long. If s is already at or beyond width, it is returned
// unchanged.
func Pad(s string, width int) string {
	runeLen := len([]rune(s))
	if runeLen >= width {
		return s
	}
	total := width - runeLen
	left := total / 2
	right := total - left
	return strings.Repeat(string(DefaultPadChar), left) + s + strings.Repeat(string(DefaultPadChar), right)
}

// CountRunes returns the number of runes (not bytes) in s.
//
// Deprecated: use utf8.RuneCountInString from the standard library
// instead; this function is kept only for teaching purposes.
func CountRunes(s string) int {
	return len([]rune(s))
}

// Counter accumulates rune counts across multiple strings. Its zero
// value is ready to use.
type Counter struct {
	total int
}

// Add adds the rune count of s to the running total and returns the
// new total, so calls can be chained.
func (c *Counter) Add(s string) int {
	c.total += len([]rune(s))
	return c.total
}

// Total returns the current accumulated rune count.
func (c *Counter) Total() int {
	return c.total
}

// Stringer is implemented by types that can render themselves as a
// human-readable string. It mirrors fmt.Stringer and exists here so
// pkgsite has an interface type to render.
type Stringer interface {
	// String returns a human-readable representation of the value.
	String() string
}

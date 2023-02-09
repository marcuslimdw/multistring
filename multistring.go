// Package multistring provides wrappers around the standard library's strings package's functions that take string
// slices.
package multistring

import (
	"strings"
)

type Wrap []string

// JoinBy joins the strings in the `Wrap` using a separator.
func (w Wrap) JoinBy(sep string) string {
	return strings.Join(w, sep)
}

// MapString applies a function over the strings in the Wrap, returning a new Wrap.
func (w Wrap) MapString(mapping func(string) string) Wrap {
	result := make(Wrap, len(w))
	for i, s := range w {
		result[i] = mapping(s)
	}

	return result
}

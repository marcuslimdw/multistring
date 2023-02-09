package multistring

import (
	"strings"
)

type Wrap []string

func (w Wrap) JoinBy(sep string) string {
	return strings.Join(w, sep)
}

func (w Wrap) MapString(mapping func(string) string) Wrap {
	result := make(Wrap, len(w))
	for i, s := range w {
		result[i] = mapping(s)
	}

	return result
}

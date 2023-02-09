package multistring

import (
	"strings"
	"unicode"
)

func (w Wrap) Count(substr string) []int {
	result0 := make([]int, len(w))
	for i, s := range w {
		result0[i] = strings.Count(s, substr)
	}
	return result0
}

func (w Wrap) Contains(substr string) []bool {
	result0 := make([]bool, len(w))
	for i, s := range w {
		result0[i] = strings.Contains(s, substr)
	}
	return result0
}

func (w Wrap) ContainsAny(chars string) []bool {
	result0 := make([]bool, len(w))
	for i, s := range w {
		result0[i] = strings.ContainsAny(s, chars)
	}
	return result0
}

func (w Wrap) ContainsRune(r rune) []bool {
	result0 := make([]bool, len(w))
	for i, s := range w {
		result0[i] = strings.ContainsRune(s, r)
	}
	return result0
}

func (w Wrap) LastIndex(substr string) []int {
	result0 := make([]int, len(w))
	for i, s := range w {
		result0[i] = strings.LastIndex(s, substr)
	}
	return result0
}

func (w Wrap) IndexByte(c byte) []int {
	result0 := make([]int, len(w))
	for i, s := range w {
		result0[i] = strings.IndexByte(s, c)
	}
	return result0
}

func (w Wrap) IndexRune(r rune) []int {
	result0 := make([]int, len(w))
	for i, s := range w {
		result0[i] = strings.IndexRune(s, r)
	}
	return result0
}

func (w Wrap) IndexAny(chars string) []int {
	result0 := make([]int, len(w))
	for i, s := range w {
		result0[i] = strings.IndexAny(s, chars)
	}
	return result0
}

func (w Wrap) LastIndexAny(chars string) []int {
	result0 := make([]int, len(w))
	for i, s := range w {
		result0[i] = strings.LastIndexAny(s, chars)
	}
	return result0
}

func (w Wrap) LastIndexByte(c byte) []int {
	result0 := make([]int, len(w))
	for i, s := range w {
		result0[i] = strings.LastIndexByte(s, c)
	}
	return result0
}

func (w Wrap) SplitN(sep string, n int) []Wrap {
	result0 := make([]Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.SplitN(s, sep, n)
	}
	return result0
}

func (w Wrap) SplitAfterN(sep string, n int) []Wrap {
	result0 := make([]Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.SplitAfterN(s, sep, n)
	}
	return result0
}

func (w Wrap) Split(sep string) []Wrap {
	result0 := make([]Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.Split(s, sep)
	}
	return result0
}

func (w Wrap) SplitAfter(sep string) []Wrap {
	result0 := make([]Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.SplitAfter(s, sep)
	}
	return result0
}

func (w Wrap) Fields() []Wrap {
	result0 := make([]Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.Fields(s)
	}
	return result0
}

func (w Wrap) FieldsFunc(f func(rune) bool) []Wrap {
	result0 := make([]Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.FieldsFunc(s, f)
	}
	return result0
}

func (w Wrap) Join(elems []string) Wrap {
	result0 := make(Wrap, len(w))
	for i, sep := range w {
		result0[i] = strings.Join(elems, sep)
	}
	return result0
}

func (w Wrap) HasPrefix(prefix string) []bool {
	result0 := make([]bool, len(w))
	for i, s := range w {
		result0[i] = strings.HasPrefix(s, prefix)
	}
	return result0
}

func (w Wrap) HasSuffix(suffix string) []bool {
	result0 := make([]bool, len(w))
	for i, s := range w {
		result0[i] = strings.HasSuffix(s, suffix)
	}
	return result0
}

func (w Wrap) Map(mapping func(rune) rune) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.Map(mapping, s)
	}
	return result0
}

func (w Wrap) Repeat(count int) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.Repeat(s, count)
	}
	return result0
}

func (w Wrap) ToUpper() Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.ToUpper(s)
	}
	return result0
}

func (w Wrap) ToLower() Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.ToLower(s)
	}
	return result0
}

func (w Wrap) ToTitle() Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.ToTitle(s)
	}
	return result0
}

func (w Wrap) ToUpperSpecial(c unicode.SpecialCase) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.ToUpperSpecial(c, s)
	}
	return result0
}

func (w Wrap) ToLowerSpecial(c unicode.SpecialCase) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.ToLowerSpecial(c, s)
	}
	return result0
}

func (w Wrap) ToTitleSpecial(c unicode.SpecialCase) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.ToTitleSpecial(c, s)
	}
	return result0
}

func (w Wrap) ToValidUTF8(replacement string) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.ToValidUTF8(s, replacement)
	}
	return result0
}

func (w Wrap) Title() Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.Title(s)
	}
	return result0
}

func (w Wrap) TrimLeftFunc(f func(rune) bool) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.TrimLeftFunc(s, f)
	}
	return result0
}

func (w Wrap) TrimRightFunc(f func(rune) bool) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.TrimRightFunc(s, f)
	}
	return result0
}

func (w Wrap) TrimFunc(f func(rune) bool) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.TrimFunc(s, f)
	}
	return result0
}

func (w Wrap) IndexFunc(f func(rune) bool) []int {
	result0 := make([]int, len(w))
	for i, s := range w {
		result0[i] = strings.IndexFunc(s, f)
	}
	return result0
}

func (w Wrap) LastIndexFunc(f func(rune) bool) []int {
	result0 := make([]int, len(w))
	for i, s := range w {
		result0[i] = strings.LastIndexFunc(s, f)
	}
	return result0
}

func (w Wrap) Trim(cutset string) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.Trim(s, cutset)
	}
	return result0
}

func (w Wrap) TrimLeft(cutset string) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.TrimLeft(s, cutset)
	}
	return result0
}

func (w Wrap) TrimRight(cutset string) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.TrimRight(s, cutset)
	}
	return result0
}

func (w Wrap) TrimSpace() Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.TrimSpace(s)
	}
	return result0
}

func (w Wrap) TrimPrefix(prefix string) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.TrimPrefix(s, prefix)
	}
	return result0
}

func (w Wrap) TrimSuffix(suffix string) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.TrimSuffix(s, suffix)
	}
	return result0
}

func (w Wrap) Replace(old string, new string, n int) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.Replace(s, old, new, n)
	}
	return result0
}

func (w Wrap) ReplaceAll(old string, new string) Wrap {
	result0 := make(Wrap, len(w))
	for i, s := range w {
		result0[i] = strings.ReplaceAll(s, old, new)
	}
	return result0
}

func (w Wrap) EqualFold(t string) []bool {
	result0 := make([]bool, len(w))
	for i, s := range w {
		result0[i] = strings.EqualFold(s, t)
	}
	return result0
}

func (w Wrap) Index(substr string) []int {
	result0 := make([]int, len(w))
	for i, s := range w {
		result0[i] = strings.Index(s, substr)
	}
	return result0
}

func (w Wrap) Cut(sep string) (before, after Wrap, found []bool) {
	result0 := make(Wrap, len(w))
	result1 := make(Wrap, len(w))
	result2 := make([]bool, len(w))
	for i, s := range w {
		result0[i], result1[i], result2[i] = strings.Cut(s, sep)
	}
	return result0, result1, result2
}

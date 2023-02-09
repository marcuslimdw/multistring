package internal_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"multistring/internal"
)

var _ = Describe("Functional", func() {

	Context("Map", func() {

		It("should return a nil slice given a nil slice", func() {
			actual := internal.Map[int, string](nil, nil)

			Expect(actual).To(BeNil())
		})

		It("should apply a function to a slice and collect the results into a new slice", func() {
			actual := internal.Map([]string{"a", "bcd", "ef"}, func(s string) int { return len(s) })
			expected := []int{1, 3, 2}

			Expect(actual).To(Equal(expected))
		})
	})

	Context("FlatMap", func() {

		It("should return a nil slice given a nil slice", func() {
			actual := internal.FlatMap[int, string](nil, nil)

			Expect(actual).To(BeNil())
		})

		It("should apply a function to a slice and collect the results into a new slice", func() {
			actual := internal.FlatMap([]string{"a b", "c def g", "h ijkl mn"}, strings.Fields)
			expected := []string{"a", "b", "c", "def", "g", "h", "ijkl", "mn"}

			Expect(actual).To(Equal(expected))
		})
	})
})

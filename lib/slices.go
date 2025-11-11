package lib

import (
	"iter"
	"slices"
)

// MapSlice is a convenience function for mapping a slice of Ts to a slice of Us.
func MapSlice[T, U any](s []T, fn func(T) U) []U {
	return slices.Collect(Map(slices.Values(s), fn))
}

// Pairs returns an iterator over every possible pair of elements in the slice.
// Every combination of two elements is yielded exactly once; the order of the elements
// in the pair will be the same as the order they appear in the slice.
func Pairs[T any](in []T) iter.Seq2[T, T] {
	return func(yield func(T, T) bool) {
		for i := range in {
			for j := i + 1; j < len(in); j++ {
				if !yield(in[i], in[j]) {
					return
				}
			}
		}
	}
}

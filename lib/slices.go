package lib

import (
	"iter"
	"slices"
)

func Count[T any](in []T, f func(T) bool) int {
	return CountSeq(slices.Values(in), f)
}

func CountSeq[T any](in iter.Seq[T], f func(T) bool) int {
	count := 0
	for v := range in {
		if f(v) {
			count++
		}
	}
	return count
}

func Every[T any](in []T, f func(T) bool) bool {
	return EverySeq(slices.Values(in), f)
}

func EverySeq[T any](in iter.Seq[T], f func(T) bool) bool {
	for v := range in {
		if !f(v) {
			return false
		}
	}
	return true
}

func Map[T, U any](in []T, f func(T) U) []U {
	return slices.Collect(MapSeq(slices.Values(in), f))
}

func MapSeq[T, U any](in iter.Seq[T], f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for v := range in {
			if !yield(f(v)) {
				return
			}
		}
	}
}

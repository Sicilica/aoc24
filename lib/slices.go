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

func Reduce[T any, U any](in []T, f func(U, T) U, init U) U {
	return ReduceSeq(slices.Values(in), f, init)
}

func ReduceSeq[T any, U any](in iter.Seq[T], f func(U, T) U, init U) U {
	var acc U
	for v := range in {
		acc = f(acc, v)
	}
	return acc
}

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

package grid2d

import (
	"iter"

	"github.com/Sicilica/aoc24/lib"
)

func Count[T any](g Grid[T], f func(T) bool) int {
	return lib.CountSeq(g.Values(), f)
}

func Find[T any](g Grid[T], f func(T) bool) iter.Seq2[Point, T] {
	return func(yield func(Point, T) bool) {
		for p, v := range g.Entries() {
			if f(v) {
				if !yield(p, v) {
					return
				}
			}
		}
	}
}

func FindIndex[T any](g Grid[T], f func(T) bool) iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for p := range Find(g, f) {
			if !yield(p) {
				return
			}
		}
	}
}

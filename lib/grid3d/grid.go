package grid3d

import (
	"iter"
	"maps"
)

type Grid[T any] interface {
	Get(p Point) T
	Set(p Point, val T) bool
	Entries() iter.Seq2[Point, T]
	Values() iter.Seq[T]
}

// SparseGrid is an infinite grid of data stored by point.
type SparseGrid[T any] map[Point]T

var _ Grid[struct{}] = SparseGrid[struct{}](nil)

func (g SparseGrid[T]) Get(p Point) T {
	return g[p]
}

func (g SparseGrid[T]) Set(p Point, val T) bool {
	g[p] = val
	return true
}

func (g SparseGrid[T]) Has(p Point) bool {
	_, ok := g[p]
	return ok
}

func (g SparseGrid[T]) Entries() iter.Seq2[Point, T] {
	return func(yield func(Point, T) bool) {
		for p, v := range g {
			if !yield(p, v) {
				return
			}
		}
	}
}

func (g SparseGrid[T]) Values() iter.Seq[T] {
	return maps.Values(g)
}

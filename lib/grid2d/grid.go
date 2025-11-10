package grid2d

import (
	"iter"
	"maps"

	"github.com/Sicilica/aoc24/lib"
)

type Grid[T any] interface {
	Get(p Point) T
	Set(p Point, val T) bool
	Entries() iter.Seq2[Point, T]
	Values() iter.Seq[T]
}

// FixedGrid contains a 2D array of data in column-major order.
type FixedGrid[T any] [][]T

var _ Grid[struct{}] = FixedGrid[struct{}](nil)

// NewFixed creates a new empty grid with the given fixed size.
func NewFixed[T any](w, h int) FixedGrid[T] {
	grid := make(FixedGrid[T], w)
	for x := range w {
		grid[x] = make([]T, h)
	}
	return grid
}

// Transpose converts data from row-major order to column-major order (or vice-versa).
func Transpose[T any](data [][]T) FixedGrid[T] {
	w := len(data[0])
	h := len(data)
	lib.Assert(lib.Every(data, func(row []T) bool {
		return len(row) == w
	}))

	transposed := make(FixedGrid[T], w)
	for x := range w {
		transposed[x] = make([]T, h)
	}
	for y, row := range data {
		for x, val := range row {
			transposed[x][y] = val
		}
	}

	return transposed
}

func (g FixedGrid[T]) Width() int {
	return len(g)
}

func (g FixedGrid[T]) Height() int {
	if len(g) == 0 {
		return 0
	}
	return len(g[0])
}

func (g FixedGrid[T]) Bounds() Rect {
	return Rect{
		Point{0, 0},
		Point{g.Width(), g.Height()},
	}
}

func (g FixedGrid[T]) Get(p Point) T {
	if g.Bounds().Contains(p) {
		return g[p[0]][p[1]]
	}
	return *new(T)
}

// Set stores the given value at the given point.
// Returns true if the value was stored, or false if the point was out of bounds.
func (g FixedGrid[T]) Set(p Point, val T) bool {
	if g.Bounds().Contains(p) {
		g[p[0]][p[1]] = val
		return true
	}
	return false
}

func (g FixedGrid[T]) Entries() iter.Seq2[Point, T] {
	return func(yield func(Point, T) bool) {
		for x := range g {
			for y := range g[x] {
				if !yield(Point{x, y}, g[x][y]) {
					return
				}
			}
		}
	}
}

func (g FixedGrid[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for x := range g {
			for y := range g[x] {
				if !yield(g[x][y]) {
					return
				}
			}
		}
	}
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

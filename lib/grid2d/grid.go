package grid2d

import (
	"iter"

	"github.com/Sicilica/aoc24/lib"
)

// Grid contains a 2D array of data in column-major order.
type Grid[T any] [][]T

// NewGrid creates a new empty grid with the given size.
func NewGrid[T any](w, h int) Grid[T] {
	grid := make(Grid[T], w)
	for x := range w {
		grid[x] = make([]T, h)
	}
	return grid
}

// Transpose converts data from row-major order to column-major order (or vice-versa).
func Transpose[T any](data [][]T) Grid[T] {
	w := len(data[0])
	h := len(data)
	lib.Assert(lib.Every(data, func(row []T) bool {
		return len(row) == w
	}))

	transposed := make(Grid[T], w)
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

func (g Grid[T]) Width() int {
	return len(g)
}

func (g Grid[T]) Height() int {
	if len(g) == 0 {
		return 0
	}
	return len(g[0])
}

func (g Grid[T]) Bounds() Rect {
	return Rect{
		Point{0, 0},
		Point{g.Width(), g.Height()},
	}
}

func (g Grid[T]) Get(p Point) T {
	if g.Bounds().Contains(p) {
		return g[p[0]][p[1]]
	}
	return *new(T)
}

func (g Grid[T]) Set(p Point, val T) bool {
	if g.Bounds().Contains(p) {
		g[p[0]][p[1]] = val
		return true
	}
	return false
}

func (g Grid[T]) Values() iter.Seq[T] {
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

func (g Grid[T]) Count(f func(T) bool) int {
	return lib.CountSeq(g.Values(), f)
}

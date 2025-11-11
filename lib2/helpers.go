package lib2

import "github.com/Sicilica/aoc24/lib"

// MakeFixedGrid2 creates a new FixedGrid2 with the given dimensions.
func MakeFixedGrid2[Data any](w, h int) FixedGrid2[Data] {
	return MakeFixedGrid2T[int, Data](w, h)
}

// MakeFixedGrid2T is similar to MakeFixedGrid2, but for FixedGrid2t.
func MakeFixedGrid2T[T ~int, Data any](w, h T) FixedGrid2t[T, Data] {
	grid := make(FixedGrid2t[T, Data], w)
	for x := range w {
		grid[x] = make([]Data, h)
	}
	return grid
}

// TransposeT is similar to Transpose, but for FixedGrid2t.
func TransposeT[T ~int, Data any](g FixedGrid2t[T, Data]) FixedGrid2t[T, Data] {
	lib.Assert(g.Valid())
	size := g.Size()
	out := make(FixedGrid2t[T, Data], size.Y())
	for y := range size.Y() {
		out[y] = make([]Data, size.X())
		for x := range size.X() {
			out[y][x] = g[x][y]
		}
	}
	return out
}

// Transpose returns a copy of the grid with the X and Y axes transposed.
//
// Since FixedGrids are stored in column-major order by convention, this is most useful
// for converting between logical and display ordering.
func Transpose[Data any](g FixedGrid2[Data]) FixedGrid2[Data] {
	return TransposeT(g)
}

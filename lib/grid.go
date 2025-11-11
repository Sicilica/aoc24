package lib

import (
	"iter"
	"maps"
)

type Grid[Pos comparable, Data any] interface {
	// All returns an iterator over all the elements in the grid.
	All() iter.Seq2[Pos, Data]

	// Get returns the data at the given position, if it exists.
	Get(Pos) (Data, bool)

	// Set stores the given data at the given position.
	// Returns true if the data was stored successfully, or false otherwise (e.g. because the position is out of bounds).
	Set(Pos, Data) bool
}

// Grid2 is a Grid representing 2D space.
type Grid2[T Real, Data any] interface {
	Grid[Vec2[T], Data]
}

// Grid2i is an alias for Grid2 with int coordinates.
type Grid2i[Data any] = Grid2[int, Data]

// Grid3 is a Grid representing 3D space.
type Grid3[T Real, Data any] interface {
	Grid[Vec3[T], Data]
}

// Grid3i is an alias for Grid3 with int coordinates.
type Grid3i[Data any] = Grid3[int, Data]

// A SparseGrid is a Grid of infinite size that only stores data at a subset of possible positions.
type SparseGrid[Pos comparable, Data any] map[Pos]Data

// SparseGrid implements Grid.
var _ Grid[Vec2[int], any] = SparseGrid[Vec2[int], any]{}

// All returns an iterator over all the elements in the grid.
func (g SparseGrid[V, D]) All() iter.Seq2[V, D] {
	return func(yield func(V, D) bool) {
		for v, d := range g {
			if !yield(v, d) {
				return
			}
		}
	}
}

func (g SparseGrid[V, D]) Copy() SparseGrid[V, D] {
	copy := make(SparseGrid[V, D], len(g))
	maps.Copy(copy, g)
	return copy
}

// Get returns the data at the given position, if it exists.
// The second return value is false if no data is stored at the position.
func (g SparseGrid[V, D]) Get(v V) (D, bool) {
	d, ok := g[v]
	return d, ok
}

// Has returns true if the grid contains the given position.
func (g SparseGrid[V, D]) Has(v V) bool {
	_, ok := g[v]
	return ok
}

// Set stores the given data at the given position.
// Returns true if the data was stored successfully.
func (g SparseGrid[V, D]) Set(v V, d D) bool {
	g[v] = d
	return true
}

// SparseGrid2 is an alias for a SparseGrid in 2D space.
type SparseGrid2[T Real, Data any] = SparseGrid[Vec2[T], Data]

// SparseGrid2 implements Grid2.
var _ Grid2[int, any] = SparseGrid2[int, any]{}

// SparseGrid2i is an alias for a SparseGrid in 2D space with int coordinates.
type SparseGrid2i[Data any] = SparseGrid2[int, Data]

// SparseGrid3 is an alias for a SparseGrid in 3D space.
type SparseGrid3[T Real, Data any] = SparseGrid[Vec3[T], Data]

// SparseGrid3 implements Grid3.
var _ Grid3[int, any] = SparseGrid3[int, any]{}

// SparseGrid3i is an alias for a SparseGrid in 3D space with int coordinates.
type SparseGrid3i[Data any] = SparseGrid3[int, Data]

// A FixedGrid is a Grid of a finite, fixed size, where all positions within the grid bounds are always valid.
type FixedGrid[Pos boxVec[Pos], Data any] interface {
	Grid[Pos, Data]

	// Bounds returns the bounds of the grid.
	Bounds() Box[Pos]

	// Size returns the size of the grid.
	Size() Pos

	// Valid returns true if the grid is valid, or false otherwise.
	// Because a FixedGrid is an N-dimensional array, it is valid if
	// the length of each dimension is consistent for each sub-array.
	Valid() bool
}

// FixedGrid2t is similar to FixedGrid2, but with typed coordinates.
type FixedGrid2t[T ~int, Data any] [][]Data

// FixedGrid2t implements Grid2.
var _ Grid2[int, any] = FixedGrid2t[int, any]{}

// FixedGrid2t implement FixedGrid.
var _ FixedGrid[Vec2[int], any] = FixedGrid2t[int, any]{}

// FixedGrid2 is an array representing 2D space.
type FixedGrid2[Data any] = FixedGrid2t[int, Data]

// All returns an iterator over all the elements in the grid.
func (g FixedGrid2t[T, Data]) All() iter.Seq2[Vec2[T], Data] {
	return func(yield func(Vec2[T], Data) bool) {
		for x := range len(g) {
			for y := range len(g[x]) {
				if !yield(Vec2[T]{T(x), T(y)}, g[x][y]) {
					return
				}
			}
		}
	}
}

// Bounds returns the bounds of the grid.
func (g FixedGrid2t[T, Data]) Bounds() Box2[T] {
	return Box2[T]{Vec2[T]{0, 0}, g.Size()}
}

// Get returns the data at the given position, if it is within the grid bounds.
func (g FixedGrid2t[T, Data]) Get(v Vec2[T]) (Data, bool) {
	if !g.Bounds().Contains(v) {
		var zero Data
		return zero, false
	}
	return g[v[0]][v[1]], true
}

// Set stores the given data at the given position.
// Returns true if the data was stored successfully, or false otherwise (e.g. because the position is out of bounds).
func (g FixedGrid2t[T, Data]) Set(v Vec2[T], d Data) bool {
	if !g.Bounds().Contains(v) {
		return false
	}
	g[v[0]][v[1]] = d
	return true
}

// Size returns the size of the grid.
func (g FixedGrid2t[T, Data]) Size() Vec2[T] {
	x := len(g)
	y := 0
	if x > 0 {
		y = len(g[0])
	}
	return Vec2[T]{T(x), T(y)}
}

// Valid returns true if the grid is valid, or false otherwise.
// Because a FixedGrid is an N-dimensional array, it is valid if
// the length of each dimension is consistent for each sub-array.
func (g FixedGrid2t[T, Data]) Valid() bool {
	size := g.Size()
	for x := range size.X() {
		if len(g[x]) != int(size.Y()) {
			return false
		}
	}
	return true
}

// FixedGrid3t is similar to FixedGrid3, but with typed coordinates.
type FixedGrid3t[T ~int, Data any] [][][]Data

// FixedGrid3t implements Grid3.
var _ Grid3[int, any] = FixedGrid3t[int, any]{}

// FixedGrid3t implement FixedGrid.
var _ FixedGrid[Vec3[int], any] = FixedGrid3t[int, any]{}

// FixedGrid3 is an array representing 3D space.
type FixedGrid3[Data any] = FixedGrid3t[int, Data]

// All returns an iterator over all the elements in the grid.
func (g FixedGrid3t[T, Data]) All() iter.Seq2[Vec3[T], Data] {
	return func(yield func(Vec3[T], Data) bool) {
		for x := range len(g) {
			for y := range len(g[x]) {
				for z := range len(g[x][y]) {
					if !yield(Vec3[T]{T(x), T(y), T(z)}, g[x][y][z]) {
						return
					}
				}
			}
		}
	}
}

// Bounds returns the bounds of the grid.
func (g FixedGrid3t[T, Data]) Bounds() Box3[T] {
	return Box3[T]{Vec3[T]{0, 0, 0}, g.Size()}
}

// Get returns the data at the given position, if it is within the grid bounds.
func (g FixedGrid3t[T, Data]) Get(v Vec3[T]) (Data, bool) {
	if !g.Bounds().Contains(v) {
		var zero Data
		return zero, false
	}
	return g[v[0]][v[1]][v[2]], true
}

// Set stores the given data at the given position.
// Returns true if the data was stored successfully, or false otherwise (e.g. because the position is out of bounds).
func (g FixedGrid3t[T, Data]) Set(v Vec3[T], d Data) bool {
	if !g.Bounds().Contains(v) {
		return false
	}
	g[v[0]][v[1]][v[2]] = d
	return true
}

// Size returns the size of the grid.
func (g FixedGrid3t[T, Data]) Size() Vec3[T] {
	x := len(g)
	y := 0
	if x > 0 {
		y = len(g[0])
	}
	z := 0
	if y > 0 {
		z = len(g[0][0])
	}
	return Vec3[T]{T(x), T(y), T(z)}
}

// Valid returns true if the grid is valid, or false otherwise.
// Because a FixedGrid is an N-dimensional array, it is valid if
// the length of each dimension is consistent for each sub-array.
func (g FixedGrid3t[T, Data]) Valid() bool {
	size := g.Size()
	for x := range size.X() {
		if len(g[x]) != int(size.Y()) {
			return false
		}
		for y := range size.Y() {
			if len(g[x][y]) != int(size.Z()) {
				return false
			}
		}
	}
	return true
}

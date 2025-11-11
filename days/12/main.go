package main

import (
	_ "embed"
	"slices"
	"strings"

	"github.com/Sicilica/aoc24/lib"
)

//go:embed input.txt
var rawInput string

func main() {
	lib.Day(
		input,
		part1,
		part2,
	)
}

func input() lib.FixedGrid2[byte] {
	return lib.Transpose(slices.Collect(lib.Map(strings.Lines(rawInput), func(l string) []byte {
		return lib.MapSlice(strings.Split(strings.TrimSpace(l), ""), func(s string) byte {
			lib.Assert(len(s) == 1)
			return s[0]
		})
	})))
}

func part1(grid lib.FixedGrid2[byte]) int {
	totalCost := 0
	for _, r := range gridToRegions(grid) {
		area := len(r.Plots)
		peri := perimeter(r.Plots)
		totalCost += area * peri
	}
	return totalCost
}

func part2(grid lib.FixedGrid2[byte]) int {
	totalCost := 0
	for _, r := range gridToRegions(grid) {
		area := len(r.Plots)
		sides := sides(r.Plots)
		totalCost += area * sides
	}
	return totalCost
}

type Region struct {
	Plant byte
	Plots lib.SparseGrid2i[struct{}]
}

var dirs = []lib.Vec2i{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func gridToRegions(grid lib.Grid2i[byte]) []Region {
	var regions []Region
	explored := make(lib.SparseGrid2i[struct{}])
	var queue []lib.Vec2i
	for plot, plant := range grid.All() {
		if explored.Has(plot) {
			continue
		}

		clear(queue)
		explored.Set(plot, struct{}{})
		queue = append(queue, plot)
		regionPlots := make(lib.SparseGrid2i[struct{}])
		for len(queue) > 0 {
			p := queue[len(queue)-1]
			queue = queue[:len(queue)-1]

			regionPlots.Set(p, struct{}{})

			for _, dir := range dirs {
				next := p.Plus(dir)
				if lib.IgnoreOK(grid.Get(next)) == plant && !explored.Has(next) {
					explored.Set(next, struct{}{})
					queue = append(queue, next)
				}
			}
		}
		regions = append(regions, Region{
			Plant: plant,
			Plots: regionPlots,
		})
	}
	return regions
}

func perimeter[T any](grid lib.SparseGrid2i[T]) int {
	sum := 0
	for p := range grid.All() {
		for _, dir := range dirs {
			if !grid.Has(p.Plus(dir)) {
				sum++
			}
		}
	}
	return sum
}

func sides[T any](grid lib.SparseGrid2i[T]) int {
	sum := 0
	explored := make(lib.SparseGrid2i[struct{}])
	var queue []lib.Vec2i
	for _, dir := range dirs {
		clear(explored)
		for p := range grid.All() {
			if explored.Has(p) {
				continue
			}
			explored.Set(p, struct{}{})

			if grid.Has(p.Plus(dir)) {
				continue
			}

			axis := axisForDir(dir)
			clear(queue)
			queue = append(queue, p)
			for len(queue) > 0 {
				p := queue[len(queue)-1]
				queue = queue[:len(queue)-1]

				left := p.Minus(axis)
				if grid.Has(left) && !explored.Has(left) && !grid.Has(left.Plus(dir)) {
					explored.Set(left, struct{}{})
					queue = append(queue, left)
				}

				right := p.Plus(axis)
				if grid.Has(right) && !explored.Has(right) && !grid.Has(right.Plus(dir)) {
					explored.Set(right, struct{}{})
					queue = append(queue, right)
				}
			}
			sum += 1
		}
	}
	return sum
}

func axisForDir(dir lib.Vec2i) lib.Vec2i {
	if dir.X() == 0 {
		return lib.Vec2i{1, 0}
	}
	return lib.Vec2i{0, 1}
}

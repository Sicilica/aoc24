package main

import (
	_ "embed"
	"slices"
	"strings"

	"github.com/Sicilica/aoc24/lib"
	"github.com/Sicilica/aoc24/lib/grid2d"
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

func input() grid2d.FixedGrid[byte] {
	return grid2d.Transpose(slices.Collect(lib.MapSeq(strings.Lines(rawInput), func(l string) []byte {
		return lib.Map(strings.Split(strings.TrimSpace(l), ""), func(s string) byte {
			lib.Assert(len(s) == 1)
			return s[0]
		})
	})))
}

func part1(grid grid2d.FixedGrid[byte]) int {
	totalCost := 0
	for _, r := range gridToRegions(grid) {
		area := len(r.Plots)
		peri := perimeter(r.Plots)
		totalCost += area * peri
	}
	return totalCost
}

func part2(grid grid2d.FixedGrid[byte]) int {
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
	Plots grid2d.SparseGrid[struct{}]
}

var dirs = []grid2d.Point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func gridToRegions(grid grid2d.Grid[byte]) []Region {
	var regions []Region
	explored := make(grid2d.SparseGrid[struct{}])
	var queue []grid2d.Point
	for plot, plant := range grid.Entries() {
		if explored.Has(plot) {
			continue
		}

		clear(queue)
		explored.Set(plot, struct{}{})
		queue = append(queue, plot)
		regionPlots := make(grid2d.SparseGrid[struct{}])
		for len(queue) > 0 {
			p := queue[len(queue)-1]
			queue = queue[:len(queue)-1]

			regionPlots.Set(p, struct{}{})

			for _, dir := range dirs {
				next := p.Plus(dir)
				if grid.Get(next) == plant && !explored.Has(next) {
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

func perimeter[T any](grid grid2d.SparseGrid[T]) int {
	sum := 0
	for p := range grid.Entries() {
		for _, dir := range dirs {
			if !grid.Has(p.Plus(dir)) {
				sum++
			}
		}
	}
	return sum
}

func sides[T any](grid grid2d.SparseGrid[T]) int {
	sum := 0
	explored := make(grid2d.SparseGrid[struct{}])
	var queue []grid2d.Point
	for _, dir := range dirs {
		clear(explored)
		for p := range grid.Entries() {
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

func axisForDir(dir grid2d.Point) grid2d.Point {
	if dir.X() == 0 {
		return grid2d.Point{1, 0}
	}
	return grid2d.Point{0, 1}
}

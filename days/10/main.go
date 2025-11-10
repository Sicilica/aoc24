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

func input() grid2d.FixedGrid[int] {
	return grid2d.Transpose(slices.Collect(lib.MapSeq(strings.Lines(rawInput), func(l string) []int {
		return lib.Map(strings.Split(strings.TrimSpace(l), ""), lib.Atoi)
	})))
}

func part1(grid grid2d.FixedGrid[int]) int {
	trailheads := grid2d.FindIndex(grid, func(height int) bool {
		return height == 0
	})

	dirs := []grid2d.Point{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	sum := 0
	peaks := grid2d.SparseGrid[struct{}]{}
	explored := grid2d.SparseGrid[struct{}]{}
	var queue []grid2d.Point
	for t := range trailheads {
		clear(peaks)
		clear(explored)
		clear(queue)

		queue = append(queue, t)

		score := 0
		for len(queue) > 0 {
			p := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			height := grid.Get(p)

			if height >= 9 {
				score++
				continue
			}

			for _, dir := range dirs {
				next := p.Plus(dir)
				if grid.Get(next) == height+1 && !explored.Has(next) {
					explored.Set(next, struct{}{})
					queue = append(queue, next)
				}
			}
		}
		sum += score
	}
	return sum
}

func part2(grid grid2d.FixedGrid[int]) int {
	trailheads := grid2d.FindIndex(grid, func(height int) bool {
		return height == 0
	})

	dirs := []grid2d.Point{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	sum := 0
	peaks := grid2d.SparseGrid[struct{}]{}
	var queue []grid2d.Point
	for t := range trailheads {
		clear(peaks)
		clear(queue)

		queue = append(queue, t)

		rating := 0
		for len(queue) > 0 {
			p := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			height := grid.Get(p)

			if height >= 9 {
				rating++
				continue
			}

			for _, dir := range dirs {
				next := p.Plus(dir)
				if grid.Get(next) == height+1 {
					queue = append(queue, next)
				}
			}
		}
		sum += rating
	}
	return sum
}

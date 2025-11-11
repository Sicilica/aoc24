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

func input() lib.FixedGrid2[int] {
	return lib.Transpose(slices.Collect(lib.Map(strings.Lines(rawInput), func(l string) []int {
		return lib.MapSlice(strings.Split(strings.TrimSpace(l), ""), lib.Atoi)
	})))
}

func part1(grid lib.FixedGrid2[int]) int {
	trailheads := lib.Indices(grid.All(), 0)

	dirs := []lib.Vec2i{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	sum := 0
	peaks := lib.SparseGrid2i[struct{}]{}
	explored := lib.SparseGrid2i[struct{}]{}
	var queue []lib.Vec2i
	for t := range trailheads {
		clear(peaks)
		clear(explored)
		clear(queue)

		queue = append(queue, t)

		score := 0
		for len(queue) > 0 {
			p := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			height := lib.OK(grid.Get(p))

			if height >= 9 {
				score++
				continue
			}

			for _, dir := range dirs {
				next := p.Plus(dir)
				if lib.IgnoreOK(grid.Get(next)) == height+1 && !explored.Has(next) {
					explored.Set(next, struct{}{})
					queue = append(queue, next)
				}
			}
		}
		sum += score
	}
	return sum
}

func part2(grid lib.FixedGrid2[int]) int {
	trailheads := lib.Indices(grid.All(), 0)

	dirs := []lib.Vec2i{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	sum := 0
	peaks := lib.SparseGrid2i[struct{}]{}
	var queue []lib.Vec2i
	for t := range trailheads {
		clear(peaks)
		clear(queue)

		queue = append(queue, t)

		rating := 0
		for len(queue) > 0 {
			p := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			height := lib.OK(grid.Get(p))

			if height >= 9 {
				rating++
				continue
			}

			for _, dir := range dirs {
				next := p.Plus(dir)
				if lib.IgnoreOK(grid.Get(next)) == height+1 {
					queue = append(queue, next)
				}
			}
		}
		sum += rating
	}
	return sum
}

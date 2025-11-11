package main

import (
	_ "embed"
	"slices"
	"strings"

	"github.com/Sicilica/aoc24/lib"
	"github.com/Sicilica/aoc24/lib2"
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

func input() lib2.FixedGrid2[int] {
	return lib2.Transpose(slices.Collect(lib.MapSeq(strings.Lines(rawInput), func(l string) []int {
		return lib.Map(strings.Split(strings.TrimSpace(l), ""), lib.Atoi)
	})))
}

func part1(grid lib2.FixedGrid2[int]) int {
	trailheads := lib2.Indices(grid.All(), 0)

	dirs := []lib2.Vec2i{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	sum := 0
	peaks := lib2.SparseGrid2i[struct{}]{}
	explored := lib2.SparseGrid2i[struct{}]{}
	var queue []lib2.Vec2i
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

func part2(grid lib2.FixedGrid2[int]) int {
	trailheads := lib2.Indices(grid.All(), 0)

	dirs := []lib2.Vec2i{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	sum := 0
	peaks := lib2.SparseGrid2i[struct{}]{}
	var queue []lib2.Vec2i
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

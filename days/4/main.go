package main

import (
	_ "embed"
	"iter"
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

var dirs = []grid2d.Point{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}

func part1(input grid2d.FixedGrid[byte]) int {
	count := 0
	for s := range search(input, 'X') {
		count += lib.Count(dirs, func(dir grid2d.Point) bool {
			return matches(input, s, dir, "XMAS")
		})
	}
	return count
}

func part2(input grid2d.FixedGrid[byte]) int {
	count := 0
	for s := range search(input, 'A') {
		if s.OnEdge(input.Bounds()) {
			// too close to the edge; skip this one
			continue
		}

		leftDiag := matches(input, s.Plus(grid2d.Point{-1, -1}), grid2d.Point{1, 1}, "MAS") || matches(input, s.Plus(grid2d.Point{1, 1}), grid2d.Point{-1, -1}, "MAS")
		rightDiag := matches(input, s.Plus(grid2d.Point{-1, 1}), grid2d.Point{1, -1}, "MAS") || matches(input, s.Plus(grid2d.Point{1, -1}), grid2d.Point{-1, 1}, "MAS")
		if leftDiag && rightDiag {
			count++
		}
	}
	return count
}

// search finds each occurrence of target anywhere in the data.
func search(data grid2d.FixedGrid[byte], target byte) iter.Seq[grid2d.Point] {
	return grid2d.FindIndex(data, func(val byte) bool {
		return val == target
	})
}

func matches(data grid2d.FixedGrid[byte], start, dir grid2d.Point, target string) bool {
	pos := start
	for i := 0; i < len(target); i++ {
		if !data.Bounds().Contains(pos) {
			return false
		}
		if data.Get(pos) != target[i] {
			return false
		}
		pos = pos.Plus(dir)
	}
	return true
}

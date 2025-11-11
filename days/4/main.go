package main

import (
	_ "embed"
	"iter"
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

var dirs = []lib.Vec2i{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}

func part1(input lib.FixedGrid2[byte]) int {
	count := 0
	for s := range search(input, 'X') {
		count += lib.CountFunc(slices.Values(dirs), func(dir lib.Vec2i) bool {
			return matches(input, s, dir, "XMAS")
		})
	}
	return count
}

func part2(input lib.FixedGrid2[byte]) int {
	count := 0
	for s := range search(input, 'A') {
		// if s.OnEdge(input.Bounds()) {
		// 	// too close to the edge; skip this one
		// 	continue
		// }

		leftDiag := matches(input, s.Plus(lib.Vec2i{-1, -1}), lib.Vec2i{1, 1}, "MAS") || matches(input, s.Plus(lib.Vec2i{1, 1}), lib.Vec2i{-1, -1}, "MAS")
		rightDiag := matches(input, s.Plus(lib.Vec2i{-1, 1}), lib.Vec2i{1, -1}, "MAS") || matches(input, s.Plus(lib.Vec2i{1, -1}), lib.Vec2i{-1, 1}, "MAS")
		if leftDiag && rightDiag {
			count++
		}
	}
	return count
}

// search finds each occurrence of target anywhere in the data.
func search(data lib.FixedGrid2[byte], target byte) iter.Seq[lib.Vec2i] {
	return lib.Indices(data.All(), target)
}

func matches(data lib.FixedGrid2[byte], start, dir lib.Vec2i, target string) bool {
	pos := start
	for i := 0; i < len(target); i++ {
		if !data.Bounds().Contains(pos) {
			return false
		}
		if lib.IgnoreOK(data.Get(pos)) != target[i] {
			return false
		}
		pos = pos.Plus(dir)
	}
	return true
}

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

func input() (lib.FixedGrid2[bool], lib.Vec2i, lib.Vec2i) {
	start := lib.Vec2i{-1, -1}
	end := lib.Vec2i{-1, -1}

	y := -1
	data := lib.Transpose(slices.Collect(lib.Map(strings.Lines(rawInput), func(l string) []bool {
		y++
		x := -1
		return lib.MapSlice(strings.Split(strings.TrimSpace(l), ""), func(s string) bool {
			x++
			lib.Assert(len(s) == 1)
			switch s[0] {
			case '#':
				return true
			case '.':
				return false
			case 'S':
				lib.Assert(start[0] == -1)
				start = lib.Vec2i{x, y}
				return false
			case 'E':
				lib.Assert(end[0] == -1)
				end = lib.Vec2i{x, y}
				return false
			default:
				panic("unrecognized char in grid")
			}
		})
	})))

	lib.Assert(start[0] != -1)
	lib.Assert(end[0] != -1)

	return data, start, end
}

func part1(grid lib.Grid2i[bool], start, goal lib.Vec2i) int {
	return 0
}

func part2(grid lib.Grid2i[bool], start, goal lib.Vec2i) int {
	return 0
}

package main

import (
	_ "embed"
	"slices"
	"strings"

	"github.com/Sicilica/aoc24/lib"
	"github.com/Sicilica/aoc24/lib/grid2d"
	"github.com/Sicilica/aoc24/lib/grid3d"
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

func input() (grid2d.FixedGrid[bool], grid2d.Point) {
	y := -1
	startPos := grid2d.Point{-1, -1}
	data := grid2d.Transpose(slices.Collect(lib.MapSeq(strings.Lines(rawInput), func(l string) []bool {
		y++
		x := -1
		return lib.Map(strings.Split(strings.TrimSpace(l), ""), func(s string) bool {
			x++
			lib.Assert(len(s) == 1)

			switch s[0] {
			case '.':
				return false
			case '#':
				return true
			case '^':
				lib.Assert(startPos[0] == -1)
				startPos = grid2d.Point{x, y}
				return false
			default:
				panic("invalid input")
			}
		})
	})))
	lib.Assert(startPos[0] != -1)
	return data, startPos
}

func part1(data grid2d.FixedGrid[bool], startPos grid2d.Point) int {
	visited := make(grid2d.SparseGrid[struct{}])

	dirs := []grid2d.Point{
		{0, -1}, // start facing up
		{1, 0},  // turn to the right
		{0, 1},
		{-1, 0},
	}

	pos := startPos
	dirIndex := 0
	for {
		visited.Set(pos, struct{}{})

		pos = pos.Plus(dirs[dirIndex])
		if !data.Bounds().Contains(pos) {
			break
		}
		if data.Get(pos) {
			pos = pos.Minus(dirs[dirIndex])
			dirIndex = (dirIndex + 1) % len(dirs)
		}
	}

	return len(visited)
}

func part2(data grid2d.FixedGrid[bool], startPos grid2d.Point) int {
	dirs := []grid2d.Point{
		{0, -1}, // start facing up
		{1, 0},  // turn to the right
		{0, 1},
		{-1, 0},
	}

	count := 0
	visitedWithDir := make(grid3d.SparseGrid[struct{}])
	for ox := range data.Width() {
		for oy := range data.Height() {
			obs := grid2d.Point{ox, oy}
			// Don't place on top of the guard
			if obs.Equals(startPos) {
				continue
			}

			clear(visitedWithDir)
			pos := startPos
			dirIndex := 0
			var looped bool
			for {
				visit := grid3d.Point{pos.X(), pos.Y(), dirIndex}
				if visitedWithDir.Has(visit) {
					looped = true
					break
				}
				visitedWithDir.Set(visit, struct{}{})

				pos = pos.Plus(dirs[dirIndex])
				if !data.Bounds().Contains(pos) {
					looped = false
					break
				}
				if data.Get(pos) || pos.Equals(obs) {
					pos = pos.Minus(dirs[dirIndex])
					dirIndex = (dirIndex + 1) % len(dirs)
				}
			}

			if looped {
				count++
			}
		}
	}

	return count
}

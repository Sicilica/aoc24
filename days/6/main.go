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

func input() (lib2.FixedGrid2[bool], lib2.Vec2i) {
	y := -1
	startPos := lib2.Vec2i{-1, -1}
	data := lib2.Transpose(slices.Collect(lib.MapSeq(strings.Lines(rawInput), func(l string) []bool {
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
				startPos = lib2.Vec2i{x, y}
				return false
			default:
				panic("invalid input")
			}
		})
	})))
	lib.Assert(startPos[0] != -1)
	return data, startPos
}

func part1(data lib2.FixedGrid2[bool], startPos lib2.Vec2i) int {
	visited := make(lib2.SparseGrid2i[struct{}])

	dirs := []lib2.Vec2i{
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
		if lib.IgnoreOK(data.Get(pos)) {
			pos = pos.Minus(dirs[dirIndex])
			dirIndex = (dirIndex + 1) % len(dirs)
		}
	}

	return len(visited)
}

func part2(data lib2.FixedGrid2[bool], startPos lib2.Vec2i) int {
	dirs := []lib2.Vec2i{
		{0, -1}, // start facing up
		{1, 0},  // turn to the right
		{0, 1},
		{-1, 0},
	}

	count := 0
	visitedWithDir := make(lib2.SparseGrid3i[struct{}])
	for ox := range data.Size().X() {
		for oy := range data.Size().Y() {
			obs := lib2.Vec2i{ox, oy}
			// Don't place on top of the guard
			if obs.Equals(startPos) {
				continue
			}

			clear(visitedWithDir)
			pos := startPos
			dirIndex := 0
			var looped bool
			for {
				visit := lib2.Vec3i{pos.X(), pos.Y(), dirIndex}
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
				if lib.IgnoreOK(data.Get(pos)) || pos.Equals(obs) {
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

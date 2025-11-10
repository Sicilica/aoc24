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

func input() (grid2d.Grid[byte], map[byte][]grid2d.Point) {
	grid := grid2d.Transpose(slices.Collect(lib.MapSeq(strings.Lines(rawInput), func(l string) []byte {
		return lib.Map(strings.Split(strings.TrimSpace(l), ""), func(s string) byte {
			lib.Assert(len(s) == 1)

			if s[0] == '.' {
				return 0
			}
			return s[0]
		})
	})))

	antennas := make(map[byte][]grid2d.Point)
	for x := range grid.Width() {
		for y := range grid.Height() {
			if grid[x][y] != 0 {
				antennas[grid[x][y]] = append(antennas[grid[x][y]], grid2d.Point{x, y})
			}
		}
	}
	return grid, antennas
}

func part1(grid grid2d.Grid[byte], antennas map[byte][]grid2d.Point) int {
	antinodes := grid2d.NewGrid[bool](grid.Width(), grid.Height())

	for _, antennas := range antennas {
		for a, b := range lib.Pairs(antennas) {
			delta := a.To(b)
			antinodes.Set(b.Plus(delta), true)
			antinodes.Set(a.Minus(delta), true)
		}
	}

	return antinodes.Count(func(val bool) bool {
		return val
	})
}

func part2(grid grid2d.Grid[byte], antennas map[byte][]grid2d.Point) int {
	antinodes := grid2d.NewGrid[bool](grid.Width(), grid.Height())

	for _, antennas := range antennas {
		for a, b := range lib.Pairs(antennas) {
			delta := a.To(b)
			for i := 0; ; i++ {
				if !antinodes.Set(b.Plus(delta.Times(i)), true) {
					break
				}
			}
			for i := 0; ; i++ {
				if !antinodes.Set(a.Minus(delta.Times(i)), true) {
					break
				}
			}
		}
	}

	return antinodes.Count(func(val bool) bool {
		return val
	})
}

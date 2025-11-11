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

func input() (lib.FixedGrid2[byte], map[byte][]lib.Vec2i) {
	grid := lib.Transpose(slices.Collect(lib.Map(strings.Lines(rawInput), func(l string) []byte {
		return lib.MapSlice(strings.Split(strings.TrimSpace(l), ""), func(s string) byte {
			lib.Assert(len(s) == 1)

			if s[0] == '.' {
				return 0
			}
			return s[0]
		})
	})))

	antennas := make(map[byte][]lib.Vec2i)
	for x := range grid.Size().X() {
		for y := range grid.Size().Y() {
			if grid[x][y] != 0 {
				antennas[grid[x][y]] = append(antennas[grid[x][y]], lib.Vec2i{x, y})
			}
		}
	}
	return grid, antennas
}

func part1(grid lib.FixedGrid2[byte], antennas map[byte][]lib.Vec2i) int {
	antinodes := lib.MakeFixedGrid2[bool](grid.Size().X(), grid.Size().Y())

	for _, antennas := range antennas {
		for a, b := range lib.Pairs(antennas) {
			delta := b.Minus(a)
			antinodes.Set(b.Plus(delta), true)
			antinodes.Set(a.Minus(delta), true)
		}
	}

	return lib.Count(lib.Indices(antinodes.All(), true))
}

func part2(grid lib.FixedGrid2[byte], antennas map[byte][]lib.Vec2i) int {
	antinodes := lib.MakeFixedGrid2[bool](grid.Size().X(), grid.Size().Y())

	for _, antennas := range antennas {
		for a, b := range lib.Pairs(antennas) {
			delta := b.Minus(a)
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

	return lib.Count(lib.Indices(antinodes.All(), true))
}

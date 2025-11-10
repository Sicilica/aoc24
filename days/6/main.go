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

func input() ([][]bool, [2]int) {
	x := -1
	startPos := [2]int{-1, -1}
	data := slices.Collect(lib.MapSeq(strings.Lines(rawInput), func(l string) []bool {
		x++
		y := -1
		return lib.Map(strings.Split(strings.TrimSpace(l), ""), func(s string) bool {
			y++
			lib.Assert(len(s) == 1)

			switch s[0] {
			case '.':
				return false
			case '#':
				return true
			case '^':
				lib.Assert(startPos[0] == -1)
				startPos = [2]int{x, y}
				return false
			default:
				panic("invalid input")
			}
		})
	}))
	lib.Assert(startPos[0] != -1)
	h := len(data[0])
	lib.Assert(lib.Every(data, func(row []bool) bool {
		return len(row) == h
	}))
	return data, startPos
}

func part1(data [][]bool, startPos [2]int) int {
	w := len(data)
	h := len(data[0])

	visited := make(map[[2]int]struct{})

	// MEMO: x/y is transposed to make parsing simple, so this will be a bit different than the prompt
	dirs := [][2]int{
		{-1, 0}, // start facing left
		{0, 1},  // turn to the left
		{1, 0},
		{0, -1},
	}

	x := startPos[0]
	y := startPos[1]
	dirIndex := 0
	for {
		visited[[2]int{x, y}] = struct{}{}

		x += dirs[dirIndex][0]
		y += dirs[dirIndex][1]
		if x < 0 || x >= w || y < 0 || y >= h {
			break
		}
		if data[x][y] {
			x -= dirs[dirIndex][0]
			y -= dirs[dirIndex][1]
			dirIndex = (dirIndex + 1) % len(dirs)
		}
	}

	return len(visited)
}

func part2(data [][]bool, startPos [2]int) int {
	w := len(data)
	h := len(data[0])

	// MEMO: x/y is transposed to make parsing simple, so this will be a bit different than the prompt
	dirs := [][2]int{
		{-1, 0}, // start facing left
		{0, 1},  // turn to the left
		{1, 0},
		{0, -1},
	}

	count := 0
	visitedWithDir := make(map[[3]int]struct{})
	for ox := range w {
		for oy := range h {
			// Don't place on top of the guard
			if ox == startPos[0] && oy == startPos[1] {
				continue
			}

			clear(visitedWithDir)
			x := startPos[0]
			y := startPos[1]
			dirIndex := 0
			var looped bool
			for {
				if _, ok := visitedWithDir[[3]int{x, y, dirIndex}]; ok {
					looped = true
					break
				}
				visitedWithDir[[3]int{x, y, dirIndex}] = struct{}{}

				x += dirs[dirIndex][0]
				y += dirs[dirIndex][1]
				if x < 0 || x >= w || y < 0 || y >= h {
					looped = false
					break
				}
				if data[x][y] || (ox == x && oy == y) {
					x -= dirs[dirIndex][0]
					y -= dirs[dirIndex][1]
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

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

func input() ([][]byte, int, int) {
	data := slices.Collect(lib.MapSeq(strings.Lines(rawInput), func(l string) []byte {
		return lib.Map(strings.Split(strings.TrimSpace(l), ""), func(s string) byte {
			lib.Assert(len(s) == 1)
			return s[0]
		})
	}))
	// TODO: we transpose x/y for convenience, but that shouldn't matter
	w := len(data)
	h := len(data[0])
	lib.Assert(lib.Every(data, func(row []byte) bool {
		return len(row) == h
	}))
	return data, w, h
}

var dirs = [][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}

func part1(input [][]byte, w, h int) int {
	count := 0
	for sx, sy := range search(input, 'X') {
		start := [2]int{sx, sy}
		count += lib.Count(dirs, func(dir [2]int) bool {
			return matches(input, start, dir, "XMAS")
		})
	}
	return count
}

func part2(input [][]byte, w, h int) int {
	count := 0
	for sx, sy := range search(input, 'A') {
		if sx == 0 || sy == 0 || sx == w-1 || sy == h-1 {
			// too close to the edge; skip this one
			continue
		}

		leftDiag := matches(input, [2]int{sx - 1, sy - 1}, [2]int{1, 1}, "MAS") || matches(input, [2]int{sx + 1, sy + 1}, [2]int{-1, -1}, "MAS")
		rightDiag := matches(input, [2]int{sx - 1, sy + 1}, [2]int{1, -1}, "MAS") || matches(input, [2]int{sx + 1, sy - 1}, [2]int{-1, 1}, "MAS")
		if leftDiag && rightDiag {
			count++
		}
	}
	return count
}

// search finds each occurrence of target anywhere in the data.
func search(data [][]byte, target byte) iter.Seq2[int, int] {
	w := len(data)
	h := len(data[0])
	return func(yield func(int, int) bool) {
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				if data[x][y] == target {
					if !yield(x, y) {
						return
					}
				}
			}
		}
	}
}

func matches(data [][]byte, start, dir [2]int, target string) bool {
	w := len(data)
	h := len(data[0])
	x, y := start[0], start[1]
	for i := 0; i < len(target); i++ {
		if x < 0 || x >= w || y < 0 || y >= h {
			return false
		}
		if data[x][y] != target[i] {
			return false
		}
		x += dir[0]
		y += dir[1]
	}
	return true
}

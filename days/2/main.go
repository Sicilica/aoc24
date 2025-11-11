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

func input() [][]int {
	return slices.Collect(lib.Map(strings.Lines(rawInput), func(l string) []int {
		return lib.MapSlice(strings.Split(l, " "), lib.Atoi)
	}))
}

func part1(reports [][]int) int {
	return lib.CountFunc(slices.Values(reports), safe)
}

func part2(reports [][]int) int {
	return lib.CountFunc(slices.Values(reports), func(r []int) bool {
		subReport := make([]int, len(r)-1)
		for i := range r {
			copy(subReport[0:], r[0:i])
			copy(subReport[i:], r[i+1:])
			if safe(subReport) {
				return true
			}
		}
		return false
	})
}

func safe(r []int) bool {
	// Two conditions must be true:
	// - strictly increasing or strictly decreasing
	// - all differences are in [1, 3]
	deltas := deltas(r)

	if slices.ContainsFunc(deltas, func(d int) bool {
		ad := lib.Abs(d)
		return ad < 1 || ad > 3
	}) {
		return false
	}

	return lib.Every(slices.Values(deltas), func(d int) bool {
		return d >= 0
	}) || lib.Every(slices.Values(deltas), func(d int) bool {
		return d <= 0
	})
}

func deltas(r []int) []int {
	deltas := make([]int, len(r)-1)
	for i := 1; i < len(r); i++ {
		deltas[i-1] = r[i] - r[i-1]
	}
	return deltas
}

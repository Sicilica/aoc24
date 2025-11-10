package main

import (
	_ "embed"
	"regexp"
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

func input() ([]int, []int) {
	lineReg := regexp.MustCompile(`(\d+)\s+(\d+)`)

	var left []int
	var right []int
	for l := range strings.Lines(rawInput) {
		m := lib.Match(lineReg, l)
		left = append(left, lib.Atoi(m[1]))
		right = append(right, lib.Atoi(m[2]))
	}

	return left, right
}

func part1(left, right []int) int {
	// Technically we mutate the inputs here, but part2 doesn't care.
	slices.Sort(left)
	slices.Sort(right)

	totalDistance := 0
	for i := range left {
		totalDistance += lib.Abs(left[i] - right[i])
	}

	return totalDistance
}

func part2(left, right []int) int {
	rightCounts := make(map[int]int)
	for _, r := range right {
		rightCounts[r]++
	}

	similarity := 0
	for _, l := range left {
		similarity += l * rightCounts[l]
	}

	return similarity
}

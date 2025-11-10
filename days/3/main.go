package main

import (
	_ "embed"
	"regexp"

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

func input() string {
	return rawInput
}

func part1(input string) int {
	sum := 0
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	for _, m := range r.FindAllStringSubmatch(input, -1) {
		sum += lib.Atoi(m[1]) * lib.Atoi(m[2])
	}
	return sum
}

func part2(input string) int {
	enabled := true
	sum := 0
	r := regexp.MustCompile(`(?:mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)
	for _, m := range r.FindAllStringSubmatch(input, -1) {
		if m[0] == "do()" {
			enabled = true
		} else if m[0] == "don't()" {
			enabled = false
		} else if enabled {
			sum += lib.Atoi(m[1]) * lib.Atoi(m[2])
		}
	}
	return sum
}

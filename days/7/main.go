package main

import (
	_ "embed"
	"fmt"
	"iter"
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

type InputEntry struct {
	Sum     int64
	Numbers []int64
}

func input() []InputEntry {
	r := regexp.MustCompile(`(\d+): ([\d ]+)`)
	return slices.Collect(lib.Map(strings.Lines(rawInput), func(l string) InputEntry {
		m := lib.Match(r, l)
		return InputEntry{
			Sum:     lib.Atoi64(m[1]),
			Numbers: lib.MapSlice(strings.Split(m[2], " "), lib.Atoi64),
		}
	}))
}

func part1(inputs []InputEntry) int64 {
	total := int64(0)
	for _, e := range inputs {
		for s := range possibleSums(e.Numbers, false) {
			if s == e.Sum {
				// total += lib.Reduce(e.Numbers, func(a, b int64) int64 {
				// 	return a + b
				// }, 0)
				total += e.Sum
				break
			}
		}
	}
	return total
}

func part2(inputs []InputEntry) int64 {
	total := int64(0)
	for _, e := range inputs {
		for s := range possibleSums(e.Numbers, true) {
			if s == e.Sum {
				// total += lib.Reduce(e.Numbers, func(a, b int64) int64 {
				// 	return a + b
				// }, 0)
				total += e.Sum
				break
			}
		}
	}
	return total
}

func possibleSums(nums []int64, allowConcat bool) iter.Seq[int64] {
	return _possibleSums(nums[0], nums[1:], allowConcat)
}

func _possibleSums(left int64, right []int64, allowConcat bool) iter.Seq[int64] {
	return func(yield func(int64) bool) {
		if len(right) == 0 {
			yield(left)
			return
		}

		next := right[0]
		rem := right[1:]

		// add
		for s := range _possibleSums(left+next, rem, allowConcat) {
			if !yield(s) {
				return
			}
		}

		// multiply
		for s := range _possibleSums(left*next, rem, allowConcat) {
			if !yield(s) {
				return
			}
		}

		// concat
		if allowConcat {
			// TODO: could use log or whatever but screw it
			concat := lib.Atoi64(fmt.Sprintf("%d%d", left, next))
			for s := range _possibleSums(concat, rem, allowConcat) {
				if !yield(s) {
					return
				}
			}
		}
	}
}

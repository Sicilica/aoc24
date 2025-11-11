package main

import (
	_ "embed"
	"iter"
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

func input() ([][2]int, [][]int) {
	var rules [][2]int
	var updates [][]int
	for l := range strings.Lines(rawInput) {
		if strings.Contains(l, "|") {
			parts := lib.MapSlice(strings.SplitN(l, "|", 2), lib.Atoi)
			rules = append(rules, [2]int{parts[0], parts[1]})
		} else if strings.Contains(l, ",") {
			update := lib.MapSlice(strings.Split(l, ","), lib.Atoi)
			lib.Assert(len(update)%2 == 1)
			updates = append(updates, update)
		}
	}
	return rules, updates
}

func part1(rules [][2]int, updates [][]int) int {
	invRules := make(map[int]map[int]struct{})
	for _, r := range rules {
		if _, ok := invRules[r[1]]; !ok {
			invRules[r[1]] = make(map[int]struct{})
		}
		invRules[r[1]][r[0]] = struct{}{}
	}

	sum := 0
	for _, u := range updates {
		if valid(u, invRules) {
			sum += u[(len(u)-1)/2]
		}
	}
	return sum
}

func part2(rules [][2]int, updates [][]int) int {
	invRules := make(map[int]map[int]struct{})
	for _, r := range rules {
		if _, ok := invRules[r[1]]; !ok {
			invRules[r[1]] = make(map[int]struct{})
		}
		invRules[r[1]][r[0]] = struct{}{}
	}

	sum := 0
	for _, u := range updates {
		if valid(u, invRules) {
			continue
		}

		fixedU := fix(u, invRules)

		sum += fixedU[(len(u)-1)/2]
	}
	return sum
}

func valid(update []int, invRules map[int]map[int]struct{}) bool {
	for i := range update {
		if !validAt(update, i, invRules) {
			return false
		}
	}
	return true
}

func validAt(update []int, index int, invRules map[int]map[int]struct{}) bool {
	for j := range index {
		if _, ok := invRules[update[j]][update[index]]; ok {
			return false
		}
	}
	for j := index + 1; j < len(update); j++ {
		if _, ok := invRules[update[index]][update[j]]; ok {
			return false
		}
	}
	return true
}

func fix(update []int, invRules map[int]map[int]struct{}) []int {
	for res := range addOrdered(nil, update, invRules) {
		return res
	}
	panic("no valid order found")
}

func addOrdered(existing, addl []int, invRules map[int]map[int]struct{}) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		// Base case #1: nothing to add
		// (really this is to card against bad inputs)
		if len(addl) == 0 {
			yield(existing)
			return
		}

		// Recursive case
		if len(addl) > 1 {
			others := make([]int, len(addl)-1)
			copy(others, addl[1:])
			// For each element:
			for i := range addl {
				// Get the set of solutions where we added every other element already
				for sub := range addOrdered(existing, others, invRules) {
					// Find a solution where we now add this element
					for res := range addOrdered(sub, addl[i:i+1], invRules) {
						if !yield(res) {
							return
						}
					}
				}
				others[i] = addl[i]
			}
		}

		// Base case #2: adding one element to an empty array
		if len(existing) == 0 {
			yield(addl)
			return
		}

		// Try inserting before each element (including after last element)
		toAdd := addl[0]
		out := make([]int, len(existing)+1)
		copy(out[1:], existing)
		for i := 0; i <= len(existing); i++ {
			if i > 0 {
				out[i-1] = out[i]
			}
			out[i] = toAdd

			if validAt(out, i, invRules) {
				if !yield(out) {
					return
				}
			}
		}
	}
}

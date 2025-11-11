package main

import (
	_ "embed"
	"fmt"
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

func input() []uint64 {
	return lib.MapSlice(strings.Split(strings.TrimSpace(rawInput), " "), func(s string) uint64 {
		i := lib.Atoi64(s)
		return uint64(i)
	})
}

func part1(stones []uint64) int {
	return lib.Sum(lib.Map(slices.Values(stones), func(stone uint64) int {
		return stoneCount(stone, 25)
	}))
}

func part2(stones []uint64) int {
	return lib.Sum(lib.Map(slices.Values(stones), func(stone uint64) int {
		return stoneCount(stone, 75)
	}))
}

type MemoKey struct {
	Stone      uint64
	Iterations int
}

var memo = map[MemoKey]int{}

func stoneCount(stone uint64, iterations int) int {
	if iterations == 0 {
		return 1
	}
	key := MemoKey{Stone: stone, Iterations: iterations}
	if res, ok := memo[key]; ok {
		return res
	}

	res := func() int {
		if stone == 0 {
			return stoneCount(1, iterations-1)
		}

		str := fmt.Sprint(stone)
		if len(str)%2 == 0 {
			half := len(str) / 2
			return stoneCount(uint64(lib.Atoi64(str[:half])), iterations-1) + stoneCount(uint64(lib.Atoi64(str[half:])), iterations-1)
		}

		return stoneCount(stone*2024, iterations-1)
	}()
	memo[key] = res
	return res
}

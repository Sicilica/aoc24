package main

import (
	_ "embed"
	"iter"
	"math"
	"regexp"
	"slices"
	"strings"

	"github.com/Sicilica/aoc24/lib"
	"github.com/Sicilica/aoc24/lib2"
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

type Machine struct {
	A     lib2.Vec2i
	B     lib2.Vec2i
	Prize lib2.Vec2i
}

func input() []Machine {
	buttonReg := regexp.MustCompile(`Button (A|B): X\+(\d+), Y\+(\d+)`)
	prizeReg := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	var out []Machine
	machine := Machine{}
	for l := range strings.Lines(rawInput) {
		if strings.HasPrefix(l, "Button") {
			m := lib.Match(buttonReg, l)
			button := lib2.Vec2i{
				lib.Atoi(m[2]),
				lib.Atoi(m[3]),
			}
			if m[1] == "A" {
				machine.A = button
			} else {
				machine.B = button
			}
		} else if strings.HasPrefix(l, "Prize") {
			m := lib.Match(prizeReg, l)
			machine.Prize = lib2.Vec2i{
				lib.Atoi(m[1]),
				lib.Atoi(m[2]),
			}
			out = append(out, machine)
		}
	}
	return out
}

func part1(machines []Machine) int {
	return lib.ReduceSeq(lib.MapSeq(slices.Values(machines), func(m Machine) int {
		return m.Cost(3, 1, 100)
	}), func(a, b int) int {
		return a + b
	}, 0)
}

func part2(machines []Machine) int64 {
	return lib.ReduceSeq(lib.MapSeq(slices.Values(machines), func(m Machine) int64 {
		return m.CostBig(3, 1, 10000000000000)
	}), func(a, b int64) int64 {
		return a + b
	}, 0)
}

func (m Machine) Cost(aCost, bCost, maxPresses int) int {
	return lib.ReduceSeq(lib.MapSeq2(m.Solutions(maxPresses), func(a, b int) int {
		return a*aCost + b*bCost
	}), func(a, b int) int {
		if a == 0 || a > b {
			return b
		}
		return a
	}, 0)
}

func (m Machine) CostBig(aCost, bCost, offset int64) int64 {
	return lib.ReduceSeq(lib.MapSeq2(m.SolutionsBig(offset), func(a, b int64) int64 {
		return a*aCost + b*bCost
	}), func(a, b int64) int64 {
		if a == 0 || a > b {
			return b
		}
		return a
	}, 0)
}

func (m Machine) Solutions(maxPresses int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for a := range maxPresses {
			rem := m.Prize.Minus(m.A.Times(a))
			if rem.X() < 0 || rem.Y() < 0 {
				break
			}

			if rem.X()%m.B.X() != 0 || rem.Y()%m.B.Y() != 0 {
				continue
			}
			b := rem.X() / m.B.X()
			if b > maxPresses || b != rem.Y()/m.B.Y() {
				continue
			}

			if !yield(a, b) {
				return
			}
		}
	}
}

func (m Machine) SolutionsBig(offset int64) iter.Seq2[int64, int64] {
	// Avoid some divide-by-zero cases
	lib.Assert(m.A.X() != 0)
	lib.Assert(m.A.Y() != 0)

	buttonA := [2]int64{int64(m.A.X()), int64(m.A.Y())}
	buttonB := [2]int64{int64(m.B.X()), int64(m.B.Y())}
	prize := [2]int64{offset + int64(m.Prize.X()), offset + int64(m.Prize.Y())}

	// If the lines are parallel, then this will be more complicated
	lib.Assert(buttonA[0]*buttonB[1] != buttonA[1]*buttonB[0])

	// There should be exactly one intersection point
	return func(yield func(int64, int64) bool) {
		// TL;DR - We do some algebra to get two lines, both representing the number of A presses required
		// - One line shows how many A presses we need to get X correct, based on the number of B presses
		// - The other shows how many are needed to get Y correct, again based on B presses
		// Any point on either line represents a combination of A and B presses that get us to the correct coordinate on the respective axis.
		// Since we need both axes to be correct, and the lines aren't parallel, they intersect at exactly one point.
		// BUT, that point could need a non-integer number of presses, negative presses, etc.
		//
		// We could also handle the case where the lines are parallel -- which either means they never cross, or are overlapping --
		// but, it turns out the input data doesn't include any such cases.
		numerator := float64(prize[0])/float64(buttonA[0]) - float64(prize[1])/float64(buttonA[1])
		denominator := float64(buttonB[0])/float64(buttonA[0]) - float64(buttonB[1])/float64(buttonA[1])
		b := int64(math.Round(numerator / denominator))
		a := (prize[0] - (buttonB[0] * b)) / buttonA[0]

		// Either of a/b could be negative, or we could be off because the real answer wasn't an integer
		if a < 0 || b < 0 || (a*buttonA[0]+b*buttonB[0] != prize[0]) || (a*buttonA[1]+b*buttonB[1] != prize[1]) {
			return
		}

		yield(a, b)
	}
}

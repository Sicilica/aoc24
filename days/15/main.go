package main

import (
	_ "embed"
	"fmt"
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

type object uint8

const (
	obj_empty object = iota
	obj_wall
	obj_box
	obj_boxL
	obj_boxR
)

func input() (lib.SparseGrid2i[object], lib.Vec2i, []lib.Vec2i) {
	grid := make(lib.SparseGrid2i[object])
	robot := lib.Vec2i{-1, -1}
	var instructions []lib.Vec2i

	parsingGrid := true
	y := 0
	for l := range strings.Lines(rawInput) {
		l = strings.TrimSpace(l)

		if parsingGrid {
			if l == "" {
				parsingGrid = false
				continue
			}

			for x, c := range l {
				switch c {
				case '#':
					grid.Set(lib.Vec2i{x, y}, obj_wall)
				case '.':
					// noop
				case 'O':
					grid.Set(lib.Vec2i{x, y}, obj_box)
				case '@':
					lib.Assert(robot[0] == -1)
					robot = lib.Vec2i{x, y}
				default:
					panic("unrecognized char in grid")
				}
			}

			y++
		} else {
			for _, c := range l {
				switch c {
				case '>':
					instructions = append(instructions, lib.Vec2i{1, 0})
				case '<':
					instructions = append(instructions, lib.Vec2i{-1, 0})
				case '^':
					instructions = append(instructions, lib.Vec2i{0, -1})
				case 'v':
					instructions = append(instructions, lib.Vec2i{0, 1})
				default:
					panic("unrecognized char in instructions")
				}
			}
		}
	}
	lib.Assert(robot[0] != -1)
	lib.Assert(!parsingGrid)

	return grid, robot, instructions
}

func part1(grid lib.SparseGrid2i[object], robot lib.Vec2i, instructions []lib.Vec2i) int {
	grid = grid.Copy()

	for _, dir := range instructions {
		firstTile := robot.Plus(dir)

		firstObj := lib.IgnoreOK(grid.Get(robot.Plus(dir)))
		if firstObj == obj_wall {
			continue
		}

		lastTile := firstTile
		for lib.IgnoreOK(grid.Get(lastTile)) == obj_box {
			lastTile = lastTile.Plus(dir)
		}

		lastObj := lib.IgnoreOK(grid.Get(lastTile))
		if lastObj == obj_wall {
			continue
		}
		lib.Assert(lastObj == obj_empty)
		grid.Set(firstTile, obj_empty)
		grid.Set(lastTile, firstObj)
		robot = firstTile
	}

	return lib.Sum(lib.Map(lib.Indices(grid.All(), obj_box), func(pos lib.Vec2i) int {
		return pos.X() + pos.Y()*100
	}))
}

func part2(grid lib.SparseGrid2i[object], robot lib.Vec2i, instructions []lib.Vec2i) int {
	grid = makeWideGrid(grid)
	robot[0] *= 2

	queue := lib.NewStack[lib.Vec2i]()
	var push []lib.Vec2i
	for _, dir := range instructions {
		firstTile := robot.Plus(dir)

		push := push[:0]
		queue.Clear()
		queue.Push(firstTile)
		blocked := false
		for queue.Len() > 0 {
			pos := queue.Pop()

			switch lib.IgnoreOK(grid.Get(pos)) {
			case obj_empty:
				// noop
			case obj_wall:
				blocked = true
				queue.Clear()
			case obj_boxR:
				pos = pos.Plus(lib.Vec2i{-1, 0})
				fallthrough
			case obj_boxL:
				push = append(push, pos)

				// dir[0]=-1: push from left
				// dir[0]=1: push from right
				// dir[0]=0: push both
				if dir[0] != 1 {
					queue.Push(pos.Plus(dir))
				}
				if dir[0] != -1 {
					queue.Push(pos.Plus(dir).Plus(lib.Vec2[int]{1, 0}))
				}
			default:
				lib.Assert(false)
			}
		}

		if blocked {
			continue
		}

		robot = firstTile

		for _, pos := range push {
			delete(grid, pos)
			delete(grid, pos.Plus(lib.Vec2i{1, 0}))
		}

		for _, pos := range push {
			pos = pos.Plus(dir)
			grid.Set(pos, obj_boxL)
			grid.Set(pos.Plus(lib.Vec2i{1, 0}), obj_boxR)
		}
	}

	if false {
		debugPrint(grid, robot)
	}

	return lib.Sum(lib.Map(lib.Indices(grid.All(), obj_boxL), func(pos lib.Vec2i) int {
		return pos.X() + pos.Y()*100
	}))
}

func makeWideGrid(grid lib.SparseGrid2i[object]) lib.SparseGrid2i[object] {
	out := make(lib.SparseGrid2i[object])
	for pos, obj := range grid {
		pos[0] *= 2
		if obj == obj_wall {
			out.Set(pos, obj)
			out.Set(pos.Plus(lib.Vec2i{1, 0}), obj)
		} else {
			lib.Assert(obj == obj_box)
			out.Set(pos, obj_boxL)
			out.Set(pos.Plus(lib.Vec2i{1, 0}), obj_boxR)
		}
	}
	return out
}

func debugPrint(grid lib.Grid2i[object], robot lib.Vec2i) {
	for y := range 10 {
		for x := range 20 {
			if robot.Equals(lib.Vec2i{x, y}) {
				fmt.Print("@")
				continue
			}

			switch lib.IgnoreOK(grid.Get(lib.Vec2i{x, y})) {
			case obj_empty:
				fmt.Print(".")
			case obj_wall:
				fmt.Print("#")
			case obj_box:
				fmt.Print("O")
			case obj_boxL:
				fmt.Print("[")
			case obj_boxR:
				fmt.Print("]")
			default:
				panic("unknown object")
			}
		}
		fmt.Println()
	}
}

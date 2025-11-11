package main

import (
	_ "embed"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
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

type Robot struct {
	Pos lib.Vec2i
	Vel lib.Vec2i
}

func input() []Robot {
	r := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	return slices.Collect(lib.Map(strings.Lines(rawInput), func(l string) Robot {
		m := lib.Match(r, l)
		return Robot{
			Pos: lib.Vec2i{
				lib.Atoi(m[1]),
				lib.Atoi(m[2]),
			},
			Vel: lib.Vec2i{
				lib.Atoi(m[3]),
				lib.Atoi(m[4]),
			},
		}
	}))
}

func part1(robots []Robot) int {
	quadrants := []int{0, 0, 0, 0}

	bounds := lib.Vec2i{101, 103}
	lib.Assert(bounds.X()%2 == 1)
	lib.Assert(bounds.Y()%2 == 1)
	center := lib.Vec2i{bounds.X() / 2, bounds.Y() / 2}

	duration := 100

	for _, r := range robots {
		p := r.Pos.Plus(lib.Vec2i{
			multiplyMod(r.Vel.X(), duration, bounds.X()),
			multiplyMod(r.Vel.Y(), duration, bounds.Y()),
		})
		p = lib.Vec2i{
			safemod(p.X(), bounds.X()),
			safemod(p.Y(), bounds.Y()),
		}

		if p.X() == center.X() || p.Y() == center.Y() {
			// Robot not in a quadrant; skip
			continue
		}

		// Count this robot in a quadrant
		q := 0
		if p.X() > center.X() {
			q += 1
		}
		if p.Y() > center.Y() {
			q += 2
		}
		quadrants[q]++
	}

	return lib.Reduce(slices.Values(quadrants), 1, func(a, b int) int {
		return a * b
	})
}

func part2(robots []Robot) string {
	// TODO: we're very destructive of robots throughout this, but oh well
	// for i, r := range robots {
	// 	robots[i] = Robot{
	// 		Pos: lib.Vec2i{r.Pos.Y(), r.Pos.X()},
	// 		Vel: lib.Vec2i{r.Vel.Y(), r.Vel.X()},
	// 	}
	// }

	bounds := lib.Vec2i{101, 103}
	// bounds = lib.Vec2i{bounds.Y(), bounds.X()}
	grid := lib.MakeFixedGrid2[int](bounds.X(), bounds.Y())
	for _, r := range robots {
		grid.Set(r.Pos, lib.IgnoreOK(grid.Get(r.Pos))+1)
	}

	duration := 10000
	dir := "days/14/output"
	lib.NoErr(os.MkdirAll(dir, 0755))

	img := image.NewGray(image.Rect(0, 0, bounds.X(), bounds.Y()))

	for i := range duration {
		if i > 0 {
			for i, r := range robots {
				grid.Set(r.Pos, lib.IgnoreOK(grid.Get(r.Pos))-1)
				robots[i].Pos = lib.Vec2i{
					safemod(r.Pos.X()+r.Vel.X(), bounds.X()),
					safemod(r.Pos.Y()+r.Vel.Y(), bounds.Y()),
				}
				grid.Set(robots[i].Pos, lib.IgnoreOK(grid.Get(robots[i].Pos))+1)
			}
		}

		for x := range bounds.X() {
			for y := range bounds.Y() {
				brightness := uint8(0)
				if lib.IgnoreOK(grid.Get(lib.Vec2i{x, y})) > 0 {
					brightness = 255
				}
				img.Set(x, y, color.Gray{Y: brightness})
			}
		}
		writeImage(img, fmt.Sprintf("%s/%d.png", dir, i))
	}
	return "output images to directory"
}

func multiplyMod(x, y, mod int) int {
	// TODO: this is only necessary for math/big
	if (x < 0) != (y < 0) {
		return -multiplyMod(-x, y, mod)
	}

	// TODO: at y=100, this doesn't actually overflow (obviously)
	// return int((int64(x) * int64(y)) % int64(mod))

	res := new(big.Int).Mul(big.NewInt(int64(x)), big.NewInt(int64(y)))
	return int(res.Mod(res, big.NewInt(int64(mod))).Int64())
}

func safemod(x, mod int) int {
	for x < 0 {
		x += mod
	}
	return x % mod
}

func writeImage(img image.Image, filename string) {
	f := lib.Must(os.Create(filename))
	defer f.Close()
	lib.NoErr(png.Encode(f, img))
}

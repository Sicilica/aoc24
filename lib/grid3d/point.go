package grid3d

type Point [3]int

func (p Point) X() int {
	return p[0]
}

func (p Point) Y() int {
	return p[1]
}

func (p Point) Z() int {
	return p[2]
}

// To returns the vector that would need to be added to `p` to get `o`.
func (p Point) To(o Point) Point {
	return o.Minus(p)
}

func (p Point) Plus(o Point) Point {
	return Point{p[0] + o[0], p[1] + o[1], p[2] + o[2]}
}

func (p Point) Minus(o Point) Point {
	return Point{p[0] - o[0], p[1] - o[1], p[2] - o[2]}
}

func (p Point) Times(s int) Point {
	return Point{p[0] * s, p[1] * s, p[2] * s}
}

func (p Point) Equals(o Point) bool {
	return p[0] == o[0] && p[1] == o[1]
}

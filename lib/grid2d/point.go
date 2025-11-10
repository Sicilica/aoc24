package grid2d

type Point [2]int

func (p Point) X() int {
	return p[0]
}

func (p Point) Y() int {
	return p[1]
}

// To returns the vector that would need to be added to `p` to get `o`.
func (p Point) To(o Point) Point {
	return o.Minus(p)
}

func (p Point) Plus(o Point) Point {
	return Point{p[0] + o[0], p[1] + o[1]}
}

func (p Point) Minus(o Point) Point {
	return Point{p[0] - o[0], p[1] - o[1]}
}

func (p Point) Times(s int) Point {
	return Point{p[0] * s, p[1] * s}
}

func (p Point) Equals(o Point) bool {
	return p[0] == o[0] && p[1] == o[1]
}

func (p Point) OnEdge(r Rect) bool {
	// return true if point p is on the edge of rectangle r, otherwise false
	return ((p[0] == r[0][0] || p[0] == r[1][0]) && (p[1] >= r[0][1] && p[1] <= r[1][1])) || ((p[1] == r[0][1] || p[1] == r[1][1]) && (p[1] >= r[0][1] && p[1] <= r[1][1]))
}

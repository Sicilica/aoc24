package grid2d

type Rect [2]Point

func (r Rect) Contains(p Point) bool {
	return p[0] >= r[0][0] && p[0] < r[1][0] && p[1] >= r[0][1] && p[1] < r[1][1]
}

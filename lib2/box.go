package lib2

type boxVec[Self any] interface {
	comparable
	Minus(Self) Self
	_contains(box [2]Self, point Self) bool
	_overlaps(a, b [2]Self) bool
}

type Box[Vec boxVec[Vec]] [2]Vec

func (b Box[Vec]) Contains(v Vec) bool {
	return b[0]._contains(b, v)
}

func (_ Vec2[T]) _contains(b [2]Vec2[T], point Vec2[T]) bool {
	return point[0] >= b[0][0] && point[0] < b[1][0] && point[1] >= b[0][1] && point[1] < b[1][1]
}

func (_ Vec3[T]) _contains(b [2]Vec3[T], point Vec3[T]) bool {
	return point[0] >= b[0][0] && point[0] < b[1][0] && point[1] >= b[0][1] && point[1] < b[1][1] && point[2] >= b[0][2] && point[2] < b[1][2]
}

func (b Box[Vec]) Overlaps(other Box[Vec]) bool {
	return b[0]._overlaps(b, other)
}

func (v Vec2[T]) _overlaps(a, b [2]Vec2[T]) bool {
	return !(a[0][0] >= b[1][0] || a[1][0] <= b[0][0] || a[0][1] >= b[1][1] || a[1][1] <= b[0][1])
}

func (v Vec3[T]) _overlaps(a, b [2]Vec3[T]) bool {
	return !(a[0][0] >= b[1][0] || a[1][0] <= b[0][0] || a[0][1] >= b[1][1] || a[1][1] <= b[0][1] || a[0][2] >= b[1][2] || a[1][2] <= b[0][2])
}

func (b Box[Vec]) Size() Vec {
	return b[1].Minus(b[0])
}

type Box2[T Real] = Box[Vec2[T]]

type Box3[T Real] = Box[Vec3[T]]

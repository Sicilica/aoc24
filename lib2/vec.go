package lib2

type Vec2[T Real] [2]T

type Vec2i = Vec2[int]

func (v Vec2[T]) X() T {
	return v[0]
}

func (v Vec2[T]) Y() T {
	return v[1]
}

func (v Vec2[T]) Plus(other Vec2[T]) Vec2[T] {
	return Vec2[T]{
		v[0] + other[0],
		v[1] + other[1],
	}
}

func (v Vec2[T]) Minus(other Vec2[T]) Vec2[T] {
	return Vec2[T]{
		v[0] - other[0],
		v[1] - other[1],
	}
}

func (v Vec2[T]) Times(scalar T) Vec2[T] {
	return Vec2[T]{
		v[0] * scalar,
		v[1] * scalar,
	}
}

func (v Vec2[T]) Equals(other Vec2[T]) bool {
	return v[0] == other[0] && v[1] == other[1]
}

type Vec3[T Real] [3]T

type Vec3i = Vec3[int]

func (v Vec3[T]) X() T {
	return v[0]
}

func (v Vec3[T]) Y() T {
	return v[1]
}

func (v Vec3[T]) Z() T {
	return v[2]
}

func (v Vec3[T]) Plus(other Vec3[T]) Vec3[T] {
	return Vec3[T]{
		v[0] + other[0],
		v[1] + other[1],
		v[2] + other[2],
	}
}

func (v Vec3[T]) Minus(other Vec3[T]) Vec3[T] {
	return Vec3[T]{
		v[0] - other[0],
		v[1] - other[1],
		v[2] - other[2],
	}
}

func (v Vec3[T]) Times(scalar T) Vec3[T] {
	return Vec3[T]{
		v[0] * scalar,
		v[1] * scalar,
		v[2] * scalar,
	}
}

func (v Vec3[T]) Equals(other Vec3[T]) bool {
	return v[0] == other[0] && v[1] == other[1] && v[2] == other[2]
}

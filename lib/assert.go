package lib

func Assert(cond bool) {
	if !cond {
		panic("assertion failed")
	}
}

func Must[T any](v T, err error) T {
	NoErr(err)
	return v
}

func NoErr(err error) {
	Assert(err == nil)
}

func NotEmpty[T any](v []T) []T {
	Assert(len(v) > 0)
	return v
}

func NotNil[T any](v *T) *T {
	Assert(v != nil)
	return v
}

func OK[T any](v T, ok bool) T {
	Assert(ok)
	return v
}

func IgnoreOK[T any](v T, ok bool) T {
	return v
}

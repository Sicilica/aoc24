package lib

import "log"

func Assert(cond bool) {
	if !cond {
		panic("assertion failed")
	}
}

func IgnoreOK[T any](v T, ok bool) T {
	return v
}

func Must[T any](v T, err error) T {
	NoErr(err)
	return v
}

func NoErr(err error) {
	if err != nil {
		log.Println(err)
		Assert(err == nil)
	}
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

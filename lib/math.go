package lib

import (
	"strconv"
	"strings"
)

func Abs[T int](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Atoi(s string) int {
	return Must(strconv.Atoi(strings.TrimSpace(s)))
}

func Atoi64(s string) int64 {
	return Must(strconv.ParseInt(strings.TrimSpace(s), 10, 64))
}

package lib

type Stack[T any] []T

func Push[T any](s Stack[T], v T) Stack[T] {
	return append(s, v)
}

func Pop[T any](s Stack[T]) (T, Stack[T]) {
	v := s[len(s)-1]
	return v, s[:len(s)-1]
}

func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

func (s *Stack[T]) Push(v T) {
	*s = Push(*s, v)
}

func (s *Stack[T]) Pop() T {
	var v T
	v, *s = Pop(*s)
	return v
}

func (s *Stack[T]) Clear() {
	*s = (*s)[:0]
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

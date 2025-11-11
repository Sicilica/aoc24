package lib2

import (
	"iter"
)

// Count returns the number of elements in the Seq[T].
func Count[T any](it iter.Seq[T]) int {
	count := 0
	for range it {
		count++
	}
	return count
}

// Count2 returns the number of elements in the Seq2[K,V].
func Count2[K, V any](it iter.Seq2[K, V]) int {
	count := 0
	for range it {
		count++
	}
	return count
}

// CountFunc returns the number of elements in the Seq[T] that satisfy the predicate.
// This is equivalent to `Count(Filter(it, fn))`.
func CountFunc[T any](it iter.Seq[T], fn func(T) bool) int {
	count := 0
	for v := range it {
		if fn(v) {
			count++
		}
	}
	return count
}

// CountFunc2 returns the number of elements in the Seq2[K,V] that satisfy the predicate.
// This is equivalent to `Count(Filter2(it, fn))`.
func CountFunc2[K, V any](it iter.Seq2[K, V], fn func(K, V) bool) int {
	count := 0
	for k, v := range it {
		if fn(k, v) {
			count++
		}
	}
	return count
}

// Every returns true if every element in the Seq[T] satisfies the predicate.
func Every[T any](it iter.Seq[T], fn func(T) bool) bool {
	for v := range it {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Every2 returns true if every element in the Seq2[K,V] satisfies the predicate.
func Every2[K, V any](it iter.Seq2[K, V], fn func(K, V) bool) bool {
	for k, v := range it {
		if !fn(k, v) {
			return false
		}
	}
	return true
}

// Filter returns a new Seq[T] that contains only the elements that satisfy the predicate.
func Filter[T any](it iter.Seq[T], fn func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		it(func(t T) bool {
			if fn(t) {
				return yield(t)
			}
			return true
		})
	}
}

// Filter2 returns a new Seq2[K,V] that contains only the elements that satisfy the predicate.
func Filter2[K, V any](it iter.Seq2[K, V], fn func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		it(func(k K, v V) bool {
			if fn(k, v) {
				return yield(k, v)
			}
			return true
		})
	}
}

// Find returns the first element in the Seq[T] that satisfies the predicate.
func Find[T any](it iter.Seq[T], fn func(T) bool) (T, bool) {
	for t := range it {
		if fn(t) {
			return t, true
		}
	}
	var zero T
	return zero, false
}

// Find2 returns the first element in the Seq2[K,V] that satisfies the predicate.
// For simple searches by key/value, use Index or Lookup instead.
func Find2[K, V any](it iter.Seq2[K, V], fn func(K, V) bool) (K, V, bool) {
	for k, v := range it {
		if fn(k, v) {
			return k, v, true
		}
	}
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}

// Index returns the first key in the Seq2[K,V] that has the given value.
func Index[K any, V comparable](it iter.Seq2[K, V], value V) (K, bool) {
	for k, v := range it {
		if v == value {
			return k, true
		}
	}
	var zero K
	return zero, false
}

// Indices returns a Seq[K] containing all the keys in the Seq2[K,V] that have the given value.
// This is equivalent to `Keys(Filter2(it, func(k, v) bool { return v == value }))`.
func Indices[K any, V comparable](it iter.Seq2[K, V], value V) iter.Seq[K] {
	return func(yield func(K) bool) {
		it(func(k K, v V) bool {
			if v == value {
				return yield(k)
			}
			return true
		})
	}
}

// Keys returns a Seq[K] containing all the keys in the Seq2[K,V].
func Keys[K, V any](it iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		it(func(k K, _ V) bool {
			return yield(k)
		})
	}
}

// Lookup returns the value associated with the given key in the Seq2[K,V].
func Lookup[K comparable, V any](it iter.Seq2[K, V], key K) (V, bool) {
	for k, v := range it {
		if k == key {
			return v, true
		}
	}
	var zero V
	return zero, false
}

// Map returns a new Seq[U] containing the results of applying the function to each element in the Seq[T].
func Map[T any, U any](it iter.Seq[T], fn func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		it(func(t T) bool {
			return yield(fn(t))
		})
	}
}

// Map2 returns a new Seq[U] containing the results of applying the function to each element in the Seq2[K,V].
func Map2[K, V, U any](it iter.Seq2[K, V], fn func(K, V) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		it(func(k K, v V) bool {
			return yield(fn(k, v))
		})
	}
}

// Reduce applies a function to each element of Seq[T], accumulating the result.
func Reduce[T any, U any](it iter.Seq[T], initial U, fn func(U, T) U) U {
	acc := initial
	for t := range it {
		acc = fn(acc, t)
	}
	return acc
}

// Some returns true if any element in the Seq[T] satisfies the predicate.
func Some[T any](it iter.Seq[T], fn func(T) bool) bool {
	for t := range it {
		if fn(t) {
			return true
		}
	}
	return false
}

// Some2 returns true if any element in the Seq2[K,V] satisfies the predicate.
func Some2[K, V any](it iter.Seq2[K, V], fn func(K, V) bool) bool {
	for k, v := range it {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Sum returns the sum of all elements in the Seq[T].
func Sum[T Real](it iter.Seq[T]) T {
	var sum T
	for t := range it {
		sum += t
	}
	return sum
}

// Values returns a Seq[V] containing all the values in the Seq2[K,V].
func Values[K, V any](it iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		it(func(_ K, v V) bool {
			return yield(v)
		})
	}
}

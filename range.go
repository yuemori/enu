package enu

import (
	"sync"
)

type IComparable[T any] interface {
	Compare(T) int
	Value() T
	Next() IComparable[T]
}

func NewRange[T any](min, max IComparable[T]) *Enumerator[T] {
	return &Enumerator[T]{iter: NewRangeEnumerator(
		min,
		max,
	)}
}

func NewRangeEnumerator[T any](min, max IComparable[T]) IEnumerable[T] {
	return &RangeEnumerator[T]{
		min: min,
		max: max,
	}
}

type RangeEnumerator[T any] struct {
	min     IComparable[T]
	max     IComparable[T]
	current IComparable[T]
	mu      sync.Mutex
}

func empty[T any]() T {
	var empty T
	return empty
}

func (e *RangeEnumerator[T]) Next() (T, bool) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.min.Compare(e.max.Value()) == 1 {
		return empty[T](), false
	}
	if e.current == nil {
		val := e.min.Value()
		e.current = e.min.Next()
		return val, true
	}
	if e.current.Compare(e.max.Value()) == 1 {
		return empty[T](), false
	}
	val := e.current.Value()
	e.current = e.current.Next()
	return val, true
}

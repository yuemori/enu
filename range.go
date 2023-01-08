package enu

import (
	"sync"
)

type RangeValuer[T1, T2 any] interface {
	Compare(T1) int
	Value() T1
	Next(step T2) RangeValuer[T1, T2]
}

func NewRange[T1, T2 any](min, max RangeValuer[T1, T2], step T2) IEnumerable[T1] {
	return &RangeEnumerator[T1, T2]{
		min:  min,
		max:  max,
		step: step,
	}
}

type RangeEnumerator[T1, T2 any] struct {
	min     RangeValuer[T1, T2]
	max     RangeValuer[T1, T2]
	current RangeValuer[T1, T2]
	step    T2
	mu      sync.Mutex
}

func empty[T any]() T {
	var empty T
	return empty
}

func (e *RangeEnumerator[T1, T2]) Stop() {}
func (e *RangeEnumerator[T1, T2]) Reset() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.current = e.min
}

func (e *RangeEnumerator[T1, T2]) Next() (T1, bool) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.min.Compare(e.max.Value()) == 1 {
		return empty[T1](), false
	}
	if e.current == nil {
		val := e.min.Value()
		e.current = e.min.Next(e.step)
		return val, true
	}
	if e.current.Compare(e.max.Value()) == 1 {
		return empty[T1](), false
	}
	val := e.current.Value()
	e.current = e.current.Next(e.step)
	return val, true
}

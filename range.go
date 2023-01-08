package enumerator

import (
	"sync"

	"golang.org/x/exp/constraints"
)

func NewRange[T constraints.Integer | constraints.Float](min, max T) *EnumeratorNumeric[T] {
	return &EnumeratorNumeric[T]{iter: NewRangeEnumerator(min, max, 1)}
}

func NewRangeWithStep[T constraints.Integer | constraints.Float](min, max, step T) *EnumeratorNumeric[T] {
	return &EnumeratorNumeric[T]{iter: NewRangeEnumerator(min, max, step)}
}

func NewRangeEnumerator[T constraints.Integer | constraints.Float](min, max T, step T) IEnumerableNumeric[T] {
	return &RangeEnumerator[T]{
		min:     min,
		max:     max,
		current: min,
		step:    step,
	}
}

type RangeEnumerator[T constraints.Integer | constraints.Float] struct {
	min     T
	max     T
	current T
	step    T
	mu      sync.Mutex
}

func empty[T any]() T {
	var empty T
	return empty
}

func (e *RangeEnumerator[T]) Next() (T, bool) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.max < e.min {
		return empty[T](), false
	}
	if e.step < 0 {
		val := e.current
		if e.current < e.max {
			return empty[T](), false
		}
		e.current -= T(e.step)
		return val, true
	} else {
		if e.max < e.current {
			return empty[T](), false
		}
		val := e.current
		e.current += T(e.step)
		return val, true
	}
}

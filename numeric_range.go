package enu

import "golang.org/x/exp/constraints"

func NewNumericRange[T constraints.Integer | constraints.Float](min, max T) *Enumerator[T] {
	return New(
		NewRange[T, T](
			Numeric[T]{value: min},
			Numeric[T]{value: max},
			T(1),
		),
	)
}

func NewNumericRangeWithStep[T constraints.Integer | constraints.Float](min, max, step T) *Enumerator[T] {
	return New(
		NewRange[T, T](
			Numeric[T]{value: min},
			Numeric[T]{value: max},
			step,
		),
	)
}

type Numeric[T constraints.Integer | constraints.Float] struct {
	value T
}

func (self Numeric[T]) Next(step T) RangeValuer[T, T] {
	return Numeric[T]{value: self.value + T(step)}
}

func (self Numeric[T]) Value() T {
	return T(self.value)
}

func (self Numeric[T]) Compare(other T) int {
	if T(self.value) < other {
		return -1
	}
	if T(self.value) > other {
		return 1
	}
	return 0
}

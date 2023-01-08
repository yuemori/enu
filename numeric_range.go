package enu

import "golang.org/x/exp/constraints"

func NewNumericRange[T constraints.Integer | constraints.Float](min, max T) *Enumerator[T] {
	return NewRange[T](
		NumericComparable[T]{value: min, step: 1},
		NumericComparable[T]{value: max, step: 1},
	)
}

func NewNumericRangeWithStep[T constraints.Integer | constraints.Float](min, max, step T) *Enumerator[T] {
	return NewRange[T](
		NumericComparable[T]{value: min, step: step},
		NumericComparable[T]{value: max, step: step},
	)
}

type NumericComparable[T constraints.Integer | constraints.Float] struct {
	value T
	step  T
}

func (self NumericComparable[T]) Next() IComparable[T] {
	return NumericComparable[T]{value: self.value + T(self.step), step: self.step}
}

func (self NumericComparable[T]) Value() T {
	return T(self.value)
}

func (self NumericComparable[T]) Compare(other T) int {
	if T(self.value) < other {
		return -1
	}
	if T(self.value) > other {
		return 1
	}
	return 0
}

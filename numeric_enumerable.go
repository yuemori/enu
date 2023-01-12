package enu

import (
	"golang.org/x/exp/constraints"
)

func FromNumeric[T constraints.Float | constraints.Integer](collection []T) *NumericEnumerable[T] {
	return NewNumeric[T](NewSliceEnumerator(collection))
}

func ToNumeric[T constraints.Float | constraints.Integer](e IEnumerable[T]) *NumericEnumerable[T] {
	return NewNumeric[T](e.GetEnumerator())
}

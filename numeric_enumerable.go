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

func (e *NumericEnumerable[T]) ToMap() map[int]T {
	result := map[int]T{}
	e.Each(func(item T, index int) {
		result[index] = item
	})
	return result
}

func (e *NumericEnumerable[T]) Min() T {
	return Min[T](e)
}

func (e *NumericEnumerable[T]) Max() T {
	return Max[T](e)
}

func (e *NumericEnumerable[T]) Sum() T {
	return Sum[T](e)
}

func (e *NumericEnumerable[T]) Sort() *NumericEnumerable[T] {
	return NewNumeric[T](Sort[T](e))
}

func (e *NumericEnumerable[T]) Uniq(item T) *NumericEnumerable[T] {
	return NewNumeric[T](Uniq[T](e))
}

func (e *NumericEnumerable[T]) Contains(item T) bool {
	return Contains[T](e, item)
}

func (e *NumericEnumerable[T]) IndexOf(item T) int {
	return IndexOf[T](e, item)
}

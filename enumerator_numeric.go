package enu

import (
	"golang.org/x/exp/constraints"
)

func FromNumeric[T constraints.Float | constraints.Integer](collection []T) *EnumeratorNumeric[T] {
	return NewNumeric[T](NewSliceEnumerator(collection))
}

func ToNumeric[T constraints.Float | constraints.Integer](e *Enumerator[T]) *EnumeratorNumeric[T] {
	return &EnumeratorNumeric[T]{
		iter: e.iter,
	}
}

func (e *EnumeratorNumeric[T]) ToMap() map[int]T {
	result := map[int]T{}
	e.Each(func(item T, index int) {
		result[index] = item
	})
	return result
}

func (e *EnumeratorNumeric[T]) Min() T {
	return Min(e.iter)
}

func (e *EnumeratorNumeric[T]) Max() T {
	return Max(e.iter)
}

func (e *EnumeratorNumeric[T]) Sum() T {
	return Sum(e.iter)
}

func (e *EnumeratorNumeric[T]) Sort() *EnumeratorNumeric[T] {
	return &EnumeratorNumeric[T]{iter: Sort(e.iter)}
}

func (e *EnumeratorNumeric[T]) Uniq(item T) *EnumeratorNumeric[T] {
	return &EnumeratorNumeric[T]{iter: Uniq(e.iter)}
}

func (e *EnumeratorNumeric[T]) Contains(item T) bool {
	return Contains(e.iter, item)
}

func (e *EnumeratorNumeric[T]) IndexOf(item T) int {
	return IndexOf(e.iter, item)
}

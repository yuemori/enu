package enu

import (
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

func FromOrdered[T constraints.Ordered](collection []T) *EnumeratorOrdered[T] {
	return NewOrdered[T](NewSliceEnumerator(collection))
}

func ToOrdered[T constraints.Ordered](e *Enumerator[T]) *EnumeratorOrdered[T] {
	return &EnumeratorOrdered[T]{
		iter: e.iter,
	}
}

func (e *EnumeratorOrdered[T]) ToMap() map[int]T {
	result := map[int]T{}
	e.Each(func(item T, index int) {
		result[index] = item
	})
	return result
}

func (e *EnumeratorOrdered[T]) Min() T {
	return lo.Min(e.ToSlice())
}

func (e *EnumeratorOrdered[T]) Max() T {
	return lo.Max(e.ToSlice())
}

func (e *EnumeratorOrdered[T]) Sort() *EnumeratorOrdered[T] {
	return &EnumeratorOrdered[T]{iter: Sort(e.iter)}
}

func (e *EnumeratorOrdered[T]) Uniq() *EnumeratorOrdered[T] {
	return &EnumeratorOrdered[T]{iter: Uniq(e.iter)}
}

func (e *EnumeratorOrdered[T]) Contains(item T) bool {
	return Contains(e.iter, item)
}

func (e *EnumeratorOrdered[T]) IndexOf(item T) int {
	return IndexOf(e.iter, item)
}

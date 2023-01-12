package enu

import (
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

func FromOrdered[T constraints.Ordered](collection []T) *OrderedEnumerable[T] {
	return NewOrdered[T](NewSliceEnumerator(collection))
}

func ToOrdered[T constraints.Ordered](e *Enumerable[T]) *OrderedEnumerable[T] {
	return NewOrdered[T](e.GetEnumerator())
}

func (e *OrderedEnumerable[T]) ToMap() map[int]T {
	result := map[int]T{}
	e.Each(func(item T, index int) {
		result[index] = item
	})
	return result
}

func (e *OrderedEnumerable[T]) Min() T {
	return lo.Min(e.ToSlice())
}

func (e *OrderedEnumerable[T]) Max() T {
	return lo.Max(e.ToSlice())
}

func (e *OrderedEnumerable[T]) Sort() *OrderedEnumerable[T] {
	return NewOrdered[T](Sort[T](e))
}

func (e *OrderedEnumerable[T]) Uniq() *OrderedEnumerable[T] {
	return NewOrdered[T](Uniq[T](e))
}

func (e *OrderedEnumerable[T]) Contains(item T) bool {
	return Contains[T](e, item)
}

func (e *OrderedEnumerable[T]) IndexOf(item T) int {
	return IndexOf[T](e, item)
}

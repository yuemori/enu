package enu

import (
	"sort"

	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

func FromOrdered[T constraints.Ordered](collection []T) *EnumeratorOrdered[T] {
	return NewOrdered[T](newSliceEnumerator(collection))
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
	res := e.ToSlice()
	sort.SliceStable(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	e.swap(res)
	return e
}

func (e *EnumeratorOrdered[T]) Uniq() *EnumeratorOrdered[T] {
	e.swap(lo.Uniq(e.ToSlice()))
	return e
}

func (e *EnumeratorOrdered[T]) Contains(item T) bool {
	return lo.Contains(e.ToSlice(), item)
}

func (e *EnumeratorOrdered[T]) IndexOf(item T) int {
	return lo.IndexOf(e.ToSlice(), item)
}

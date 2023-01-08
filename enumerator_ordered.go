package enu

import (
	"sort"

	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

func FromOrdered[T constraints.Ordered](collection []T) *EnumeratorOrdered[T] {
	return NewOrdered[T](newSliceEnumerator(collection))
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

package enu

import (
	"sort"

	"github.com/samber/lo"
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
	return lo.Min(e.ToSlice())
}

func (e *EnumeratorNumeric[T]) Max() T {
	return lo.Max(e.ToSlice())
}

func (e *EnumeratorNumeric[T]) Sort() *EnumeratorNumeric[T] {
	res := e.ToSlice()
	sort.SliceStable(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	e.swap(res)
	return e
}

func (e *EnumeratorNumeric[T]) Sum() T {
	return lo.Sum(e.ToSlice())
}

func (e *EnumeratorNumeric[T]) Uniq(item T) *EnumeratorNumeric[T] {
	e.swap(lo.Uniq(e.ToSlice()))
	return e
}

func (e *EnumeratorNumeric[T]) Contains(item T) bool {
	return lo.Contains(e.ToSlice(), item)
}

func (e *EnumeratorNumeric[T]) IndexOf(item T) int {
	return lo.IndexOf(e.ToSlice(), item)
}

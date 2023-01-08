package enumerator

import (
	"github.com/samber/lo"
)

func (e *EnumeratorC[T]) ToMap() map[int]T {
	result := map[int]T{}
	_ = e.Each(func(item T, index int) {
		result[index] = item
	})
	return result
}

func (e *EnumeratorC[T]) Uniq() *EnumeratorC[T] {
	e.iter = newSliceEnumerator(lo.Uniq(e.ToSlice()))
	return e
}

func (e *EnumeratorC[T]) Contains(item T) bool {
	return lo.Contains(e.ToSlice(), item)
}

func Comparable[T comparable](e Enumerator[T]) *EnumeratorC[T] {
	return &EnumeratorC[T]{
		iter: e.iter,
		err:  e.err,
	}
}

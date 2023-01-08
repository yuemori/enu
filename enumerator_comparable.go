package enumerator

import (
	"github.com/samber/lo"
)

func FromComparable[T comparable](collection []T) *EnumeratorComparable[T] {
	return NewComparable[T](newSliceEnumerator(collection))
}

func (e *EnumeratorComparable[T]) ToMap() map[int]T {
	result := map[int]T{}
	_ = e.Each(func(item T, index int) {
		result[index] = item
	})
	return result
}

func (e *EnumeratorComparable[T]) Uniq() *EnumeratorComparable[T] {
	e.iter = newSliceEnumerator(lo.Uniq(e.ToSlice()))
	return e
}

func (e *EnumeratorComparable[T]) Contains(item T) bool {
	return lo.Contains(e.ToSlice(), item)
}

func ToComparable[T comparable](e Enumerator[T]) *EnumeratorComparable[T] {
	return &EnumeratorComparable[T]{
		iter: e.iter,
		err:  e.err,
	}
}

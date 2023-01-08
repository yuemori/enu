package enu

import "github.com/samber/lo"

func MapE[T, R any](e *Enumerator[T], iteratee func(item T, index int) R) *Enumerator[R] {
	return &Enumerator[R]{
		iter: newSliceEnumerator(lo.Map(e.ToSlice(), iteratee)),
	}
}

func MapC[T comparable, R any](e *EnumeratorComparable[T], iteratee func(item T, index int) R) *Enumerator[R] {
	return &Enumerator[R]{
		iter: newSliceEnumerator(lo.Map(e.ToSlice(), iteratee)),
	}
}

func Map[T, R any](collection []T, iteratee func(item T, index int) R) *Enumerator[R] {
	return &Enumerator[R]{
		iter: newSliceEnumerator(lo.Map(collection, iteratee)),
	}
}

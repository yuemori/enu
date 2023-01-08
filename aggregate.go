package enu

import "github.com/samber/lo"

func MapE[T, R any](e *Enumerator[T], iteratee func(item T, index int) R) *Enumerator[R] {
	result := lo.Map(e.ToSlice(), iteratee)
	return &Enumerator[R]{
		iter:      newSliceEnumerator(result),
		result:    result,
		isStopped: true,
	}
}

func MapC[T comparable, R any](e *EnumeratorComparable[T], iteratee func(item T, index int) R) *Enumerator[R] {
	result := lo.Map(e.ToSlice(), iteratee)
	return &Enumerator[R]{
		iter:      newSliceEnumerator(result),
		result:    result,
		isStopped: true,
	}
}

func Map[T, R any](collection []T, iteratee func(item T, index int) R) *Enumerator[R] {
	result := lo.Map(collection, iteratee)
	return &Enumerator[R]{
		iter:      newSliceEnumerator(result),
		result:    result,
		isStopped: true,
	}
}

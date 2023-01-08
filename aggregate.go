package enumerator

import "github.com/samber/lo"

func MapE[T, R any](e *Enumerator[T], iteratee func(item T, index int) R) *Enumerator[R] {
	return &Enumerator[R]{
		iter: newSliceEnumerator(lo.Map(e.ToSlice(), iteratee)),
		err:  e.err,
	}
}

func MapC[T comparable, R any](e *EnumeratorComparable[T], iteratee func(item T, index int) R) *Enumerator[R] {
	return &Enumerator[R]{
		iter: newSliceEnumerator(lo.Map(e.ToSlice(), iteratee)),
		err:  e.err,
	}
}

func Map[T, R any](collection []T, iteratee func(item T, index int) R) *Enumerator[R] {
	return &Enumerator[R]{
		iter: newSliceEnumerator(lo.Map(collection, iteratee)),
	}
}

package enu

import "github.com/samber/lo"

func Map[T, R any](collection []T, iteratee func(item T, index int) R) *Enumerator[R] {
	result := lo.Map(collection, iteratee)
	return &Enumerator[R]{
		iter:      newSliceEnumerator(result),
		result:    result,
		isStopped: true,
	}
}

func Reduce[T, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	return lo.Reduce(collection, accumulator, initial)
}

package enu

func From[T any](collection []T) *Enumerable[T] {
	return New[T](NewSliceEnumerator(collection))
}

func (e *Enumerable[T]) ToMap() map[int]T {
	return Reduce[T](e, func(agg map[int]T, item T, index int) map[int]T {
		agg[index] = item
		return agg
	}, map[int]T{})
}

func (e *Enumerable[T]) Aggregate(accumulator func(agg []T, item T, index int) []T) []T {
	return Reduce[T](e, accumulator, []T{})
}

type IEnumerable[T any] interface {
	GetEnumerator() IEnumerator[T]
}

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

type IEnumerable[T any] interface {
	GetEnumerator() Enumerator[T]
}

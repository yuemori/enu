package enu

func Map[T, R any](iter IEnumerable[T], iteratee func(item T, index int) R) *Enumerator[R] {
	result := []R{}
	each(iter, func(item T, index int) bool {
		result = append(result, iteratee(item, index))
		return true
	})

	return &Enumerator[R]{
		iter: NewSliceEnumerator(result),
	}
}

func Reduce[T, R any](iter IEnumerable[T], accumulator func(agg R, item T, index int) R, initial R) R {
	each(iter, func(item T, index int) bool {
		initial = accumulator(initial, item, index)
		return true
	})

	return initial
}

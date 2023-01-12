package enu

func Map[T, R any](e IEnumerable[T], iteratee func(item T, index int) R) *Enumerable[R] {
	result := []R{}
	each(e, func(item T, index int) bool {
		result = append(result, iteratee(item, index))
		return true
	})

	return New[R](NewSliceEnumerator(result))
}

func Reduce[T, R any](e IEnumerable[T], accumulator func(agg R, item T, index int) R, initial R) R {
	each(e, func(item T, index int) bool {
		initial = accumulator(initial, item, index)
		return true
	})

	return initial
}

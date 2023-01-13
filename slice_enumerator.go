package enu

type SliceEnumerator[T any] struct {
	collection []T
	index      int
}

func NewSliceEnumerator[T any](collection []T) *SliceEnumerator[T] {
	return &SliceEnumerator[T]{collection: collection}
}

func (e *SliceEnumerator[T]) Dispose() {
	e.index = 0
}

func (e *SliceEnumerator[T]) Next() (T, bool) {
	if len(e.collection) > e.index {
		item := e.collection[e.index]
		e.index += 1
		return item, true
	}
	var empty T
	return empty, false
}

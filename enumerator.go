package enu

func From[T any](collection []T) *Enumerator[T] {
	return New[T](newSliceEnumerator(collection))
}

func (e *Enumerator[T]) ToMap() map[int]T {
	result := map[int]T{}
	_ = e.Each(func(item T, index int) {
		result[index] = item
	})
	return result
}

type sliceEnumerator[T any] struct {
	collection []T
	index      int
}

func newSliceEnumerator[T any](collection []T) *sliceEnumerator[T] {
	return &sliceEnumerator[T]{collection: collection}
}

func (e *sliceEnumerator[T]) Stop() {}
func (e *sliceEnumerator[T]) Next() (T, bool) {
	if len(e.collection) > e.index {
		item := e.collection[e.index]
		e.index += 1
		return item, true
	}
	var empty T
	return empty, false
}

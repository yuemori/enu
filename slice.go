package enumerator

func FromSlice[T any](collection []T) *Enumerator[T] {
	return &Enumerator[T]{iter: newSliceEnumerator(collection)}
}

type sliceEnumerator[T any] struct {
	collection []T
	index      int
}

func newSliceEnumerator[T any](collection []T) *sliceEnumerator[T] {
	return &sliceEnumerator[T]{collection: collection}
}

func (e *sliceEnumerator[T]) Stop() {}
func (e *sliceEnumerator[T]) Next() (T, error) {
	if len(e.collection) > e.index {
		item := e.collection[e.index]
		e.index += 1
		return item, nil
	}
	var empty T
	return empty, Done
}

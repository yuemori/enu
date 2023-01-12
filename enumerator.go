package enu

func From[T any](collection []T) *Enumerator[T] {
	return New[T](NewSliceEnumerator(collection))
}

func (e *Enumerator[T]) ToMap() map[int]T {
	result := map[int]T{}
	e.Each(func(item T, index int) {
		result[index] = item
	})
	return result
}

type IEnumerable[T any] interface {
	Next() (T, bool)
	Reset()
	Stop()
}

type Query[T any] func(T, int) (T, bool)

func (fn Query[T]) Apply(item T, index int) (T, bool) {
	return fn(item, index)
}

type Queryer[T any] interface {
	Apply(item T, index int) (T, bool)
}

type SliceEnumerator[T any] struct {
	collection []T
	index      int
}

func NewSliceEnumerator[T any](collection []T) *SliceEnumerator[T] {
	return &SliceEnumerator[T]{collection: collection}
}

func (e *SliceEnumerator[T]) Reset() {
	e.index = 0
}

func (e *SliceEnumerator[T]) Stop() {
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

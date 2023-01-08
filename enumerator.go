package enu

import "sync"

func From[T any](collection []T) *Enumerator[T] {
	return New[T](newSliceEnumerator(collection))
}

func (e *Enumerator[T]) ToMap() map[int]T {
	result := map[int]T{}
	e.Each(func(item T, index int) {
		result[index] = item
	})
	return result
}

type sliceEnumerator[T any] struct {
	collection []T
	index      int
	mu         sync.Mutex
}

func newSliceEnumerator[T any](collection []T) *sliceEnumerator[T] {
	return &sliceEnumerator[T]{collection: collection}
}

func (e *sliceEnumerator[T]) Reset() {
	e.index = 0
}
func (e *sliceEnumerator[T]) Stop() {}
func (e *sliceEnumerator[T]) Next() (T, bool) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if len(e.collection) > e.index {
		item := e.collection[e.index]
		e.index += 1
		return item, true
	}
	var empty T
	return empty, false
}

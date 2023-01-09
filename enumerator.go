package enu

import (
	"sync"
)

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

type SliceEnumerator[T any] struct {
	collection []T
	index      int
	mu         sync.Mutex
}

func NewSliceEnumerator[T any](collection []T) *SliceEnumerator[T] {
	return &SliceEnumerator[T]{collection: collection}
}

func (e *SliceEnumerator[T]) Reset() {
	e.index = 0
}
func (e *SliceEnumerator[T]) Stop() {}
func (e *SliceEnumerator[T]) Next() (T, bool) {
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

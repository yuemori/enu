package enumerator

import "github.com/samber/lo"

func FromMap[K comparable, V any](collection map[K]V) *Enumerator[KeyValuePair[K, V]] {
	return &Enumerator[KeyValuePair[K, V]]{iter: newMapEnumerator(collection)}
}

type KeyValuePair[K comparable, V any] struct {
	Key   K
	Value V
}

type mapEnumerator[K comparable, V any] struct {
	collection map[K]V
	index      int
	keys       []K
}

func newMapEnumerator[K comparable, V any](collection map[K]V) *mapEnumerator[K, V] {
	return &mapEnumerator[K, V]{collection: collection, index: 0, keys: lo.Keys(collection)}
}

func (e *mapEnumerator[K, V]) Stop() {}
func (e *mapEnumerator[K, V]) Next() (KeyValuePair[K, V], error) {
	if len(e.keys) > e.index {
		key := e.keys[e.index]
		value := e.collection[key]
		e.index += 1
		return KeyValuePair[K, V]{Key: key, Value: value}, nil
	}
	var empty KeyValuePair[K, V]
	return empty, Done
}

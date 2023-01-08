package enumerator

import "github.com/samber/lo"

func FromMap[K comparable, V any](collection map[K]V) *EnumeratorMap[K, V] {
	return &EnumeratorMap[K, V]{iter: newMapEnumerator(collection)}
}

func (e *EnumeratorMap[K, V]) ToMap() map[K]V {
	result := map[K]V{}
	e.Each(func(kv KeyValuePair[K, V], _ int) {
		result[kv.Key] = kv.Value
	})
	return result
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
func (e *mapEnumerator[K, V]) Next() (KeyValuePair[K, V], bool) {
	if len(e.keys) > e.index {
		key := e.keys[e.index]
		value := e.collection[key]
		e.index += 1
		return KeyValuePair[K, V]{Key: key, Value: value}, true
	}
	var empty KeyValuePair[K, V]
	return empty, false
}

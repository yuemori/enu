package enu

import "github.com/samber/lo"

func FromMap[K comparable, V any](collection map[K]V) *EnumeratorMap[K, V] {
	return &EnumeratorMap[K, V]{iter: NewMapEnumerator(collection)}
}

func ToMap[K comparable, V any](e *Enumerator[KeyValuePair[K, V]]) *EnumeratorMap[K, V] {
	m := lo.Reduce(e.ToSlice(), func(agg map[K]V, kv KeyValuePair[K, V], index int) map[K]V {
		agg[kv.Key] = kv.Value
		return agg
	}, map[K]V{})
	return &EnumeratorMap[K, V]{iter: NewMapEnumerator(m)}
}

func (e *EnumeratorMap[K, V]) ToMap() map[K]V {
	result := map[K]V{}
	e.Each(func(kv KeyValuePair[K, V], _ int) {
		result[kv.Key] = kv.Value
	})
	return result
}

func (e *EnumeratorMap[K, V]) Keys() []K {
	return lo.Map(e.ToSlice(), func(kv KeyValuePair[K, V], _ int) K {
		return kv.Key
	})
}

func (e *EnumeratorMap[K, V]) Values() []V {
	return lo.Map(e.ToSlice(), func(kv KeyValuePair[K, V], _ int) V {
		return kv.Value
	})
}

type KeyValuePair[K comparable, V any] struct {
	Key   K
	Value V
}

type MapEnumerator[K comparable, V any] struct {
	collection map[K]V
	index      int
	keys       []K
}

func NewMapEnumerator[K comparable, V any](collection map[K]V) *MapEnumerator[K, V] {
	return &MapEnumerator[K, V]{collection: collection, index: 0, keys: lo.Keys(collection)}
}

func (e *MapEnumerator[K, V]) Reset() {
	e.index = 0
}

func (e *MapEnumerator[K, V]) Stop() {
	e.index = 0
}

func (e *MapEnumerator[K, V]) Next() (KeyValuePair[K, V], bool) {
	if len(e.keys) > e.index {
		key := e.keys[e.index]
		value := e.collection[key]
		e.index += 1
		return KeyValuePair[K, V]{Key: key, Value: value}, true
	}
	var empty KeyValuePair[K, V]
	return empty, false
}

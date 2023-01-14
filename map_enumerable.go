package enu

func FromMap[K comparable, V any](collection map[K]V) *MapEnumerable[K, V] {
	return NewMap[K, V](NewMapEnumerator(collection))
}

func ToMap[K comparable, V any](e IEnumerable[KeyValuePair[K, V]]) *MapEnumerable[K, V] {
	m := Reduce(e, func(agg map[K]V, kv KeyValuePair[K, V], index int) map[K]V {
		agg[kv.Key] = kv.Value
		return agg
	}, map[K]V{})
	return NewMap[K, V](NewMapEnumerator(m))
}

func (e *MapEnumerable[K, V]) ToMap() map[K]V {
	return Reduce[KeyValuePair[K, V]](e, func(agg map[K]V, kv KeyValuePair[K, V], index int) map[K]V {
		agg[kv.Key] = kv.Value
		return agg
	}, map[K]V{})
}

func (e *MapEnumerable[K, V]) Keys() []K {
	keys := []K{}
	e.Each(func(kv KeyValuePair[K, V], _ int) {
		keys = append(keys, kv.Key)
	})
	return keys
}

func (e *MapEnumerable[K, V]) Values() []V {
	values := []V{}
	e.Each(func(kv KeyValuePair[K, V], _ int) {
		values = append(values, kv.Value)
	})
	return values
}

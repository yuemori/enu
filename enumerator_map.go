package enumerator

func (e *EnumeratorMap[K, V]) ToMap() map[K]V {
	result := map[K]V{}
	e.Each(func(kv KeyValuePair[K, V], _ int) {
		result[kv.Key] = kv.Value
	})
	return result
}

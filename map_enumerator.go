package enu

type MapEnumerator[K comparable, V any] struct {
	collection map[K]V
	index      int
	keys       []K
}

func NewMapEnumerator[K comparable, V any](collection map[K]V) *MapEnumerator[K, V] {
	keys := make([]K, 0, len(collection))
	for k := range collection {
		keys = append(keys, k)
	}
	return &MapEnumerator[K, V]{collection: collection, index: 0, keys: keys}
}

func (e *MapEnumerator[K, V]) Reset() {
	e.index = 0
}

func (e *MapEnumerator[K, V]) Dispose() {
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

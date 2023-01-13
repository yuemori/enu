// This code was generated by a tool.
package enu

type MapEnumerable[K comparable, V any] struct {
	enumerator IEnumerator[KeyValuePair[K, V]]
}

func NewMap[K comparable, V any](e IEnumerator[KeyValuePair[K, V]]) *MapEnumerable[K, V] {
	return &MapEnumerable[K, V]{enumerator: e}
}

func (e *MapEnumerable[K, V]) Each(iteratee func(KeyValuePair[K, V], int)) *MapEnumerable[K, V] {
	each[KeyValuePair[K, V]](e, func(item KeyValuePair[K, V], index int) bool {
		iteratee(item, index)
		return true
	})
	return e
}

func (e *MapEnumerable[K, V]) ToSlice() []KeyValuePair[K, V] {
	return ToSlice[KeyValuePair[K, V]](e)
}

func (e *MapEnumerable[K, V]) Count() int {
	return Count[KeyValuePair[K, V]](e)
}

func (e *MapEnumerable[K, V]) Filter(predicate func(KeyValuePair[K, V], int) bool) *MapEnumerable[K, V] {
	return &MapEnumerable[K, V]{
		enumerator: Filter[KeyValuePair[K, V]](e, predicate),
	}
}

func (e *MapEnumerable[K, V]) Nth(nth int) (KeyValuePair[K, V], bool) {
	return Nth[KeyValuePair[K, V]](e, nth)
}

func (e *MapEnumerable[K, V]) Find(predicate func(KeyValuePair[K, V], int) bool) (KeyValuePair[K, V], bool) {
	return Find[KeyValuePair[K, V]](e, predicate)
}

func (e *MapEnumerable[K, V]) First() (KeyValuePair[K, V], bool) {
	return First[KeyValuePair[K, V]](e)
}

func (e *MapEnumerable[K, V]) Last() (KeyValuePair[K, V], bool) {
	return Last[KeyValuePair[K, V]](e)
}

func (e *MapEnumerable[K, V]) Reverse() *MapEnumerable[K, V] {
	return &MapEnumerable[K, V]{enumerator: Reverse[KeyValuePair[K, V]](e)}
}

func (e *MapEnumerable[K, V]) SortBy(sorter func(i, j KeyValuePair[K, V]) bool) *MapEnumerable[K, V] {
	return &MapEnumerable[K, V]{enumerator: SortBy[KeyValuePair[K, V]](e, sorter)}
}

func (e *MapEnumerable[K, V]) Reject(predicate func(KeyValuePair[K, V], int) bool) *MapEnumerable[K, V] {
	return &MapEnumerable[K, V]{enumerator: Reject[KeyValuePair[K, V]](e, predicate)}
}

func (e *MapEnumerable[K, V]) IsAll(predicate func(KeyValuePair[K, V]) bool) bool {
	return IsAll[KeyValuePair[K, V]](e, predicate)
}

func (e *MapEnumerable[K, V]) IsAny(predicate func(KeyValuePair[K, V]) bool) bool {
	return IsAny[KeyValuePair[K, V]](e, predicate)
}

func (e *MapEnumerable[K, V]) Take(num uint) *MapEnumerable[K, V] {
	return &MapEnumerable[K, V]{enumerator: Take[KeyValuePair[K, V]](e, num)}
}

func (e *MapEnumerable[K, V]) Result(out *[]KeyValuePair[K, V]) *MapEnumerable[K, V] {
	Result[KeyValuePair[K, V]](e, out)

	return e
}

func (e *MapEnumerable[K, V]) Err() error {
	if p, ok := e.enumerator.(ErrorProvider); ok {
		return p.Err()
	}
	return nil
}

func (e *MapEnumerable[K, V]) GetEnumerator() IEnumerator[KeyValuePair[K, V]] {
	return e.enumerator
}

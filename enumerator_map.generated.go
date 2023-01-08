package enumerator

import (
	"sort"

	"github.com/samber/lo"
  
)

type IEnumerableMap[K comparable, V any] interface {
	Next() (KeyValuePair[K, V], bool)
}

type EnumeratorMap[K comparable, V any] struct {
	iter IEnumerableMap[K, V]
}

func NewMap[K comparable, V any](e IEnumerableMap[K, V]) *EnumeratorMap[K, V] {
	return &EnumeratorMap[K, V]{iter: e}
}

func (e *EnumeratorMap[K, V]) Each(iteratee func(item KeyValuePair[K, V], index int)) *EnumeratorMap[K, V] {
  eachMap(e.iter, iteratee)

	return e
}

func (e *EnumeratorMap[K, V]) Count() int {
	v := 0
	eachMap(e.iter, func(item KeyValuePair[K, V], _ int) {
		v += 1
	})
	return v
}

func (e *EnumeratorMap[K, V]) ToSlice() []KeyValuePair[K, V] {
	result := make([]KeyValuePair[K, V], 0)

	for {
		item, ok := e.iter.Next()
		if !ok {
			break
		}
		result = append(result, item)
	}
	return result
}

func (e *EnumeratorMap[K, V]) Filter(predicate func(item KeyValuePair[K, V], index int) bool) *EnumeratorMap[K, V] {
	e.iter = newSliceEnumerator(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorMap[K, V]) First() (KeyValuePair[K, V], bool) {
	item, ok := e.iter.Next()
	if !ok {
		var empty KeyValuePair[K, V]
		return empty, false
	}
	return item, true
}

func (e *EnumeratorMap[K, V]) Last() (KeyValuePair[K, V], bool) {
	prev, ok := e.iter.Next()
	if !ok {
		var empty KeyValuePair[K, V]
		return empty, false
	}
	for {
		item, ok := e.iter.Next()
		if !ok {
			return prev, true
		}
		prev = item
	}
}

func (e *EnumeratorMap[K, V]) Reverse() *EnumeratorMap[K, V] {
	e.iter = newSliceEnumerator(lo.Reverse(e.ToSlice()))
	return e
}

func (e *EnumeratorMap[K, V]) SortBy(sorter func(i, j KeyValuePair[K, V]) bool) *EnumeratorMap[K, V] {
	res := e.ToSlice()
	sort.SliceStable(res, func(i, j int) bool {
		return sorter(res[i], res[j])
	})
	e.iter = newSliceEnumerator(res)
	return e
}

func eachMap[K comparable, V any](iter IEnumerableMap[K, V], iteratee func(item KeyValuePair[K, V], index int)) {
	index := 0
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}
		iteratee(item, index)
		index += 1
	}
}


func (e *EnumeratorMap[K, V]) Reject(predicate func(item KeyValuePair[K, V], index int) bool) *EnumeratorMap[K, V] {
	e.iter = newSliceEnumerator(lo.Reject(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorMap[K, V]) IsAll(predicate func(item KeyValuePair[K, V]) bool) bool {
	flag := true
	for {
		item, ok := e.iter.Next()
		if !ok {
			break
		}
		if !predicate(item) {
			flag = false
			break
		}
	}
	return flag
}

func (e *EnumeratorMap[K, V]) IsAny(predicate func(item KeyValuePair[K, V]) bool) bool {
	flag := false
	for {
		item, ok := e.iter.Next()
		if !ok {
			break
		}
		if predicate(item) {
			flag = true
			break
		}
	}
	return flag
}

func (e *EnumeratorMap[K, V]) Take(num int) *EnumeratorMap[K, V] {
	result := []KeyValuePair[K, V]{}
	index := 0
	for {
		item, ok := e.iter.Next()
		if !ok {
			break
		}
		if index >= num {
			break
		}
		result = append(result, item)
		index += 1
	}
	e.iter = newSliceEnumerator(result)
	return e
}

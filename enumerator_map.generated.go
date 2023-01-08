package enumerator

import (
	"sort"

	"github.com/samber/lo"
  
)

type IEnumerableMap[K comparable, V any] interface {
	Next() (KeyValuePair[K, V], error)
	Stop()
}

type EnumeratorMap[K comparable, V any] struct {
	iter IEnumerableMap[K, V]
	err  error
}

func NewMap[K comparable, V any](e IEnumerableMap[K, V]) *EnumeratorMap[K, V] {
	return &EnumeratorMap[K, V]{iter: e}
}

func (e *EnumeratorMap[K, V]) Error() error {
	return e.err
}

func (e *EnumeratorMap[K, V]) Each(iteratee func(item KeyValuePair[K, V], index int)) *EnumeratorMap[K, V] {
	if e.err == nil {
		eachMap(e.iter, iteratee)
	}

	return e
}

func (e *EnumeratorMap[K, V]) Count() int {
	v := 0
	if e.err != nil {
		return v
	}
	eachMap(e.iter, func(item KeyValuePair[K, V], _ int) {
		v += 1
	})
	return v
}

func (e *EnumeratorMap[K, V]) ToSlice() []KeyValuePair[K, V] {
	result := make([]KeyValuePair[K, V], 0)
	if e.err != nil {
		return result
	}

	for {
		item, err := e.iter.Next()
		if err == Done {
			break
		}
		if err != nil {
			e.err = err
			break
		}
		result = append(result, item)
	}
	return result
}

func (e *EnumeratorMap[K, V]) Filter(predicate func(item KeyValuePair[K, V], index int) bool) *EnumeratorMap[K, V] {
	if e.err != nil {
		return e
	}
	e.iter = newSliceEnumerator(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorMap[K, V]) First() (KeyValuePair[K, V], bool) {
	if e.err != nil {
		var empty KeyValuePair[K, V]
		return empty, false
	}
	item, err := e.iter.Next()
	if err != nil {
		var empty KeyValuePair[K, V]
		if err != Done {
			e.err = err
		}
		return empty, false
	}
	return item, true
}

func (e *EnumeratorMap[K, V]) Last() (KeyValuePair[K, V], bool) {
	if e.err != nil {
		var empty KeyValuePair[K, V]
		return empty, false
	}
	prev, err := e.iter.Next()
	if err == Done {
		var empty KeyValuePair[K, V]
		return empty, false
	}
	for {
		item, err := e.iter.Next()
		if err == Done {
			return prev, true
		}
		prev = item
		if err != nil {
			var empty KeyValuePair[K, V]
			e.err = err
			return empty, false
		}
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
		item, err := iter.Next()
		if err == Done {
			break
		}
		iteratee(item, index)
		index += 1
	}
}

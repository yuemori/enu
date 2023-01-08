package enu

import (
	"sort"

	"github.com/samber/lo"
  
)

type IEnumerableMap[K comparable, V any] interface {
	Next() (KeyValuePair[K, V], bool)
	Stop()
	Reset()
}

type EnumeratorMap[K comparable, V any] struct {
	iter      IEnumerableMap[K, V]
	result    []KeyValuePair[K, V]
	isStopped bool
}

func NewMap[K comparable, V any](e IEnumerableMap[K, V]) *EnumeratorMap[K, V] {
	return &EnumeratorMap[K, V]{iter: e}
}

func (e *EnumeratorMap[K, V]) Each(iteratee func(item KeyValuePair[K, V], index int)) {
	if e.isStopped {
		lo.ForEach(e.result, iteratee)
		return
	}

	result := []KeyValuePair[K, V]{}
	index := 0
	for {
		item, ok := e.iter.Next()
		if !ok {
			break
		}
		iteratee(item, index)
		index += 1
		result = append(result, item)
	}
	e.iter.Stop()
	e.isStopped = true
	e.iter = newSliceEnumerator(result)
	e.result = result
}

func (e *EnumeratorMap[K, V]) ToSlice() []KeyValuePair[K, V] {
	if e.isStopped {
		return e.result
	}
	e.Each(func(KeyValuePair[K, V], int) {})
	return e.result
}

func (e *EnumeratorMap[K, V]) Count() int {
	v := 0
	e.Each(func(item KeyValuePair[K, V], _ int) {
		v += 1
	})
	return v
}

func (e *EnumeratorMap[K, V]) Filter(predicate func(item KeyValuePair[K, V], index int) bool) *EnumeratorMap[K, V] {
	e.swap(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorMap[K, V]) Nth(index int) (KeyValuePair[K, V], bool) {
	item, err := lo.Nth(e.ToSlice(),index)
	if err != nil {
		return empty[KeyValuePair[K, V]](), false
	}

	return item, true
}

func (e *EnumeratorMap[K, V]) First() (KeyValuePair[K, V], bool) {
	if e.isStopped {
		if len(e.result) == 0 {
			return empty[KeyValuePair[K, V]](), false
		}
		return e.result[0], true
	}
	item, ok := e.iter.Next()
	if !ok {
		e.swap([]KeyValuePair[K, V]{})
		e.iter.Stop()
		e.isStopped = true
		return empty[KeyValuePair[K, V]](), false
	}
	e.iter.Reset()
	return item, true
}

func (e *EnumeratorMap[K, V]) Last() (KeyValuePair[K, V], bool) {
	result := e.ToSlice()
	if len(result) == 0 {
		return empty[KeyValuePair[K, V]](), false
	}
	return result[len(result)-1], true
}

func (e *EnumeratorMap[K, V]) Reverse() *EnumeratorMap[K, V] {
	e.swap(lo.Reverse(e.ToSlice()))
	return e
}

func (e *EnumeratorMap[K, V]) SortBy(sorter func(i, j KeyValuePair[K, V]) bool) *EnumeratorMap[K, V] {
	res := e.ToSlice()
	sort.SliceStable(res, func(i, j int) bool {
		return sorter(res[i], res[j])
	})
	e.swap(res)
	return e
}

func (e *EnumeratorMap[K, V]) Reject(predicate func(item KeyValuePair[K, V], index int) bool) *EnumeratorMap[K, V] {
	e.swap(lo.Reject(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorMap[K, V]) IsAll(predicate func(item KeyValuePair[K, V]) bool) bool {
	if e.isStopped {
		for _, item := range e.ToSlice() {
			if !predicate(item) {
				return false
			}
			return true
		}
	}

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
	e.iter.Reset()
	return flag
}

func (e *EnumeratorMap[K, V]) IsAny(predicate func(item KeyValuePair[K, V]) bool) bool {
	if e.isStopped {
		for _, item := range e.ToSlice() {
			if predicate(item) {
				return true
			}
			return false
		}
	}

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
	e.iter.Reset()
	return flag
}

func (e *EnumeratorMap[K, V]) Take(num uint) *EnumeratorMap[K, V] {
	if e.isStopped {
		e.swap(lo.Subset(e.result, 0, num-1))
		return e
	}

	result := []KeyValuePair[K, V]{}
	index := 0
	for {
		item, ok := e.iter.Next()
		if !ok {
			break
		}
		if uint(index) >= num {
			break
		}
		result = append(result, item)
		index += 1
	}
	e.swap(result)
	return e
}

func (e *EnumeratorMap[K, V]) swap(result []KeyValuePair[K, V]) {
	if !e.isStopped {
		e.iter.Stop()
		e.isStopped = true
	}
	e.iter = newSliceEnumerator(result)
	e.result = result
}

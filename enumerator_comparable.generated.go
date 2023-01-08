package enumerator

import (
	"sort"

	"github.com/samber/lo"
  
)

type IEnumerableComparable[T comparable] interface {
	Next() (T, bool)
	Stop()
}

type EnumeratorComparable[T comparable] struct {
	iter IEnumerableComparable[T]
}

func NewComparable[T comparable](e IEnumerableComparable[T]) *EnumeratorComparable[T] {
	return &EnumeratorComparable[T]{iter: e}
}

func (e *EnumeratorComparable[T]) Each(iteratee func(item T, index int)) *EnumeratorComparable[T] {
  eachComparable(e.iter, iteratee)

	return e
}

func (e *EnumeratorComparable[T]) Count() int {
	v := 0
	eachComparable(e.iter, func(item T, _ int) {
		v += 1
	})
	return v
}

func (e *EnumeratorComparable[T]) ToSlice() []T {
	result := make([]T, 0)

	for {
		item, ok := e.iter.Next()
		if !ok {
			break
		}
		result = append(result, item)
	}
	return result
}

func (e *EnumeratorComparable[T]) Filter(predicate func(item T, index int) bool) *EnumeratorComparable[T] {
	e.iter = newSliceEnumerator(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorComparable[T]) First() (T, bool) {
	item, ok := e.iter.Next()
	if !ok {
		var empty T
		return empty, false
	}
	return item, true
}

func (e *EnumeratorComparable[T]) Last() (T, bool) {
	prev, ok := e.iter.Next()
	if !ok {
		var empty T
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

func (e *EnumeratorComparable[T]) Reverse() *EnumeratorComparable[T] {
	e.iter = newSliceEnumerator(lo.Reverse(e.ToSlice()))
	return e
}

func (e *EnumeratorComparable[T]) SortBy(sorter func(i, j T) bool) *EnumeratorComparable[T] {
	res := e.ToSlice()
	sort.SliceStable(res, func(i, j int) bool {
		return sorter(res[i], res[j])
	})
	e.iter = newSliceEnumerator(res)
	return e
}

func eachComparable[T comparable](iter IEnumerableComparable[T], iteratee func(item T, index int)) {
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


func (e *EnumeratorComparable[T]) Reject(predicate func(item T, index int) bool) *EnumeratorComparable[T] {
	e.iter = newSliceEnumerator(lo.Reject(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorComparable[T]) IsAll(predicate func(item T) bool) bool {
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

func (e *EnumeratorComparable[T]) IsAny(predicate func(item T) bool) bool {
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

func (e *EnumeratorComparable[T]) Take(num int) *EnumeratorComparable[T] {
	result := []T{}
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

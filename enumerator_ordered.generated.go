package enumerator

import (
	"sort"

	"github.com/samber/lo"
  "golang.org/x/exp/constraints"
)

type IEnumerableOrdered[T constraints.Ordered] interface {
	Next() (T, bool)
	Stop()
}

type EnumeratorOrdered[T constraints.Ordered] struct {
	iter IEnumerableOrdered[T]
}

func NewOrdered[T constraints.Ordered](e IEnumerableOrdered[T]) *EnumeratorOrdered[T] {
	return &EnumeratorOrdered[T]{iter: e}
}

func (e *EnumeratorOrdered[T]) Each(iteratee func(item T, index int)) *EnumeratorOrdered[T] {
  eachOrdered(e.iter, iteratee)

	return e
}

func (e *EnumeratorOrdered[T]) Count() int {
	v := 0
	eachOrdered(e.iter, func(item T, _ int) {
		v += 1
	})
	return v
}

func (e *EnumeratorOrdered[T]) ToSlice() []T {
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

func (e *EnumeratorOrdered[T]) Filter(predicate func(item T, index int) bool) *EnumeratorOrdered[T] {
	e.iter = newSliceEnumerator(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorOrdered[T]) First() (T, bool) {
	item, ok := e.iter.Next()
	if !ok {
		var empty T
		return empty, false
	}
	return item, true
}

func (e *EnumeratorOrdered[T]) Last() (T, bool) {
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

func (e *EnumeratorOrdered[T]) Reverse() *EnumeratorOrdered[T] {
	e.iter = newSliceEnumerator(lo.Reverse(e.ToSlice()))
	return e
}

func (e *EnumeratorOrdered[T]) SortBy(sorter func(i, j T) bool) *EnumeratorOrdered[T] {
	res := e.ToSlice()
	sort.SliceStable(res, func(i, j int) bool {
		return sorter(res[i], res[j])
	})
	e.iter = newSliceEnumerator(res)
	return e
}

func eachOrdered[T constraints.Ordered](iter IEnumerableOrdered[T], iteratee func(item T, index int)) {
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


func (e *EnumeratorOrdered[T]) Reject(predicate func(item T, index int) bool) *EnumeratorOrdered[T] {
	e.iter = newSliceEnumerator(lo.Reject(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorOrdered[T]) IsAll(predicate func(item T) bool) bool {
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

func (e *EnumeratorOrdered[T]) IsAny(predicate func(item T) bool) bool {
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

func (e *EnumeratorOrdered[T]) Take(num int) *EnumeratorOrdered[T] {
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

package enumerator

import (
	"sort"

	"github.com/samber/lo"
  
)

type IEnumerable[T any] interface {
	Next() (T, bool)
	Stop()
}

type Enumerator[T any] struct {
	iter IEnumerable[T]
}

func New[T any](e IEnumerable[T]) *Enumerator[T] {
	return &Enumerator[T]{iter: e}
}

func (e *Enumerator[T]) Each(iteratee func(item T, index int)) *Enumerator[T] {
  each(e.iter, iteratee)

	return e
}

func (e *Enumerator[T]) Count() int {
	v := 0
	each(e.iter, func(item T, _ int) {
		v += 1
	})
	return v
}

func (e *Enumerator[T]) ToSlice() []T {
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

func (e *Enumerator[T]) Filter(predicate func(item T, index int) bool) *Enumerator[T] {
	e.iter = newSliceEnumerator(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *Enumerator[T]) First() (T, bool) {
	item, ok := e.iter.Next()
	if !ok {
		var empty T
		return empty, false
	}
	return item, true
}

func (e *Enumerator[T]) Last() (T, bool) {
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

func (e *Enumerator[T]) Reverse() *Enumerator[T] {
	e.iter = newSliceEnumerator(lo.Reverse(e.ToSlice()))
	return e
}

func (e *Enumerator[T]) SortBy(sorter func(i, j T) bool) *Enumerator[T] {
	res := e.ToSlice()
	sort.SliceStable(res, func(i, j int) bool {
		return sorter(res[i], res[j])
	})
	e.iter = newSliceEnumerator(res)
	return e
}

func each[T any](iter IEnumerable[T], iteratee func(item T, index int)) {
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


func (e *Enumerator[T]) Reject(predicate func(item T, index int) bool) *Enumerator[T] {
	e.iter = newSliceEnumerator(lo.Reject(e.ToSlice(), predicate))
	return e
}

func (e *Enumerator[T]) IsAll(predicate func(item T) bool) bool {
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

func (e *Enumerator[T]) IsAny(predicate func(item T) bool) bool {
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

func (e *Enumerator[T]) Take(num int) *Enumerator[T] {
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

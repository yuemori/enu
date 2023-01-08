package enumerator

import (
	"sort"

	"github.com/samber/lo"
  "golang.org/x/exp/constraints"
)

type IEnumerableNumeric[T constraints.Integer | constraints.Float] interface {
	Next() (T, bool)
	Stop()
}

type EnumeratorNumeric[T constraints.Integer | constraints.Float] struct {
	iter IEnumerableNumeric[T]
}

func NewNumeric[T constraints.Integer | constraints.Float](e IEnumerableNumeric[T]) *EnumeratorNumeric[T] {
	return &EnumeratorNumeric[T]{iter: e}
}

func (e *EnumeratorNumeric[T]) Each(iteratee func(item T, index int)) *EnumeratorNumeric[T] {
  eachNumeric(e.iter, iteratee)

	return e
}

func (e *EnumeratorNumeric[T]) Count() int {
	v := 0
	eachNumeric(e.iter, func(item T, _ int) {
		v += 1
	})
	return v
}

func (e *EnumeratorNumeric[T]) ToSlice() []T {
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

func (e *EnumeratorNumeric[T]) Filter(predicate func(item T, index int) bool) *EnumeratorNumeric[T] {
	e.iter = newSliceEnumerator(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorNumeric[T]) First() (T, bool) {
	item, ok := e.iter.Next()
	if !ok {
		var empty T
		return empty, false
	}
	return item, true
}

func (e *EnumeratorNumeric[T]) Last() (T, bool) {
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

func (e *EnumeratorNumeric[T]) Reverse() *EnumeratorNumeric[T] {
	e.iter = newSliceEnumerator(lo.Reverse(e.ToSlice()))
	return e
}

func (e *EnumeratorNumeric[T]) SortBy(sorter func(i, j T) bool) *EnumeratorNumeric[T] {
	res := e.ToSlice()
	sort.SliceStable(res, func(i, j int) bool {
		return sorter(res[i], res[j])
	})
	e.iter = newSliceEnumerator(res)
	return e
}

func eachNumeric[T constraints.Integer | constraints.Float](iter IEnumerableNumeric[T], iteratee func(item T, index int)) {
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


func (e *EnumeratorNumeric[T]) Reject(predicate func(item T, index int) bool) *EnumeratorNumeric[T] {
	e.iter = newSliceEnumerator(lo.Reject(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorNumeric[T]) IsAll(predicate func(item T) bool) bool {
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

func (e *EnumeratorNumeric[T]) IsAny(predicate func(item T) bool) bool {
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

func (e *EnumeratorNumeric[T]) Take(num int) *EnumeratorNumeric[T] {
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

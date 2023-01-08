package enu

import (
	"sort"

	"github.com/samber/lo"
  "golang.org/x/exp/constraints"
)

type IEnumerableOrdered[T constraints.Ordered] interface {
	Next() (T, bool)
	Stop()
	Reset()
}

type EnumeratorOrdered[T constraints.Ordered] struct {
	iter      IEnumerableOrdered[T]
	result    []T
	isStopped bool
}

func NewOrdered[T constraints.Ordered](e IEnumerableOrdered[T]) *EnumeratorOrdered[T] {
	return &EnumeratorOrdered[T]{iter: e}
}

func (e *EnumeratorOrdered[T]) Each(iteratee func(item T, index int)) {
	if e.isStopped {
		lo.ForEach(e.result, iteratee)
		return
	}

	result := []T{}
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

func (e *EnumeratorOrdered[T]) ToSlice() []T {
	if e.isStopped {
		return e.result
	}
	e.Each(func(T, int) {})
	return e.result
}

func (e *EnumeratorOrdered[T]) Count() int {
	v := 0
	e.Each(func(item T, _ int) {
		v += 1
	})
	return v
}

func (e *EnumeratorOrdered[T]) Filter(predicate func(item T, index int) bool) *EnumeratorOrdered[T] {
	e.swap(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorOrdered[T]) Nth(index int) (T, bool) {
	item, err := lo.Nth(e.ToSlice(),index)
	if err != nil {
		return empty[T](), false
	}

	return item, true
}

func (e *EnumeratorOrdered[T]) Find(predicate func(item T) bool) (T, bool) {
	if e.isStopped {
		return lo.Find(e.result, predicate)
	}

	result := []T{}
	for {
		item, ok := e.iter.Next()
		if !ok {
			break
		}
		if predicate(item) {
			e.iter.Reset()
			return item, true
		}
	}
	e.isStopped = true
	e.swap(result)
	return empty[T](), false
}

func (e *EnumeratorOrdered[T]) First() (T, bool) {
	if e.isStopped {
		if len(e.result) == 0 {
			return empty[T](), false
		}
		return e.result[0], true
	}
	item, ok := e.iter.Next()
	if !ok {
		e.swap([]T{})
		e.iter.Stop()
		e.isStopped = true
		return empty[T](), false
	}
	e.iter.Reset()
	return item, true
}

func (e *EnumeratorOrdered[T]) Last() (T, bool) {
	result := e.ToSlice()
	if len(result) == 0 {
		return empty[T](), false
	}
	return result[len(result)-1], true
}

func (e *EnumeratorOrdered[T]) Reverse() *EnumeratorOrdered[T] {
	e.swap(lo.Reverse(e.ToSlice()))
	return e
}

func (e *EnumeratorOrdered[T]) SortBy(sorter func(i, j T) bool) *EnumeratorOrdered[T] {
	res := e.ToSlice()
	sort.SliceStable(res, func(i, j int) bool {
		return sorter(res[i], res[j])
	})
	e.swap(res)
	return e
}

func (e *EnumeratorOrdered[T]) Reject(predicate func(item T, index int) bool) *EnumeratorOrdered[T] {
	e.swap(lo.Reject(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorOrdered[T]) IsAll(predicate func(item T) bool) bool {
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

func (e *EnumeratorOrdered[T]) IsAny(predicate func(item T) bool) bool {
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

func (e *EnumeratorOrdered[T]) Take(num uint) *EnumeratorOrdered[T] {
	if e.isStopped {
		e.swap(lo.Subset(e.result, 0, num-1))
		return e
	}

	result := []T{}
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

func (e *EnumeratorOrdered[T]) swap(result []T) {
	if !e.isStopped {
		e.iter.Stop()
		e.isStopped = true
	}
	e.iter = newSliceEnumerator(result)
	e.result = result
}

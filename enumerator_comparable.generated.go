package enu

import (
	"sort"

	"github.com/samber/lo"
  
)

type IEnumerableComparable[T comparable] interface {
	Next() (T, bool)
	Stop()
	Reset()
}

type EnumeratorComparable[T comparable] struct {
	iter      IEnumerableComparable[T]
	result    []T
	isStopped bool
}

func NewComparable[T comparable](e IEnumerableComparable[T]) *EnumeratorComparable[T] {
	return &EnumeratorComparable[T]{iter: e}
}

func (e *EnumeratorComparable[T]) Each(iteratee func(item T, index int)) {
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

func (e *EnumeratorComparable[T]) ToSlice() []T {
	if e.isStopped {
		return e.result
	}
	e.Each(func(T, int) {})
	return e.result
}

func (e *EnumeratorComparable[T]) Count() int {
	v := 0
	e.Each(func(item T, _ int) {
		v += 1
	})
	return v
}

func (e *EnumeratorComparable[T]) Filter(predicate func(item T, index int) bool) *EnumeratorComparable[T] {
	e.swap(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorComparable[T]) First() (T, bool) {
	if e.isStopped {
		if len(e.result) == 0 {
			return empty[T](), false
		}
		return e.result[0], true
	}
	item, ok := e.iter.Next()
	e.iter.Reset()
	if !ok {
		return empty[T](), false
	}
	return item, true
}

func (e *EnumeratorComparable[T]) Last() (T, bool) {
	result := e.ToSlice()
	if len(result) == 0 {
		return empty[T](), false
	}
	return result[len(result)-1], true
}

func (e *EnumeratorComparable[T]) Reverse() *EnumeratorComparable[T] {
	e.swap(lo.Reverse(e.ToSlice()))
	return e
}

func (e *EnumeratorComparable[T]) SortBy(sorter func(i, j T) bool) *EnumeratorComparable[T] {
	res := e.ToSlice()
	sort.SliceStable(res, func(i, j int) bool {
		return sorter(res[i], res[j])
	})
	e.swap(res)
	return e
}

func (e *EnumeratorComparable[T]) Reject(predicate func(item T, index int) bool) *EnumeratorComparable[T] {
	e.swap(lo.Reject(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorComparable[T]) IsAll(predicate func(item T) bool) bool {
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

func (e *EnumeratorComparable[T]) IsAny(predicate func(item T) bool) bool {
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

func (e *EnumeratorComparable[T]) Take(num uint) *EnumeratorComparable[T] {
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

func (e *EnumeratorComparable[T]) swap(result []T) {
	if !e.isStopped {
		e.iter.Stop()
		e.isStopped = true
	}
	e.iter = newSliceEnumerator(result)
	e.result = result
}

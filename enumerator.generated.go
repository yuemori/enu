package enu

import (
	"sort"

	"github.com/samber/lo"
  
)

type IEnumerable[T any] interface {
	Next() (T, bool)
	Stop()
	Reset()
}

type Enumerator[T any] struct {
	iter      IEnumerable[T]
	result    []T
	isStopped bool
}

func New[T any](e IEnumerable[T]) *Enumerator[T] {
	return &Enumerator[T]{iter: e}
}

func (e *Enumerator[T]) Each(iteratee func(item T, index int)) {
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

func (e *Enumerator[T]) ToSlice() []T {
	if e.isStopped {
		return e.result
	}
	e.Each(func(T, int) {})
	return e.result
}

func (e *Enumerator[T]) Count() int {
	v := 0
	e.Each(func(item T, _ int) {
		v += 1
	})
	return v
}

func (e *Enumerator[T]) Filter(predicate func(item T, index int) bool) *Enumerator[T] {
	e.swap(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *Enumerator[T]) First() (T, bool) {
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

func (e *Enumerator[T]) Last() (T, bool) {
	result := e.ToSlice()
	if len(result) == 0 {
		return empty[T](), false
	}
	return result[len(result)-1], true
}

func (e *Enumerator[T]) Reverse() *Enumerator[T] {
	e.swap(lo.Reverse(e.ToSlice()))
	return e
}

func (e *Enumerator[T]) SortBy(sorter func(i, j T) bool) *Enumerator[T] {
	res := e.ToSlice()
	sort.SliceStable(res, func(i, j int) bool {
		return sorter(res[i], res[j])
	})
	e.swap(res)
	return e
}

func (e *Enumerator[T]) Reject(predicate func(item T, index int) bool) *Enumerator[T] {
	e.swap(lo.Reject(e.ToSlice(), predicate))
	return e
}

func (e *Enumerator[T]) IsAll(predicate func(item T) bool) bool {
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

func (e *Enumerator[T]) IsAny(predicate func(item T) bool) bool {
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

func (e *Enumerator[T]) Take(num uint) *Enumerator[T] {
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

func (e *Enumerator[T]) swap(result []T) {
	if !e.isStopped {
		e.iter.Stop()
		e.isStopped = true
	}
	e.iter = newSliceEnumerator(result)
	e.result = result
}

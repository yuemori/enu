package enumerator

import (
	"sort"

	"github.com/samber/lo"
  
)

type IEnumerableComparable[T comparable] interface {
	Next() (T, error)
	Stop()
}

type EnumeratorComparable[T comparable] struct {
	iter IEnumerableComparable[T]
	err  error
}

func NewComparable[T comparable](e IEnumerableComparable[T]) *EnumeratorComparable[T] {
	return &EnumeratorComparable[T]{iter: e}
}

func (e *EnumeratorComparable[T]) Error() error {
	return e.err
}

func (e *EnumeratorComparable[T]) Each(iteratee func(item T, index int)) *EnumeratorComparable[T] {
	if e.err == nil {
		eachComparable(e.iter, iteratee)
	}

	return e
}

func (e *EnumeratorComparable[T]) Count() int {
	v := 0
	if e.err != nil {
		return v
	}
	eachComparable(e.iter, func(item T, _ int) {
		v += 1
	})
	return v
}

func (e *EnumeratorComparable[T]) ToSlice() []T {
	result := make([]T, 0)
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

func (e *EnumeratorComparable[T]) Filter(predicate func(item T, index int) bool) *EnumeratorComparable[T] {
	if e.err != nil {
		return e
	}
	e.iter = newSliceEnumerator(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorComparable[T]) First() (T, bool) {
	if e.err != nil {
		var empty T
		return empty, false
	}
	item, err := e.iter.Next()
	if err != nil {
		var empty T
		if err != Done {
			e.err = err
		}
		return empty, false
	}
	return item, true
}

func (e *EnumeratorComparable[T]) Last() (T, bool) {
	if e.err != nil {
		var empty T
		return empty, false
	}
	prev, err := e.iter.Next()
	if err == Done {
		var empty T
		return empty, false
	}
	for {
		item, err := e.iter.Next()
		if err == Done {
			return prev, true
		}
		prev = item
		if err != nil {
			var empty T
			e.err = err
			return empty, false
		}
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
		item, err := iter.Next()
		if err == Done {
			break
		}
		iteratee(item, index)
		index += 1
	}
}

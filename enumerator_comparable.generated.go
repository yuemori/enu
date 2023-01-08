package enumerator

import (
	"sort"

	"github.com/samber/lo"
)

type IEnumerableC[T comparable] interface {
	Next() (T, error)
	Stop()
}

type EnumeratorC[T comparable] struct {
	iter IEnumerableC[T]
	err  error
}

func NewC[T comparable](e IEnumerableC[T]) *EnumeratorC[T] {
	return &EnumeratorC[T]{iter: e}
}

func (e *EnumeratorC[T]) Error() error {
	return e.err
}

func (e *EnumeratorC[T]) Each(iteratee func(item T, index int)) *EnumeratorC[T] {
	if e.err == nil {
		eachC(e.iter, iteratee)
	}

	return e
}

func (e *EnumeratorC[T]) Count() int {
	v := 0
	if e.err != nil {
		return v
	}
	eachC(e.iter, func(item T, _ int) {
		v += 1
	})
	return v
}

func (e *EnumeratorC[T]) ToSlice() []T {
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

func (e *EnumeratorC[T]) Filter(predicate func(item T, index int) bool) *EnumeratorC[T] {
	if e.err != nil {
		return e
	}
	e.iter = newSliceEnumerator(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorC[T]) First() (T, bool) {
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

func (e *EnumeratorC[T]) Last() (T, bool) {
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

func (e *EnumeratorC[T]) SortBy(sorter func(i, j T) bool) *EnumeratorC[T] {
	res := e.ToSlice()
	sort.SliceStable(res, func(i, j int) bool {
		return sorter(res[i], res[j])
	})
	e.iter = newSliceEnumerator(res)
	return e
}

func eachC[T comparable](iter IEnumerableC[T], iteratee func(item T, index int)) {
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

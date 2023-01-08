package enumerator

import (
	"sort"

	"github.com/samber/lo"
  "golang.org/x/exp/constraints"
)

type IEnumerableOrdered[T constraints.Ordered] interface {
	Next() (T, error)
	Stop()
}

type EnumeratorOrdered[T constraints.Ordered] struct {
	iter IEnumerableOrdered[T]
	err  error
}

func NewOrdered[T constraints.Ordered](e IEnumerableOrdered[T]) *EnumeratorOrdered[T] {
	return &EnumeratorOrdered[T]{iter: e}
}

func (e *EnumeratorOrdered[T]) Error() error {
	return e.err
}

func (e *EnumeratorOrdered[T]) Each(iteratee func(item T, index int)) *EnumeratorOrdered[T] {
	if e.err == nil {
		eachOrdered(e.iter, iteratee)
	}

	return e
}

func (e *EnumeratorOrdered[T]) Count() int {
	v := 0
	if e.err != nil {
		return v
	}
	eachOrdered(e.iter, func(item T, _ int) {
		v += 1
	})
	return v
}

func (e *EnumeratorOrdered[T]) ToSlice() []T {
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

func (e *EnumeratorOrdered[T]) Filter(predicate func(item T, index int) bool) *EnumeratorOrdered[T] {
	if e.err != nil {
		return e
	}
	e.iter = newSliceEnumerator(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorOrdered[T]) First() (T, bool) {
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

func (e *EnumeratorOrdered[T]) Last() (T, bool) {
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
		item, err := iter.Next()
		if err == Done {
			break
		}
		iteratee(item, index)
		index += 1
	}
}

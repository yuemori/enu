package enumerator

import (
	"sort"

	"github.com/samber/lo"
)

type IEnumerable[T any] interface {
	Next() (T, error)
	Stop()
}

type Enumerator[T any] struct {
	iter IEnumerable[T]
	err  error
}

func New[T any](e IEnumerable[T]) *Enumerator[T] {
	return &Enumerator[T]{iter: e}
}

func (e *Enumerator[T]) Error() error {
	return e.err
}

func (e *Enumerator[T]) Each(iteratee func(item T, index int)) *Enumerator[T] {
	if e.err == nil {
		each(e.iter, iteratee)
	}

	return e
}

func (e *Enumerator[T]) Count() int {
	v := 0
	if e.err != nil {
		return v
	}
	each(e.iter, func(item T, _ int) {
		v += 1
	})
	return v
}

func (e *Enumerator[T]) ToSlice() []T {
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

func (e *Enumerator[T]) Filter(predicate func(item T, index int) bool) *Enumerator[T] {
	if e.err != nil {
		return e
	}
	e.iter = newSliceEnumerator(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *Enumerator[T]) First() (T, bool) {
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

func (e *Enumerator[T]) Last() (T, bool) {
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
		item, err := iter.Next()
		if err == Done {
			break
		}
		iteratee(item, index)
		index += 1
	}
}

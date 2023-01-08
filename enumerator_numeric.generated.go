package enumerator

import (
	"sort"

	"github.com/samber/lo"
  "golang.org/x/exp/constraints"
)

type IEnumerableNumeric[T constraints.Integer | constraints.Float] interface {
	Next() (T, error)
	Stop()
}

type EnumeratorNumeric[T constraints.Integer | constraints.Float] struct {
	iter IEnumerableNumeric[T]
	err  error
}

func NewNumeric[T constraints.Integer | constraints.Float](e IEnumerableNumeric[T]) *EnumeratorNumeric[T] {
	return &EnumeratorNumeric[T]{iter: e}
}

func (e *EnumeratorNumeric[T]) Error() error {
	return e.err
}

func (e *EnumeratorNumeric[T]) Each(iteratee func(item T, index int)) *EnumeratorNumeric[T] {
	if e.err == nil {
		eachNumeric(e.iter, iteratee)
	}

	return e
}

func (e *EnumeratorNumeric[T]) Count() int {
	v := 0
	if e.err != nil {
		return v
	}
	eachNumeric(e.iter, func(item T, _ int) {
		v += 1
	})
	return v
}

func (e *EnumeratorNumeric[T]) ToSlice() []T {
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

func (e *EnumeratorNumeric[T]) Filter(predicate func(item T, index int) bool) *EnumeratorNumeric[T] {
	if e.err != nil {
		return e
	}
	e.iter = newSliceEnumerator(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorNumeric[T]) First() (T, bool) {
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

func (e *EnumeratorNumeric[T]) Last() (T, bool) {
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
		item, err := iter.Next()
		if err == Done {
			break
		}
		iteratee(item, index)
		index += 1
	}
}


func (e *EnumeratorNumeric[T]) Reject(predicate func(item T, index int) bool) *EnumeratorNumeric[T] {
	if e.err != nil {
		return e
	}
	e.iter = newSliceEnumerator(lo.Reject(e.ToSlice(), predicate))
	return e
}

func (e *EnumeratorNumeric[T]) IsAll(predicate func(item T) bool) bool {
	if e.err != nil {
		return false
	}
	flag := true
	for {
		item, err := e.iter.Next()
		if err == Done {
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
	if e.err != nil {
		return false
	}
	flag := false
	for {
		item, err := e.iter.Next()
		if err == Done {
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
		item, err := e.iter.Next()
		if err == Done {
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

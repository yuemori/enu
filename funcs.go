package enu

import (
	"sort"

	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

func Each[T any](iter IEnumerable[T], iteratee func(item T, index int)) {
	each(iter, func(item T, index int) bool {
		iteratee(item, index)
		return true
	})
}

func each[T any](iter IEnumerable[T], iteratee func(item T, index int) bool) {
	defer iter.Stop()

	index := 0
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}
		if !iteratee(item, index) {
			break
		}
		index++
	}
}

func Nth[T any](iter IEnumerable[T], nth int) (T, bool) {
	defer iter.Reset()

	index := 0
	for {
		item, ok := iter.Next()
		if !ok {
			return empty[T](), false
		}
		if nth == index {
			return item, true
		}
		index++
	}
}

func Count[T any](iter IEnumerable[T]) int {
	v := 0
	Each(iter, func(item T, _ int) {
		v += 1
	})
	return v
}

func Find[T any](iter IEnumerable[T], predicate func(T, int) bool) (T, bool) {
	result := empty[T]()
	ok := false
	each(iter, func(item T, index int) bool {
		if predicate(item, index) {
			result = item
			ok = true
			return false
		}
		return true
	})
	return result, ok
}

func First[T any](iter IEnumerable[T]) (T, bool) {
	defer iter.Reset()

	item, ok := iter.Next()
	if !ok {
		return empty[T](), false
	}
	return item, true
}

func Last[T any](iter IEnumerable[T]) (T, bool) {
	result := ToSlice(iter)
	if len(result) == 0 {
		return empty[T](), false
	}
	return result[len(result)-1], true
}

func ToSlice[T any](iter IEnumerable[T]) []T {
	result := make([]T, 0)
	each(iter, func(item T, _ int) bool {
		result = append(result, item)
		return true
	})
	return result
}

func Reverse[T any](iter IEnumerable[T]) *SliceEnumerator[T] {
	return NewSliceEnumerator(lo.Reverse(ToSlice(iter)))
}

func Sort[T constraints.Ordered](iter IEnumerable[T]) *SliceEnumerator[T] {
	res := ToSlice(iter)
	sort.SliceStable(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return NewSliceEnumerator(res)
}

func SortBy[T any](iter IEnumerable[T], sorter func(i, j T) bool) *SliceEnumerator[T] {
	res := ToSlice(iter)
	sort.SliceStable(res, func(i, j int) bool {
		return sorter(res[i], res[j])
	})
	return NewSliceEnumerator(res)
}

func Sum[T constraints.Integer | constraints.Float | constraints.Complex](iter IEnumerable[T]) T {
	return lo.Sum(ToSlice(iter))
}

func Min[T constraints.Ordered](iter IEnumerable[T]) T {
	return lo.Min(ToSlice(iter))
}

func Max[T constraints.Ordered](iter IEnumerable[T]) T {
	return lo.Max(ToSlice(iter))
}

func Uniq[T comparable](iter IEnumerable[T]) *SliceEnumerator[T] {
	return NewSliceEnumerator(lo.Uniq(ToSlice(iter)))
}

func Contains[T comparable](iter IEnumerable[T], element T) bool {
	ok := false
	each(iter, func(item T, _ int) bool {
		if item == element {
			ok = true
			return false
		}
		return true
	})
	return ok
}

func IndexOf[T comparable](iter IEnumerable[T], element T) int {
	i := -1
	each(iter, func(item T, index int) bool {
		if item == element {
			i = index
			return false
		}
		return true
	})
	return i
}

func IsAll[T any](iter IEnumerable[T], predicate func(item T) bool) bool {
	flag := true

	each(iter, func(item T, _ int) bool {
		if !predicate(item) {
			flag = false
			return false
		}
		return true
	})

	return flag
}

func IsAny[T any](iter IEnumerable[T], predicate func(item T) bool) bool {
	flag := false

	each(iter, func(item T, _ int) bool {
		if predicate(item) {
			flag = true
			return false
		}
		return true
	})

	return flag
}

func Filter[T any](iter IEnumerable[T], predicate func(item T, index int) bool) *FilterEnumerator[T] {
	return &FilterEnumerator[T]{
		iter:      iter,
		predicate: predicate,
	}
}

func Reject[T any](iter IEnumerable[T], predicate func(item T, index int) bool) *RejectEnumerator[T] {
	return &RejectEnumerator[T]{
		iter:      iter,
		predicate: predicate,
	}
}

func Take[T any](iter IEnumerable[T], size uint) *TakeEnumerator[T] {
	return &TakeEnumerator[T]{
		iter: iter,
		size: size,
	}
}

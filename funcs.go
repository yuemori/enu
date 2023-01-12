package enu

import (
	"sort"

	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

func Each[T any](e IEnumerable[T], iteratee func(item T, index int)) {
	each(e, func(item T, index int) bool {
		iteratee(item, index)
		return true
	})
}

func each[T any](enumerable IEnumerable[T], iteratee func(item T, index int) bool) {
	iter := enumerable.GetEnumerator()
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

func Nth[T any](e IEnumerable[T], nth int) (T, bool) {
	iter := e.GetEnumerator()
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

func Count[T any](e IEnumerable[T]) int {
	v := 0
	Each(e, func(item T, _ int) {
		v += 1
	})
	return v
}

func Find[T any](e IEnumerable[T], predicate func(T, int) bool) (T, bool) {
	result := empty[T]()
	ok := false
	each(e, func(item T, index int) bool {
		if predicate(item, index) {
			result = item
			ok = true
			return false
		}
		return true
	})
	return result, ok
}

func First[T any](e IEnumerable[T]) (T, bool) {
	iter := e.GetEnumerator()
	defer iter.Reset()

	item, ok := iter.Next()
	if !ok {
		return empty[T](), false
	}
	return item, true
}

func Last[T any](e IEnumerable[T]) (T, bool) {
	result := ToSlice(e)
	if len(result) == 0 {
		return empty[T](), false
	}
	return result[len(result)-1], true
}

func ToSlice[T any](e IEnumerable[T]) []T {
	result := make([]T, 0)
	each(e, func(item T, _ int) bool {
		result = append(result, item)
		return true
	})
	return result
}

func Reverse[T any](e IEnumerable[T]) *SliceEnumerator[T] {
	return NewSliceEnumerator(lo.Reverse(ToSlice(e)))
}

func Sort[T constraints.Ordered](e IEnumerable[T]) *SliceEnumerator[T] {
	res := ToSlice(e)
	sort.SliceStable(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return NewSliceEnumerator(res)
}

func SortBy[T any](e IEnumerable[T], sorter func(i, j T) bool) *SliceEnumerator[T] {
	res := ToSlice(e)
	sort.SliceStable(res, func(i, j int) bool {
		return sorter(res[i], res[j])
	})
	return NewSliceEnumerator(res)
}

func Sum[T constraints.Integer | constraints.Float | constraints.Complex](e IEnumerable[T]) T {
	return lo.Sum(ToSlice(e))
}

func Min[T constraints.Ordered](e IEnumerable[T]) T {
	return lo.Min(ToSlice(e))
}

func Max[T constraints.Ordered](e IEnumerable[T]) T {
	return lo.Max(ToSlice(e))
}

func Uniq[T comparable](e IEnumerable[T]) *SliceEnumerator[T] {
	return NewSliceEnumerator(lo.Uniq(ToSlice(e)))
}

func Contains[T comparable](e IEnumerable[T], element T) bool {
	ok := false
	each(e, func(item T, _ int) bool {
		if item == element {
			ok = true
			return false
		}
		return true
	})
	return ok
}

func IndexOf[T comparable](e IEnumerable[T], element T) int {
	i := -1
	each(e, func(item T, index int) bool {
		if item == element {
			i = index
			return false
		}
		return true
	})
	return i
}

func IsAll[T any](e IEnumerable[T], predicate func(item T) bool) bool {
	flag := true

	each(e, func(item T, _ int) bool {
		if !predicate(item) {
			flag = false
			return false
		}
		return true
	})

	return flag
}

func IsAny[T any](e IEnumerable[T], predicate func(item T) bool) bool {
	flag := false

	each(e, func(item T, _ int) bool {
		if predicate(item) {
			flag = true
			return false
		}
		return true
	})

	return flag
}

func Filter[T any](e IEnumerable[T], predicate func(item T, index int) bool) *FilterEnumerator[T] {
	return &FilterEnumerator[T]{
		iter:      e.GetEnumerator(),
		predicate: predicate,
	}
}

func Reject[T any](e IEnumerable[T], predicate func(item T, index int) bool) *RejectEnumerator[T] {
	return &RejectEnumerator[T]{
		iter:      e.GetEnumerator(),
		predicate: predicate,
	}
}

func Take[T any](e IEnumerable[T], size uint) *TakeEnumerator[T] {
	return &TakeEnumerator[T]{
		iter: e.GetEnumerator(),
		size: size,
	}
}

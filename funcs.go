package enu

import (
	"sort"

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
	defer iter.Dispose()

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
	if nth < 0 {
		collection := ToSlice(e)
		l := len(collection)
		if l < -nth {
			return empty[T](), false
		}
		return collection[l+nth], true
	}

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
	Each(e, func(item T, _ int) {
		result = append(result, item)
	})
	return result
}

func Reverse[T any](e IEnumerable[T]) *SliceEnumerator[T] {
	collection := ToSlice(e)
	length := len(collection)
	half := length / 2

	for i := 0; i < half; i = i + 1 {
		j := length - 1 - i
		collection[i], collection[j] = collection[j], collection[i]
	}

	return NewSliceEnumerator(collection)
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
	var sum T = 0
	for _, val := range ToSlice(e) {
		sum += val
	}
	return sum
}

func Min[T constraints.Ordered](e IEnumerable[T]) T {
	var min T
	collection := ToSlice(e)
	if len(collection) == 0 {
		return empty[T]()
	}
	min = collection[0]
	for _, item := range collection {
		if min > item {
			min = item
		}
	}
	return min
}

func Max[T constraints.Ordered](e IEnumerable[T]) T {
	var max T
	collection := ToSlice(e)
	if len(collection) == 0 {
		return empty[T]()
	}
	max = collection[0]
	for _, item := range collection {
		if max < item {
			max = item
		}
	}
	return max
}

func Uniq[T comparable](e IEnumerable[T]) *UniqEnumerable[T] {
	return &UniqEnumerable[T]{iter: e.GetEnumerator()}
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

func Filter[T any](e IEnumerable[T], predicate func(item T, index int) bool) *FilterEnumerable[T] {
	return &FilterEnumerable[T]{
		iter:      e.GetEnumerator(),
		predicate: predicate,
	}
}

func Reject[T any](e IEnumerable[T], predicate func(item T, index int) bool) *RejectEnumerable[T] {
	return &RejectEnumerable[T]{
		iter:      e.GetEnumerator(),
		predicate: predicate,
	}
}

func Take[T any](e IEnumerable[T], size uint) *TakeEnumerable[T] {
	return &TakeEnumerable[T]{
		iter: e.GetEnumerator(),
		size: size,
	}
}

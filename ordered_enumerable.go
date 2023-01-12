package enu

import (
	"golang.org/x/exp/constraints"
)

func FromOrdered[T constraints.Ordered](collection []T) *OrderedEnumerable[T] {
	return NewOrdered[T](NewSliceEnumerator(collection))
}

func ToOrdered[T constraints.Ordered](e IEnumerable[T]) *OrderedEnumerable[T] {
	return NewOrdered(e.GetEnumerator())
}

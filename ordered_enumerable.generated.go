// This code was generated by a tool.
package enu
import "golang.org/x/exp/constraints"


type OrderedEnumerable[T constraints.Ordered] struct {
	enumerator Enumerator[T]
}

func NewOrdered[T constraints.Ordered](e Enumerator[T]) *OrderedEnumerable[T] {
	return &OrderedEnumerable[T]{enumerator: e}
}

func (e *OrderedEnumerable[T]) Each(iteratee func(T, int)) {
	each[T](e, func(item T, index int) bool {
		iteratee(item, index)
		return true
	})
}

func (e *OrderedEnumerable[T]) ToSlice() []T {
	return ToSlice[T](e)
}

func (e *OrderedEnumerable[T]) Count() int {
	return Count[T](e)
}

func (e *OrderedEnumerable[T]) Filter(predicate func(T, int) bool) *OrderedEnumerable[T] {
	return &OrderedEnumerable[T]{
		enumerator: Filter[T](e, predicate),
	}
}

func (e *OrderedEnumerable[T]) Nth(nth int) (T, bool) {
	return Nth[T](e, nth)
}

func (e *OrderedEnumerable[T]) Find(predicate func(T, int) bool) (T, bool) {
	return Find[T](e, predicate)
}

func (e *OrderedEnumerable[T]) First() (T, bool) {
	return First[T](e)
}

func (e *OrderedEnumerable[T]) Last() (T, bool) {
	return Last[T](e)
}

func (e *OrderedEnumerable[T]) Reverse() *OrderedEnumerable[T] {
	return &OrderedEnumerable[T]{enumerator: Reverse[T](e)}
}

func (e *OrderedEnumerable[T]) SortBy(sorter func(i, j T) bool) *OrderedEnumerable[T] {
	return &OrderedEnumerable[T]{enumerator: SortBy[T](e, sorter)}
}

func (e *OrderedEnumerable[T]) Reject(predicate func(T, int) bool) *OrderedEnumerable[T] {
	return &OrderedEnumerable[T]{enumerator: Reject[T](e, predicate)}
}

func (e *OrderedEnumerable[T]) IsAll(predicate func(T) bool) bool {
	return IsAll[T](e, predicate)
}

func (e *OrderedEnumerable[T]) IsAny(predicate func(T) bool) bool {
	return IsAny[T](e, predicate)
}

func (e *OrderedEnumerable[T]) Take(num uint) *OrderedEnumerable[T] {
	return &OrderedEnumerable[T]{enumerator: Take[T](e, num)}
}

func (e *OrderedEnumerable[T]) GetEnumerator() Enumerator[T] {
	return e.enumerator
}
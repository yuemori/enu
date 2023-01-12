// This code was generated by a tool.
package enu
import "golang.org/x/exp/constraints"


type EnumeratorNumeric[T constraints.Integer | constraints.Float] struct {
	iter   IEnumerable[T]
}

func NewNumeric[T constraints.Integer | constraints.Float](e IEnumerable[T]) *EnumeratorNumeric[T] {
	return &EnumeratorNumeric[T]{iter: e}
}

func (e *EnumeratorNumeric[T]) Each(iteratee func(T, int)) {
	each(e.iter, func(item T, index int) bool {
		iteratee(item, index)
		return true
	})
}

func (e *EnumeratorNumeric[T]) ToSlice() []T {
	return ToSlice(e.iter)
}

func (e *EnumeratorNumeric[T]) Count() int {
	return Count(e.iter)
}

func (e *EnumeratorNumeric[T]) Filter(predicate func(T, int) bool) *EnumeratorNumeric[T] {
	return &EnumeratorNumeric[T]{
		iter: Filter(e.iter, predicate),
	}
}

func (e *EnumeratorNumeric[T]) Nth(nth int) (T, bool) {
	return Nth(e.iter, nth)
}

func (e *EnumeratorNumeric[T]) Find(predicate func(T, int) bool) (T, bool) {
	return Find(e.iter, predicate)
}

func (e *EnumeratorNumeric[T]) First() (T, bool) {
	return First(e.iter)
}

func (e *EnumeratorNumeric[T]) Last() (T, bool) {
	return Last(e.iter)
}

func (e *EnumeratorNumeric[T]) Reverse() *EnumeratorNumeric[T] {
	return &EnumeratorNumeric[T]{iter: Reverse(e.iter)}
}

func (e *EnumeratorNumeric[T]) SortBy(sorter func(i, j T) bool) *EnumeratorNumeric[T] {
	return &EnumeratorNumeric[T]{iter: SortBy(e.iter, sorter)}
}

func (e *EnumeratorNumeric[T]) Reject(predicate func(T, int) bool) *EnumeratorNumeric[T] {
	return &EnumeratorNumeric[T]{iter: Reject(e.iter, predicate)}
}

func (e *EnumeratorNumeric[T]) IsAll(predicate func(T) bool) bool {
	return IsAll(e.iter, predicate)
}

func (e *EnumeratorNumeric[T]) IsAny(predicate func(T) bool) bool {
	return IsAny(e.iter, predicate)
}

func (e *EnumeratorNumeric[T]) Take(num uint) *EnumeratorNumeric[T] {
	return &EnumeratorNumeric[T]{iter: Take(e.iter, num)}
}

func (e *EnumeratorNumeric[T]) GetEnumerator() *Enumerator[T] {
	return &Enumerator[T]{iter: e.iter}
}

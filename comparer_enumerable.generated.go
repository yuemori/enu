// This code was generated by a tool.
package enu

type ComparerEnumerable[T comparable] struct {
	enumerator Enumerator[T]
}

func NewComparer[T comparable](e Enumerator[T]) *ComparerEnumerable[T] {
	return &ComparerEnumerable[T]{enumerator: e}
}

func (e *ComparerEnumerable[T]) Each(iteratee func(T, int)) {
	each[T](e, func(item T, index int) bool {
		iteratee(item, index)
		return true
	})
}

func (e *ComparerEnumerable[T]) ToSlice() []T {
	return ToSlice[T](e)
}

func (e *ComparerEnumerable[T]) Count() int {
	return Count[T](e)
}

func (e *ComparerEnumerable[T]) Filter(predicate func(T, int) bool) *ComparerEnumerable[T] {
	return &ComparerEnumerable[T]{
		enumerator: Filter[T](e, predicate),
	}
}

func (e *ComparerEnumerable[T]) Nth(nth int) (T, bool) {
	return Nth[T](e, nth)
}

func (e *ComparerEnumerable[T]) Find(predicate func(T, int) bool) (T, bool) {
	return Find[T](e, predicate)
}

func (e *ComparerEnumerable[T]) First() (T, bool) {
	return First[T](e)
}

func (e *ComparerEnumerable[T]) Last() (T, bool) {
	return Last[T](e)
}

func (e *ComparerEnumerable[T]) Reverse() *ComparerEnumerable[T] {
	return &ComparerEnumerable[T]{enumerator: Reverse[T](e)}
}

func (e *ComparerEnumerable[T]) SortBy(sorter func(i, j T) bool) *ComparerEnumerable[T] {
	return &ComparerEnumerable[T]{enumerator: SortBy[T](e, sorter)}
}

func (e *ComparerEnumerable[T]) Reject(predicate func(T, int) bool) *ComparerEnumerable[T] {
	return &ComparerEnumerable[T]{enumerator: Reject[T](e, predicate)}
}

func (e *ComparerEnumerable[T]) IsAll(predicate func(T) bool) bool {
	return IsAll[T](e, predicate)
}

func (e *ComparerEnumerable[T]) IsAny(predicate func(T) bool) bool {
	return IsAny[T](e, predicate)
}

func (e *ComparerEnumerable[T]) Take(num uint) *ComparerEnumerable[T] {
	return &ComparerEnumerable[T]{enumerator: Take[T](e, num)}
}

func (e *ComparerEnumerable[T]) GetEnumerator() Enumerator[T] {
	return e.enumerator
}

func (e *ComparerEnumerable[T]) Contains(item T) bool {
	return Contains[T](e, item)
}

func (e *ComparerEnumerable[T]) IndexOf(item T) int {
	return IndexOf[T](e, item)
}

func (e *ComparerEnumerable[T]) ToMap() map[int]T {
	return Reduce[T](e, func(agg map[int]T, item T, index int) map[int]T {
		agg[index] = item
		return agg
	}, map[int]T{})
}

func (e *ComparerEnumerable[T]) Uniq() *ComparerEnumerable[T] {
	return NewComparer[T](Uniq[T](e))
}

// This code was generated by a tool.
package enu

type Enumerable[T any] struct {
	enumerator IEnumerator[T]
}

func New[T any](e IEnumerator[T]) *Enumerable[T] {
	return &Enumerable[T]{enumerator: e}
}

func (e *Enumerable[T]) Each(iteratee func(T, int)) {
	each[T](e, func(item T, index int) bool {
		iteratee(item, index)
		return true
	})
}

func (e *Enumerable[T]) ToSlice() []T {
	return ToSlice[T](e)
}

func (e *Enumerable[T]) Count() int {
	return Count[T](e)
}

func (e *Enumerable[T]) Filter(predicate func(T, int) bool) *Enumerable[T] {
	return &Enumerable[T]{
		enumerator: Filter[T](e, predicate),
	}
}

func (e *Enumerable[T]) Nth(nth int) (T, bool) {
	return Nth[T](e, nth)
}

func (e *Enumerable[T]) Find(predicate func(T, int) bool) (T, bool) {
	return Find[T](e, predicate)
}

func (e *Enumerable[T]) First() (T, bool) {
	return First[T](e)
}

func (e *Enumerable[T]) Last() (T, bool) {
	return Last[T](e)
}

func (e *Enumerable[T]) Reverse() *Enumerable[T] {
	return &Enumerable[T]{enumerator: Reverse[T](e)}
}

func (e *Enumerable[T]) SortBy(sorter func(i, j T) bool) *Enumerable[T] {
	return &Enumerable[T]{enumerator: SortBy[T](e, sorter)}
}

func (e *Enumerable[T]) Reject(predicate func(T, int) bool) *Enumerable[T] {
	return &Enumerable[T]{enumerator: Reject[T](e, predicate)}
}

func (e *Enumerable[T]) IsAll(predicate func(T) bool) bool {
	return IsAll[T](e, predicate)
}

func (e *Enumerable[T]) IsAny(predicate func(T) bool) bool {
	return IsAny[T](e, predicate)
}

func (e *Enumerable[T]) Take(num uint) *Enumerable[T] {
	return &Enumerable[T]{enumerator: Take[T](e, num)}
}

func (e *Enumerable[T]) Result(out *[]T) *Enumerable[T] {
	Result[T](e, out)

	return e
}

func (e *Enumerable[T]) Err() error {
	if p, ok := e.enumerator.(ErrorProvider); ok {
		return p.Err()
	}
	return nil
}

func (e *Enumerable[T]) GetEnumerator() IEnumerator[T] {
	return e.enumerator
}

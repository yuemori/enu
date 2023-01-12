package enu

func FromComparable[T comparable](collection []T) *EnumeratorComparable[T] {
	return NewComparable[T](NewSliceEnumerator(collection))
}

func ToComparable[T comparable](e Enumerator[T]) *EnumeratorComparable[T] {
	return &EnumeratorComparable[T]{
		iter: e.iter,
	}
}

func (e *EnumeratorComparable[T]) ToMap() map[int]T {
	result := map[int]T{}
	e.Each(func(item T, index int) {
		result[index] = item
	})
	return result
}

func (e *EnumeratorComparable[T]) Uniq() *EnumeratorComparable[T] {
	return &EnumeratorComparable[T]{iter: Uniq(e.iter)}
}

func (e *EnumeratorComparable[T]) Contains(item T) bool {
	return Contains(e.iter, item)
}

func (e *EnumeratorComparable[T]) IndexOf(item T) int {
	return IndexOf(e.iter, item)
}

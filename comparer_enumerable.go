package enu

func FromComparable[T comparable](collection []T) *ComparerEnumerable[T] {
	return NewComparer[T](NewSliceEnumerator(collection))
}

func ToComparable[T comparable](e IEnumerable[T]) *ComparerEnumerable[T] {
	return NewComparer(e.GetEnumerator())
}

func (e *ComparerEnumerable[T]) ToMap() map[int]T {
	result := map[int]T{}
	e.Each(func(item T, index int) {
		result[index] = item
	})
	return result
}

func (e *ComparerEnumerable[T]) Uniq() *ComparerEnumerable[T] {
	return NewComparer[T](Uniq[T](e))
}

func (e *ComparerEnumerable[T]) Contains(item T) bool {
	return Contains[T](e, item)
}

func (e *ComparerEnumerable[T]) IndexOf(item T) int {
	return IndexOf[T](e, item)
}

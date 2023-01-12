package enu

func FromComparable[T comparable](collection []T) *ComparerEnumerable[T] {
	return NewComparer[T](NewSliceEnumerator(collection))
}

func ToComparable[T comparable](e IEnumerable[T]) *ComparerEnumerable[T] {
	return NewComparer(e.GetEnumerator())
}

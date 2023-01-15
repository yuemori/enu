package enu

func From[T any](collection []T) *Enumerable[T] {
	return New[T](NewSliceEnumerator(collection))
}

func To[T any](e IEnumerable[T]) *Enumerable[T] {
	return New(e.GetEnumerator())
}

// IEnumerable[T any] is an interface for using Enumerable functions.
type IEnumerable[T any] interface {
	// GetEnumerator returns IEnumerator[T] .
	GetEnumerator() IEnumerator[T]
}

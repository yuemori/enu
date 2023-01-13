package enu

// IEnumerator[T any] supports iteration over a generic collection.
type IEnumerator[T any] interface {
	// Next returns a next item of collection. If Next passes the end of the collection, the empty item and true.
	Next() (T, bool)

	// Dispose disposes the managed resources in the collection.
	// This method called when the iteration is completed.
	Dispose()
}

package enu

// IEnumerator[T any] supports iteration over a generic collection.
type IEnumerator[T any] interface {
	// Next returns a next item of collection. If Next passes the end of the collection, the empty item and true.
	Next() (T, bool)

	// Dispose disposes the managed resources in the collection.
	// This method called when the iteration is completed.
	Dispose()
}

// ErrorProvider supports iteration error. If an Enumerable raises an error during execution of Next() or Dispose(), implement this interface.
type ErrorProvider interface {
	// Err returns error during execution of Next() or Dispose()
	Err() error
}

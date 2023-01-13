package enu

// IEnumerator[T any] supports iteration over a generic collection.
type IEnumerator[T any] interface {
	// Next returns a next item of collection. If Next passes the end of the collection, the empty item and true.
	Next() (T, bool)

	// Reset resets the collection. If the collection does not support Reset, you don't have to do anything.
	// This method is similar to called on a non-lazy Enumerable functions is called:
	//
	// ```
	// e := enu.From([]int{1, 2, 3})
	//
	// // Filter is lazy function, Reset does not called.
	// e2 := e.Filter(func(item, _ int) bool { return item % 2 == 0 })
	//
	// // Take is lazy function, Reset does not called.
	// e3 := e2.Take(2)
	//
	// // ToSlice is non-lazy function, Reset called.
	// r := e3.ToSlice()
	// ```
	//
	Reset()
	Dispose()
}

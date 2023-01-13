package enu

type FilterEnumerable[T any] struct {
	iter      IEnumerator[T]
	predicate func(item T, index int) bool
	index     int
}

func (e *FilterEnumerable[T]) GetEnumerator() IEnumerator[T] {
	return e
}

func (e *FilterEnumerable[T]) Reset() {
	e.iter.Reset()
	e.index = 0
}

func (e *FilterEnumerable[T]) Stop() {
	e.iter.Stop()
	e.index = 0
}

func (e *FilterEnumerable[T]) Next() (T, bool) {
	for {
		item, ok := e.iter.Next()
		if !ok {
			return empty[T](), false
		}
		if e.predicate(item, e.index) {
			return item, true
		}
		e.index++
	}
}

type RejectEnumerable[T any] struct {
	iter      IEnumerator[T]
	predicate func(item T, index int) bool
	index     int
}

func (e *RejectEnumerable[T]) GetEnumerator() IEnumerator[T] {
	return e
}

func (e *RejectEnumerable[T]) Reset() {
	e.iter.Reset()
	e.index = 0
}

func (e *RejectEnumerable[T]) Stop() {
	e.iter.Stop()
	e.index = 0
}

func (e *RejectEnumerable[T]) Next() (T, bool) {
	for {
		item, ok := e.iter.Next()
		if !ok {
			return empty[T](), false
		}
		if !e.predicate(item, e.index) {
			return item, true
		}
		e.index++
	}
}

type TakeEnumerable[T any] struct {
	iter  IEnumerator[T]
	size  uint
	index uint
}

func (e *TakeEnumerable[T]) GetEnumerator() IEnumerator[T] {
	return e
}

func (e *TakeEnumerable[T]) Reset() {
	e.iter.Reset()
	e.index = 0
}

func (e *TakeEnumerable[T]) Stop() {
	e.iter.Stop()
	e.index = 0
}

func (e *TakeEnumerable[T]) Next() (T, bool) {
	if e.size == e.index {
		e.Stop()
		return empty[T](), false
	}
	item, ok := e.iter.Next()
	if !ok {
		e.Stop()
		return empty[T](), false
	}
	e.index++
	return item, true
}

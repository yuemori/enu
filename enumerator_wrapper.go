package enu

type FilterEnumerator[T any] struct {
	iter      IEnumerable[T]
	predicate func(item T, index int) bool
	index     int
}

func (e *FilterEnumerator[T]) Reset() {
	e.iter.Reset()
	e.index = 0
}

func (e *FilterEnumerator[T]) Stop() {
	e.iter.Stop()
	e.index = 0
}

func (e *FilterEnumerator[T]) Next() (T, bool) {
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

type RejectEnumerator[T any] struct {
	iter      IEnumerable[T]
	predicate func(item T, index int) bool
	index     int
}

func (e *RejectEnumerator[T]) Reset() {
	e.iter.Reset()
	e.index = 0
}

func (e *RejectEnumerator[T]) Stop() {
	e.iter.Stop()
	e.index = 0
}

func (e *RejectEnumerator[T]) Next() (T, bool) {
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

type TakeEnumerator[T any] struct {
	iter  IEnumerable[T]
	size  uint
	index uint
}

func (e *TakeEnumerator[T]) Reset() {
	e.iter.Reset()
	e.index = 0
}

func (e *TakeEnumerator[T]) Stop() {
	e.iter.Stop()
	e.index = 0
}

func (e *TakeEnumerator[T]) Next() (T, bool) {
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

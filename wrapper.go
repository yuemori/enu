package enu

type FilterEnumerable[T any] struct {
	iter      IEnumerator[T]
	predicate func(item T, index int) bool
	index     int
}

func (e *FilterEnumerable[T]) GetEnumerator() IEnumerator[T] {
	return e
}

func (e *FilterEnumerable[T]) Dispose() {
	e.iter.Dispose()
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

func (e *RejectEnumerable[T]) Dispose() {
	e.iter.Dispose()
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

func (e *TakeEnumerable[T]) Dispose() {
	e.iter.Dispose()
	e.index = 0
}

func (e *TakeEnumerable[T]) Next() (T, bool) {
	if e.size == e.index {
		return empty[T](), false
	}
	item, ok := e.iter.Next()
	if !ok {
		return empty[T](), false
	}
	e.index++
	return item, true
}

type UniqEnumerable[T comparable] struct {
	iter IEnumerator[T]
	seen map[T]struct{}
}

func (e *UniqEnumerable[T]) GetEnumerator() IEnumerator[T] {
	return e
}

func (e *UniqEnumerable[T]) Dispose() {
	e.iter.Dispose()
	e.seen = nil
}

func (e *UniqEnumerable[T]) Next() (T, bool) {
	if e.seen == nil {
		e.seen = map[T]struct{}{}
	}
	for {
		item, ok := e.iter.Next()
		if !ok {
			return empty[T](), false
		}
		if _, ok := e.seen[item]; ok {
			continue
		}
		e.seen[item] = struct{}{}
		return item, true
	}
}

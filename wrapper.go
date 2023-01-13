package enu

type FilterEnumerator[T any] struct {
	iter      IEnumerator[T]
	predicate func(item T, index int) bool
	index     int
}

func (e *FilterEnumerator[T]) GetEnumerator() IEnumerator[T] {
	return e
}

func (e *FilterEnumerator[T]) Dispose() {
	e.iter.Dispose()
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
	iter      IEnumerator[T]
	predicate func(item T, index int) bool
	index     int
}

func (e *RejectEnumerator[T]) GetEnumerator() IEnumerator[T] {
	return e
}

func (e *RejectEnumerator[T]) Dispose() {
	e.iter.Dispose()
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
	iter  IEnumerator[T]
	size  uint
	index uint
}

func (e *TakeEnumerator[T]) GetEnumerator() IEnumerator[T] {
	return e
}

func (e *TakeEnumerator[T]) Dispose() {
	e.iter.Dispose()
	e.index = 0
}

func (e *TakeEnumerator[T]) Next() (T, bool) {
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

type UniqEnumerator[T comparable] struct {
	iter IEnumerator[T]
	seen map[T]struct{}
}

func (e *UniqEnumerator[T]) GetEnumerator() IEnumerator[T] {
	return e
}

func (e *UniqEnumerator[T]) Dispose() {
	e.iter.Dispose()
	e.seen = nil
}

func (e *UniqEnumerator[T]) Next() (T, bool) {
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

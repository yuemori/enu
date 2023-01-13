package enu

type RangeValuer[T1, T2 any] interface {
	Compare(T1) int
	Value() T1
	Next(step T2) RangeValuer[T1, T2]
}

func FromRange[T1, T2 any](min, max RangeValuer[T1, T2], step T2) *Enumerable[T1] {
	return New[T1](NewRange(min, max, step))
}

func NewRange[T1, T2 any](min, max RangeValuer[T1, T2], step T2) *RangeEnumerator[T1, T2] {
	return &RangeEnumerator[T1, T2]{
		min:  min,
		max:  max,
		step: step,
	}
}

type RangeEnumerator[T1, T2 any] struct {
	min     RangeValuer[T1, T2]
	max     RangeValuer[T1, T2]
	current RangeValuer[T1, T2]
	step    T2
}

func empty[T any]() T {
	var empty T
	return empty
}

func (e *RangeEnumerator[T1, T2]) Dispose() {
	e.current = nil
}

func (e *RangeEnumerator[T1, T2]) Next() (T1, bool) {
	if e.min.Compare(e.max.Value()) == 1 {
		return empty[T1](), false
	}
	if e.current == nil {
		val := e.min.Value()
		e.current = e.min.Next(e.step)
		return val, true
	}
	if e.current.Compare(e.max.Value()) == 1 {
		return empty[T1](), false
	}
	val := e.current.Value()
	e.current = e.current.Next(e.step)
	return val, true
}

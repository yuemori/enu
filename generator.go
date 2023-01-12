package enu

func FromFunc[T any](generator func(index int) (T, bool)) *Enumerable[T] {
	return New[T](NewGenerator(generator))
}

func NewGenerator[T any](generator func(index int) (T, bool)) *FuncEnumerator[T] {
	return &FuncEnumerator[T]{fn: generator}
}

type FuncEnumerator[T any] struct {
	fn    func(int) (T, bool)
	index int
}

func (g *FuncEnumerator[T]) Next() (T, bool) {
	item, ok := g.fn(g.index)
	if !ok {
		return empty[T](), false
	}
	g.index += 1
	return item, true
}

func (g *FuncEnumerator[T]) Stop() {}

func (g *FuncEnumerator[T]) Reset() {
	g.index = 0
}

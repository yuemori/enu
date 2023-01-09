package enu

func FromFunc[T any](generator func(index int) (T, bool)) *Enumerator[T] {
	return &Enumerator[T]{iter: NewGenerator(generator)}
}

func NewGenerator[T any](generator func(index int) (T, bool)) *Generator[T] {
	return &Generator[T]{generator: generator}
}

type Generator[T any] struct {
	generator func(int) (T, bool)
	generated []T
	index     int
}

func (g *Generator[T]) Next() (T, bool) {
	if len(g.generated)-1 < g.index {
		item, ok := g.generator(g.index)
		if !ok {
			return empty[T](), false
		}
		g.generated = append(g.generated, item)
		g.index += 1
		return item, true
	}

	item := g.generated[g.index]
	g.index++
	return item, true
}

func (g *Generator[T]) Stop() {}

func (g *Generator[T]) Reset() {
	g.index = 0
}

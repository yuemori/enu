package enu

func FromFunc[T any](generator func(index int) (T, bool)) *Enumerator[T] {
	return &Enumerator[T]{iter: NewGenerator(generator)}
}

func NewGenerator[T any](generator func(index int) (T, bool)) *Generator[T] {
	return &Generator[T]{generator: generator}
}

type Generator[T any] struct {
	generator func(int) (T, bool)
	index     int
}

func (s *Generator[T]) Next() (T, bool) {
	item, ok := s.generator(s.index)
	if !ok {
		return empty[T](), false
	}
	s.index += 1
	return item, true
}

func (s *Generator[T]) Stop() {}

func (s *Generator[T]) Reset() {
	s.index = 0
}

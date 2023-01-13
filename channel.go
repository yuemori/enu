package enu

func FromChannel[T any](ch chan (T)) *Enumerable[T] {
	return New[T](NewChannelEnumerator(ch))
}

type ChannelEnumerator[T any] struct {
	sender chan (T)
}

func NewChannelEnumerator[T any](ch chan (T)) *ChannelEnumerator[T] {
	return &ChannelEnumerator[T]{
		sender: ch,
	}
}

func (e *ChannelEnumerator[T]) Dispose() {
}

func (e *ChannelEnumerator[T]) Next() (T, bool) {
	v, ok := <-e.sender
	if !ok {
		return empty[T](), false
	}

	return v, true
}

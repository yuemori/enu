package enu

func FromChannel[T any](ch chan (T)) *Enumerator[T] {
	return &Enumerator[T]{iter: NewChannelEnumerator(ch)}
}

type ChannelEnumerator[T any] struct {
	sender chan (T)
}

func NewChannelEnumerator[T any](ch chan (T)) *ChannelEnumerator[T] {
	return &ChannelEnumerator[T]{
		sender: ch,
	}
}

func (e *ChannelEnumerator[T]) Reset() {
}

func (e *ChannelEnumerator[T]) Stop() {
}

func (e *ChannelEnumerator[T]) Next() (T, bool) {
	for {
		select {
		case v, ok := <-e.sender:
			if !ok {
				return empty[T](), false
			}

			return v, true
		default:
		}
	}
}

package enu

func FromChannel[T any](ch chan (T)) *Enumerator[T] {
	return &Enumerator[T]{iter: &ChannelEnumerator[T]{sender: ch}}
}

func FromChannelWithDone[T any](ch chan (T), done chan (struct{})) *Enumerator[T] {
	return &Enumerator[T]{iter: &ChannelEnumerator[T]{sender: ch, doneCh: done}}
}

type ChannelEnumerator[T any] struct {
	sender chan (T)
	done   bool
	doneCh chan (struct{})
	result []T
	index  int
}

func NewChannelEnumerator[T any](ch chan (T), done chan (struct{})) *ChannelEnumerator[T] {
	return &ChannelEnumerator[T]{sender: ch, doneCh: done}
}

func (e *ChannelEnumerator[T]) Stop() {
	e.done = true
	go func() {
		e.doneCh <- struct{}{}
	}()
}

func (e *ChannelEnumerator[T]) Done() chan (struct{}) {
	return e.doneCh
}

func (e *ChannelEnumerator[T]) Reset() {
	e.index = 0
}

func (e *ChannelEnumerator[T]) Next() (T, bool) {
	if len(e.result)-1 < e.index {
		for {
			select {
			case v, ok := <-e.sender:
				if !ok {
					return empty[T](), false
				}
				if e.done {
					return empty[T](), false
				}
				e.result = append(e.result, v)
				e.index += 1

				return v, true
			default:
				if e.done {
					return empty[T](), false
				}
			}
		}
	}

	item := e.result[e.index]
	e.index++
	return item, true
}

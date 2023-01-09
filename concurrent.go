package enu

import (
	"sync"
)

func Repeat[T any](bufferSize, repeatCount int, fn func(int) T) chan (T) {
	ch := make(chan (T), bufferSize)
	wg := new(sync.WaitGroup)

	for i := 0; i < repeatCount; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			ch <- fn(j)
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func Parallel[T any](bufferSize int, funcs ...func() T) chan (T) {
	ch := make(chan (T), bufferSize)
	wg := new(sync.WaitGroup)

	for _, fn := range funcs {
		wg.Add(1)
		go func(f func() T) {
			defer wg.Done()
			ch <- f()
		}(fn)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func Retry[T any](bufferSize, maxRetry int, fn func(int) (T, error)) chan (Tuple2[T, bool]) {
	ch := make(chan (Tuple2[T, bool]), bufferSize)

	go func() {
		defer close(ch)
		retry := 0
		for {
			v, err := fn(retry)
			if err == nil {
				ch <- Tuple2[T, bool]{v, true}
				break
			}
			if maxRetry < retry {
				ch <- Tuple2[T, bool]{empty[T](), false}
				break
			}
			retry++
		}
	}()

	return ch
}

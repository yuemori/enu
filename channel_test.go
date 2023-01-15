package enu_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enu"
)

func isChannelClosed[T any](ch chan (T)) bool {
	select {
	case _, ok := <-ch:
		if ok {
			return false
		}
		return true
	default:
		return false
	}
}

func TestChannel(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ch := make(chan (int), 1)

	go func() {
		defer close(ch)

		for i := 1; i < 6; i++ {
			ch <- i
		}
	}()

	r := enu.FromChannel(ch).ToSlice()
	is.Equal([]int{1, 2, 3, 4, 5}, r)
	is.True(isChannelClosed(ch))
}

func TestChannelWithFirst(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ch := make(chan (int))

	go func() {
		defer close(ch)

		for i := 1; i < 6; i++ {
			ch <- i
		}
	}()

	r, ok := enu.FromChannel(ch).First()
	is.True(ok)
	is.Equal(1, r)
	// Important: func 6 times to send channel, enumerator receive value from channel at 1 times because First() called.
	is.False(isChannelClosed(ch))
}

func TestChannelWithTake(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ch := make(chan (int))

	go func() {
		defer close(ch)

		for i := 1; i < 6; i++ {
			ch <- i
		}
	}()

	stream := enu.FromChannel(ch)
	r1 := stream.Take(3).ToSlice()
	is.Equal([]int{1, 2, 3}, r1)

	r2 := stream.Take(1).ToSlice()
	is.Equal([]int{4}, r2)
	is.False(isChannelClosed(ch))
}

func TestChannelWithFind(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ch := make(chan (int), 1)

	go func() {
		defer close(ch)

		for i := 1; i < 6; i++ {
			ch <- i
		}
	}()

	channel := enu.FromChannel(ch)

	r1, ok := channel.Find(func(i, _ int) bool {
		return i == 3
	})
	is.Equal(true, ok)
	is.Equal(3, r1)
	// Important: func 6 times to send channel, enumerator 3 times received from channel because Find(3) called.
	is.False(isChannelClosed(ch))
}

func TestChannelWithClose(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ch := make(chan (int), 1)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3

		close(ch)
	}()

	r1 := enu.FromChannel(ch).ToSlice()
	is.Equal([]int{1, 2, 3}, r1)
	is.True(isChannelClosed(ch))
}

//	func TestChannelWithDone(t *testing.T) {
//		t.Parallel()
//		is := assert.New(t)
//
//		ch := make(chan (int))
//		done := make(chan struct{})
//
//		go func() {
//			defer close(ch)
//			defer close(done)
//
//			i := 0
//			for {
//				select {
//				case <-done:
//					break
//				default:
//					ch <- i
//					i++
//				}
//			}
//		}()
//
//		channel := enu.FromChannelWithDone(ch, done)
//		r1, ok := channel.First()
//		is.Equal(true, ok)
//		is.Equal(0, r1)
//
//		r2 := channel.Take(10).ToSlice()
//		is.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, r2)
//
//		<-done
//
//		r3 := channel.Take(10).ToSlice()
//		is.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, r3)
//	}
//
//	func TestChannelWithDoneWithFind(t *testing.T) {
//		t.Parallel()
//		is := assert.New(t)
//
//		ch := make(chan (int))
//		done := make(chan struct{})
//
//		go func() {
//			defer close(ch)
//			defer close(done)
//
//			i := 0
//			for {
//				select {
//				case <-done:
//					break
//				default:
//					ch <- i
//					i++
//				}
//			}
//		}()
//
//		channel := enu.FromChannelWithDone(ch, done)
//		r1, ok := channel.Find(func(item, _ int) bool {
//			return item == 5
//		})
//		is.Equal(true, ok)
//		is.Equal(5, r1)
//
//		r2 := channel.Take(10).ToSlice()
//		is.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, r2)
//	}
func TestChannelWithRetry(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	mockHttpRequest := func() func(int) (string, error) {
		return func(n int) (string, error) {
			if n < 5 {
				return "", errors.New("some error")
			}
			return "OK", nil
		}
	}

	// Returns [result[T], nil] if function was succeed.
	r1, ok := enu.FromChannel(enu.Retry(3, 5, mockHttpRequest())).First()
	is.Equal(true, ok)
	is.Equal("OK", r1.A)
	is.NoError(r1.B)

	// Returns [empty[T], error] if function was failed with reach to maxRetry.
	r2, ok := enu.FromChannel(enu.Retry(3, 3, mockHttpRequest())).First()
	is.Equal(true, ok)
	is.Errorf(r2.B, "some error")
}

func TestChannelWithRepeat(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	mockHttpRequest := func() func(int) enu.Tuple2[string, error] {
		return func(n int) enu.Tuple2[string, error] {
			if n%2 == 0 {
				return enu.Tuple2[string, error]{"OK", nil}
			}
			return enu.Tuple2[string, error]{"", errors.New("some error")}
		}
	}

	// Aggregate function result by repeat and async
	r := enu.FromChannel(enu.Repeat(2, 5, mockHttpRequest())).ToSlice()

	is.ElementsMatch([]enu.Tuple2[string, error]{
		{"OK", nil},
		{"OK", nil},
		{"OK", nil},
		{"", errors.New("some error")},
		{"", errors.New("some error")},
	}, r)
}

func TestChannelWithParallel(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	request1 := func() enu.Tuple2[string, error] {
		time.Sleep(100 * time.Millisecond)
		return enu.Tuple2[string, error]{"https://google.com", nil}
	}

	request2 := func() enu.Tuple2[string, error] {
		time.Sleep(100 * time.Millisecond)
		return enu.Tuple2[string, error]{"https://amazon.co.jp", nil}
	}

	request3 := func() enu.Tuple2[string, error] {
		time.Sleep(100 * time.Millisecond)
		return enu.Tuple2[string, error]{"", errors.New("some error")}
	}

	now := time.Now()

	// Aggregate multiple function result by async
	r := enu.FromChannel(enu.Parallel(3, request1, request2, request3)).ToSlice()
	is.True(time.Since(now) < (200 * time.Millisecond))

	is.ElementsMatch([]enu.Tuple2[string, error]{
		{"", errors.New("some error")},
		{"https://amazon.co.jp", nil},
		{"https://google.com", nil},
	}, r)
}

func TestChannelWithAsync(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enu.FromChannel(enu.Async(func(ch chan (int)) {
		ch <- 1
		ch <- 2
		ch <- 3
	})).ToSlice()
	is.Equal([]int{1, 2, 3}, r1)
}

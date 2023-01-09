package enu_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enu"
)

func TestChannel(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ch := make(chan (int), 3)

	go func() {
		defer close(ch)

		for i := 1; i < 6; i++ {
			ch <- i
		}
	}()

	channel := enu.FromChannel(ch)
	r1, ok := channel.First()
	is.Equal(true, ok)
	is.Equal(1, r1)

	r2 := channel.ToSlice()
	is.Equal([]int{1, 2, 3, 4, 5}, r2)
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

	r := enu.FromChannel(ch).Take(10).ToSlice()
	is.Equal([]int{1, 2, 3, 4, 5}, r)
}

func TestChannelWithFind(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ch := make(chan (int))

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

	r2 := channel.ToSlice()
	is.Equal([]int{1, 2, 3, 4, 5}, r2)
}

func TestChannelWithDone(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ch := make(chan (int))
	done := make(chan struct{})

	go func() {
		defer close(ch)
		defer close(done)

		i := 0
		for {
			select {
			case <-done:
				break
			default:
				ch <- i
				i++
			}
		}
	}()

	channel := enu.FromChannelWithDone(ch, done)
	r1, ok := channel.First()
	is.Equal(true, ok)
	is.Equal(0, r1)

	r2 := channel.Take(10).ToSlice()
	is.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, r2)

	<-done

	r3 := channel.Take(10).ToSlice()
	is.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, r3)
}

func TestChannelWithDoneWithFind(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ch := make(chan (int))
	done := make(chan struct{})

	go func() {
		defer close(ch)
		defer close(done)

		i := 0
		for {
			select {
			case <-done:
				break
			default:
				ch <- i
				i++
			}
		}
	}()

	channel := enu.FromChannelWithDone(ch, done)
	r1, ok := channel.Find(func(item, _ int) bool {
		return item == 5
	})
	is.Equal(true, ok)
	is.Equal(5, r1)

	r2 := channel.Take(10).ToSlice()
	is.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, r2)
}

func TestChannelWithRetry(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	mockHttpRequest := func() func(int) (string, error) {
		return func(n int) (string, error) {
			if n < 5 {
				n += 1
				return "", errors.New("some error")
			}
			return "OK", nil
		}
	}

	// Returns [result[T], true] if function was succeed.
	r1, ok := enu.FromChannel(enu.Retry(3, 5, mockHttpRequest())).First()
	is.Equal(true, ok)
	is.Equal("OK", r1.A)
	is.Equal(true, r1.B)

	// Returns [empty[T], false] if function was failed with reach to maxRetry.
	r2, ok := enu.FromChannel(enu.Retry(3, 3, mockHttpRequest())).First()
	is.Equal(true, ok)
	is.Equal(false, r2.B)
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
	r := enu.FromChannel(enu.Repeat(2, 5, mockHttpRequest())).SortBy(func(i, j enu.Tuple2[string, error]) bool {
		return j.B != nil
	}).ToSlice()

	is.Equal([]enu.Tuple2[string, error]{
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
		return enu.Tuple2[string, error]{"https://google.com", nil}
	}

	request2 := func() enu.Tuple2[string, error] {
		return enu.Tuple2[string, error]{"https://amazon.co.jp", nil}
	}

	request3 := func() enu.Tuple2[string, error] {
		return enu.Tuple2[string, error]{"", errors.New("some error")}
	}

	// Aggregate multiple function result by async
	r := enu.FromChannel(enu.Parallel(2, request1, request2, request3)).SortBy(func(i, j enu.Tuple2[string, error]) bool {
		return strings.Compare(i.A, j.A) == -1
	}).ToSlice()

	is.Equal([]enu.Tuple2[string, error]{
		{"", errors.New("some error")},
		{"https://amazon.co.jp", nil},
		{"https://google.com", nil},
	}, r)
}

package enu_test

import (
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

	r1, ok := channel.Find(func(i int) bool {
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
	r1, ok := channel.Find(func(item int) bool {
		return item == 5
	})
	is.Equal(true, ok)
	is.Equal(5, r1)

	r2 := channel.Take(10).ToSlice()
	is.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, r2)
}

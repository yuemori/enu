package enu_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enu"
)

func TestGeneratorFunc(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	fibonacci := func() func(int) (int, bool) {
		n, n1 := 0, 1
		return func(i int) (int, bool) {
			fn := func(index int) (int, bool) {
				// Reset variable if index given zero.
				if index == 0 {
					n, n1 = 0, 1
				}
				v := n
				n, n1 = n1, n+n1
				return v, true
			}
			return fn(i)
		}
	}

	g := enu.FromFunc(fibonacci())

	r1, ok := g.First()
	is.Equal(true, ok)
	is.Equal(0, r1)

	r2 := g.Take(10).ToSlice()
	is.Equal([]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, r2)

	r3, ok := g.First()
	is.Equal(true, ok)
	is.Equal(0, r3)

	r4, ok := g.Last()
	is.Equal(true, ok)
	is.Equal(34, r4)
}

func TestGeneratorLimitedFunc(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.FromFunc(func(index int) (int, bool) {
		if index < 10 {
			return index * 2, true
		}
		// Stop enumerator if returns false.
		return 0, false
	}).ToSlice()

	is.Equal([]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}, r)
}

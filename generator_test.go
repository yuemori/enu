package enu_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enu"
)

func TestGeneratorFunc(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	n, n1 := 0, 1
	g := enu.FromFunc(func(index int) (int, bool) {
		if index == 0 {
			n, n1 = 0, 1
		}
		fn := func() int {
			v := n
			n, n1 = n1, n+n1
			return v
		}
		return fn(), true
	})

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

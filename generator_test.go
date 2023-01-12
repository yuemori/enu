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
			v := n
			n, n1 = n1, n+n1
			return v, true
		}
	}

	r1 := enu.FromFunc(fibonacci()).Take(10).ToSlice()
	is.Equal([]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, r1)

	g := enu.FromFunc(fibonacci())
	r2, ok := g.First()
	is.Equal(true, ok)
	is.Equal(0, r2)

	r3 := g.Take(10).ToSlice()
	is.Equal([]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}, r3)

	r4, ok := g.First()
	is.Equal(true, ok)
	is.Equal(89, r4)

	r5 := enu.FromFunc(fibonacci()).Take(5).Filter(func(item int, _ int) bool {
		return item%2 == 0
	}).ToSlice()
	is.Equal([]int{0, 2}, r5)

	r6 := enu.FromFunc(fibonacci()).Filter(func(item int, _ int) bool {
		return item%2 == 0
	}).Take(5).ToSlice()
	is.Equal([]int{0, 2, 8, 34, 144}, r6)
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

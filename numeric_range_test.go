package enu_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enu"
)

func TestNumericRange(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enu.NewNumericRange(1, 5).ToSlice()
	is.Equal([]int{1, 2, 3, 4, 5}, r1)

	r2 := enu.NewNumericRange(1, 10).ToSlice()
	is.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, r2)

	r3 := enu.NewNumericRange(-5, -1).ToSlice()
	is.Equal([]int{-5, -4, -3, -2, -1}, r3)

	r4 := enu.NewNumericRange(-1, -5).ToSlice()
	is.Equal([]int{}, r4)

	r5 := enu.NewNumericRange(5, 1).ToSlice()
	is.Equal([]int{}, r5)

	r6 := enu.NewNumericRange(1, math.MaxInt).Take(10).ToSlice()
	is.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, r6)

	r7 := enu.NewNumericRange(1, math.Inf(0)).Take(3).ToSlice()
	is.Equal([]float64{1, 2, 3}, r7)
}

func TestNumericRangeWithFind(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r, ok := enu.NewNumericRange(1, math.MaxInt).Find(func(i int) bool {
		return i > 1000
	})
	is.Equal(true, ok)
	is.Equal(1001, r)
}

func TestNumericRangeWithStep(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enu.NewNumericRangeWithStep(1, 5, 1).ToSlice()
	is.Equal([]int{1, 2, 3, 4, 5}, r1)

	r2 := enu.NewNumericRangeWithStep(1, 5, 2).ToSlice()
	is.Equal([]int{1, 3, 5}, r2)

	r3 := enu.NewNumericRangeWithStep(1, 5, 3).ToSlice()
	is.Equal([]int{1, 4}, r3)

	r4 := enu.NewNumericRangeWithStep(-5, -1, 1).ToSlice()
	is.Equal([]int{-5, -4, -3, -2, -1}, r4)

	r5 := enu.NewNumericRangeWithStep(-5, -1, 2).ToSlice()
	is.Equal([]int{-5, -3, -1}, r5)

	r6 := enu.NewNumericRangeWithStep(1.1, 5.1, 1.0).ToSlice()
	is.Equal([]float64{1.1, 2.1, 3.1, 4.1, 5.1}, r6)

	r7 := enu.NewNumericRangeWithStep(1.1, math.Inf(0), 2.0).Take(5).ToSlice()
	is.Equal([]float64{1.1, 3.1, 5.1, 7.1, 9.1}, r7)
}

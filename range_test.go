package enumerator_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enumerator"
)

func TestRange(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enumerator.NewRange(1, 5).ToSlice()
	is.Equal([]int{1, 2, 3, 4, 5}, r1)

	r2 := enumerator.NewRange(1, 10).ToSlice()
	is.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, r2)

	r3 := enumerator.NewRange(-5, -1).ToSlice()
	is.Equal([]int{-5, -4, -3, -2, -1}, r3)

	r4 := enumerator.NewRange(-1, -5).ToSlice()
	is.Equal([]int{}, r4)

	r5 := enumerator.NewRange(5, 1).ToSlice()
	is.Equal([]int{}, r5)

	r6 := enumerator.NewRange(1, math.MaxInt).Take(10).ToSlice()
	is.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, r6)

	r7 := enumerator.NewRange(1, math.Inf(0)).Take(3).ToSlice()
	is.Equal([]float64{1, 2, 3}, r7)
}

func TestRangeWithStep(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enumerator.NewRangeWithStep(1, 5, 1).ToSlice()
	is.Equal([]int{1, 2, 3, 4, 5}, r1)

	r2 := enumerator.NewRangeWithStep(1, 5, 2).ToSlice()
	is.Equal([]int{1, 3, 5}, r2)

	r3 := enumerator.NewRangeWithStep(1, 5, 3).ToSlice()
	is.Equal([]int{1, 4}, r3)

	r4 := enumerator.NewRangeWithStep(-5, -1, 1).ToSlice()
	is.Equal([]int{-5, -4, -3, -2, -1}, r4)

	r5 := enumerator.NewRangeWithStep(-5, -1, 2).ToSlice()
	is.Equal([]int{-5, -3, -1}, r5)

	r6 := enumerator.NewRangeWithStep(1.1, 5.1, 1.0).ToSlice()
	is.Equal([]float64{1.1, 2.1, 3.1, 4.1, 5.1}, r6)

	r7 := enumerator.NewRangeWithStep(1.1, math.Inf(0), 2.0).Take(5).ToSlice()
	is.Equal([]float64{1.1, 3.1, 5.1, 7.1, 9.1}, r7)
}

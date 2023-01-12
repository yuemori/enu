package enu_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enu"
)

func TestMin(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enu.FromOrdered([]int{2, 1, 3}).Min()
	is.Equal(1, r1)

	r2 := enu.FromOrdered([]int{}).Min()
	is.Equal(0, r2)
}

func TestMax(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enu.FromOrdered([]int{2, 1, 3}).Max()
	is.Equal(3, r1)

	r2 := enu.FromOrdered([]int{}).Max()
	is.Equal(0, r2)
}

func TestSort(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.FromOrdered([]string{"c", "a", "b"}).Sort().ToSlice()

	is.Equal([]string{"a", "b", "c"}, r)
}

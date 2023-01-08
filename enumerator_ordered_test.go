package enu_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enu"
)

func TestMin(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.FromOrdered([]int{2, 1, 3}).Min()

	is.Equal(1, r)
}

func TestMax(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.FromOrdered([]int{2, 1, 3}).Max()

	is.Equal(3, r)
}

func TestSort(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.FromOrdered([]int{2, 1, 3}).Sort().ToSlice()

	is.Equal([]int{1, 2, 3}, r)
}

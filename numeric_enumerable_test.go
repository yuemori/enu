package enu_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enu"
)

func TestSum(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enu.FromNumeric([]int{2, 1, 3}).Sum()
	is.Equal(6, r1)

	r2 := enu.FromNumeric([]int{}).Sum()
	is.Equal(0, r2)
}

package enu_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enu"
)

func TestMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.Map[int](enu.From([]int{1, 2, 3}), func(e, i int) string {
		return strconv.Itoa(e)
	}).ToSlice()

	is.Equal([]string{"1", "2", "3"}, r)
}

package enumerator_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	enu "github.com/yuemori/enumerator"
)

func TestMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.Map([]int{1, 2, 3}, func(e, i int) string {
		return strconv.Itoa(e)
	}).ToSlice()

	is.Equal([]string{"1", "2", "3"}, r)
}

func TestMapE(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.MapE(enu.From([]int{1, 2, 3}), func(e, i int) string {
		return strconv.Itoa(e)
	}).ToSlice()

	is.Equal([]string{"1", "2", "3"}, r)
}

func TestMapC(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.MapC(enu.FromComparable([]int{1, 2, 2, 3, 3, 3}).Uniq(), func(e, i int) string {
		return strconv.Itoa(e)
	}).ToSlice()

	is.Equal([]string{"1", "2", "3"}, r)
}

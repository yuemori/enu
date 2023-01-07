package enumerator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enumerator"
)

func TestUniq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enumerator.FromSliceC([]int{1, 1, 2, 3, 3}).Uniq().ToSlice()

	is.Equal([]int{1, 2, 3}, r)
}

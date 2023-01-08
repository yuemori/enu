package enumerator_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enumerator"
)

func TestCount(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enumerator.From([]int{1, 2, 3, 4, 5}).Count()

	is.Equal(5, r)
}

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enumerator.From([]int{1, 2, 3, 4, 5}).Filter(func(i int, _ int) bool {
		return i%2 == 0
	}).ToSlice()

	is.Equal([]int{2, 4}, r)
}

func TestReject(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enumerator.From([]int{1, 2, 3, 4, 5}).Reject(func(i int, _ int) bool {
		return i%2 == 0
	}).ToSlice()

	is.Equal([]int{1, 3, 5}, r)
}

type errorE[T any] struct{}

func (e errorE[T]) Stop() {}
func (e errorE[T]) Next() (T, error) {
	var empty T
	return empty, errors.New("error")
}

func TestFirst(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1, ok := enumerator.From([]int{1, 2, 3, 4, 5}).First()
	is.Equal(true, ok)
	is.Equal(1, r1)

	r2, ok := enumerator.From([]int{}).First()
	is.Equal(false, ok)
	is.Equal(0, r2)

	type dummy struct{}

	r3, ok := enumerator.From([]dummy{}).First()
	is.Equal(false, ok)
	is.Equal(dummy{}, r3)

	r4, ok := enumerator.From([]*dummy{}).First()
	is.Equal(false, ok)
	is.Nil(r4)

	enum := enumerator.New[int](errorE[int]{})
	r5, ok := enum.First()
	is.Equal(false, ok)
	is.Equal(0, r5)
	is.Error(enum.Error())
}

func TestLast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1, ok := enumerator.From([]int{1, 2, 3, 4, 5}).Last()
	is.Equal(true, ok)
	is.Equal(5, r1)

	r2, ok := enumerator.From([]int{}).Last()
	is.Equal(false, ok)
	is.Equal(0, r2)

	type dummy struct{}

	r3, ok := enumerator.From([]dummy{}).Last()
	is.Equal(false, ok)
	is.Equal(dummy{}, r3)

	r4, ok := enumerator.From([]*dummy{}).Last()
	is.Equal(false, ok)
	is.Nil(r4)

	enum := enumerator.New[int](errorE[int]{})
	r5, ok := enum.First()
	is.Equal(false, ok)
	is.Equal(0, r5)
	is.Error(enum.Error())

	r6, ok := enumerator.FromMap(map[string]int{
		"a": 1,
	}).Last()
	is.Equal(true, ok)
	is.Equal(r6, enumerator.KeyValuePair[string, int]{Key: "a", Value: 1})
}

func TestSortBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enumerator.From([]string{"aa", "bbb", "c"}).SortBy(func(i, j string) bool {
		return len(i) > len(j)
	}).ToSlice()
	is.Equal([]string{"bbb", "aa", "c"}, r1)
}

func TestReverse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enumerator.From([]int{1, 2, 3}).Reverse().ToSlice()
	is.Equal([]int{3, 2, 1}, r1)
}

func TestToMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enumerator.From([]int{3, 2, 1}).ToMap()
	is.Equal(map[int]int{0: 3, 1: 2, 2: 1}, r1)
}

func TestTake(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enumerator.From([]int{1, 2, 3, 4, 5}).Take(3).ToSlice()
	is.Equal([]int{1, 2, 3}, r1)

	r2 := enumerator.From([]int{1, 2, 3, 4, 5}).Take(0).ToSlice()
	is.Equal([]int{}, r2)

	r3 := enumerator.From([]int{1, 2, 3, 4, 5}).Take(6).ToSlice()
	is.Equal([]int{1, 2, 3, 4, 5}, r3)

	r4 := enumerator.From([]int{1, 2, 3, 4, 5}).Take(-1).ToSlice()
	is.Equal([]int{}, r4)
}

func TestIsAll(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enumerator.From([]int{1, 1, 1, 1, 1}).IsAll(func(item int) bool { return item == 1 })
	is.Equal(true, r1)

	r2 := enumerator.From([]int{1, 2, 1, 1, 1}).IsAll(func(item int) bool { return item == 1 })
	is.Equal(false, r2)
}

func TestIsAny(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enumerator.From([]int{1, 1, 1, 1, 1}).IsAny(func(item int) bool { return item == 2 })
	is.Equal(false, r1)

	r2 := enumerator.From([]int{1, 2, 1, 1, 1}).IsAny(func(item int) bool { return item == 2 })
	is.Equal(true, r2)
}

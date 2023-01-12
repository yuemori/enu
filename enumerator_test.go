package enu_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enu"
)

func TestCount(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.From([]int{1, 2, 3, 4, 5}).Count()

	is.Equal(5, r)
}

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.From([]int{1, 2, 3, 4, 5}).Filter(func(i int, _ int) bool {
		return i%2 == 0
	}).ToSlice()

	is.Equal([]int{2, 4}, r)
}

func TestReject(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.From([]int{1, 2, 3, 4, 5}).Reject(func(i int, _ int) bool {
		return i%2 == 0
	}).ToSlice()

	is.Equal([]int{1, 3, 5}, r)
}

func TestFind(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	slices := enu.From([]int{1, 2, 3, 4, 5})

	r1, ok := slices.Find(func(item, _ int) bool {
		return item%2 == 0
	})
	is.Equal(true, ok)
	is.Equal(2, r1)

	r2, ok := slices.Find(func(item, _ int) bool {
		return item%2 == 0
	})
	is.Equal(true, ok)
	is.Equal(2, r2)

	_, ok = slices.Find(func(item, _ int) bool {
		return item > 10
	})
	is.Equal(false, ok)

	r3, ok := slices.Find(func(item, _ int) bool {
		return item%2 == 0
	})
	is.Equal(true, ok)
	is.Equal(2, r3)

	r4, ok := slices.Find(func(_, index int) bool {
		return index == 3
	})
	is.Equal(true, ok)
	is.Equal(4, r4)
}

func TestNth(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	slices := enu.From([]int{1, 2, 3, 4, 5})

	r1, ok := slices.Nth(3)
	is.Equal(true, ok)
	is.Equal(4, r1)

	r2, ok := slices.Nth(1)
	is.Equal(true, ok)
	is.Equal(2, r2)

	_, ok = slices.Nth(5)
	is.Equal(false, ok)
	//
	// r3, ok := slices.Nth(-1)
	// is.Equal(true, ok)
	// is.Equal(5, r3)
}

func TestFirst(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	slices := enu.From([]int{1, 2, 3, 4, 5})

	r1, ok := slices.First()
	is.Equal(true, ok)
	is.Equal(1, r1)

	r2, ok := slices.First()
	is.Equal(true, ok)
	is.Equal(1, r2)

	r3, ok := enu.From([]int{}).First()
	is.Equal(false, ok)
	is.Equal(0, r3)

	type dummy struct{}

	r4, ok := enu.From([]dummy{}).First()
	is.Equal(false, ok)
	is.Equal(dummy{}, r4)

	r5, ok := enu.From([]*dummy{}).First()
	is.Equal(false, ok)
	is.Nil(r5)
}

func TestLast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1, ok := enu.From([]int{1, 2, 3, 4, 5}).Last()
	is.Equal(true, ok)
	is.Equal(5, r1)

	r2, ok := enu.From([]int{}).Last()
	is.Equal(false, ok)
	is.Equal(0, r2)

	type dummy struct{}

	r3, ok := enu.From([]dummy{}).Last()
	is.Equal(false, ok)
	is.Equal(dummy{}, r3)

	r4, ok := enu.From([]*dummy{}).Last()
	is.Equal(false, ok)
	is.Nil(r4)

	r6, ok := enu.FromMap(map[string]int{
		"a": 1,
	}).Last()
	is.Equal(true, ok)
	is.Equal(r6, enu.KeyValuePair[string, int]{Key: "a", Value: 1})
}

func TestSortBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enu.From([]string{"aa", "bbb", "c"}).SortBy(func(i, j string) bool {
		return len(i) > len(j)
	}).ToSlice()
	is.Equal([]string{"bbb", "aa", "c"}, r1)
}

func TestReverse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enu.From([]int{1, 2, 3}).Reverse().ToSlice()
	is.Equal([]int{3, 2, 1}, r1)
}

func TestToMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enu.From([]int{3, 2, 1}).ToMap()
	is.Equal(map[int]int{0: 3, 1: 2, 2: 1}, r1)
}

func TestTake(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enu.From([]int{1, 2, 3, 4, 5}).Take(3).ToSlice()
	is.Equal([]int{1, 2, 3}, r1)

	r2 := enu.From([]int{1, 2, 3, 4, 5}).Take(0).ToSlice()
	is.Equal([]int{}, r2)

	r3 := enu.From([]int{1, 2, 3, 4, 5}).Take(6).ToSlice()
	is.Equal([]int{1, 2, 3, 4, 5}, r3)

	r4 := enu.From([]int{1, 2, 3, 4, 5}).Take(2).ToSlice()
	is.Equal([]int{1, 2}, r4)
}

func TestIsAll(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enu.From([]int{1, 1, 1, 1, 1}).IsAll(func(item int) bool { return item == 1 })
	is.Equal(true, r1)

	r2 := enu.From([]int{1, 2, 1, 1, 1}).IsAll(func(item int) bool { return item == 1 })
	is.Equal(false, r2)
}

func TestIsAny(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enu.From([]int{1, 1, 1, 1, 1}).IsAny(func(item int) bool { return item == 2 })
	is.Equal(false, r1)

	r2 := enu.From([]int{1, 2, 1, 1, 1}).IsAny(func(item int) bool { return item == 2 })
	is.Equal(true, r2)
}

func TestGetEnumerator(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	enumerator := enu.FromOrdered([]int{1, 2, 3}).GetEnumerator()
	r := enu.ToNumeric(enumerator).Sum()
	is.Equal(6, r)
}

package enu_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enu"
)

func TestMapToSlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.FromMap(map[int]string{1: "foo", 2: "bar", 3: "baz"}).ToSlice()

	is.ElementsMatch([]enu.KeyValuePair[int, string]{
		{Key: 1, Value: "foo"},
		{Key: 2, Value: "bar"},
		{Key: 3, Value: "baz"},
	}, r)
}

func TestMapToCount(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.FromMap(map[int]string{1: "foo", 2: "bar", 3: "baz"}).Count()
	is.Equal(3, r)
}

func TestMapFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.FromMap(map[int]string{1: "foo", 2: "bar", 3: "baz", 4: "boo"}).Filter(func(kv enu.KeyValuePair[int, string], _ int) bool {
		return kv.Key%2 == 0
	}).ToMap()

	is.Equal(map[int]string{2: "bar", 4: "boo"}, r)
}

func TestMapKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.FromMap(map[int]string{1: "foo", 2: "bar", 3: "baz"}).Keys()

	is.ElementsMatch([]int{1, 2, 3}, r)
}

func TestMapValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := enu.FromMap(map[int]string{1: "foo", 2: "bar", 3: "baz"}).Values()
	is.ElementsMatch([]string{"foo", "bar", "baz"}, r)
}

func TestEnumeratorToMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	e := enu.From([]enu.KeyValuePair[int, string]{
		{Key: 1, Value: "foo"},
		{Key: 2, Value: "bar"},
		{Key: 3, Value: "baz"},
	})
	r := enu.ToMap[int, string](e).Keys()

	is.ElementsMatch([]int{1, 2, 3}, r)
}

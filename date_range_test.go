package enu_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enu"
)

func TestDateRange(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enu.NewDateRange(
		time.Date(2023, 5, 20, 21, 59, 59, 0, time.UTC),
		time.Date(2023, 5, 20, 23, 59, 59, 0, time.UTC),
		time.Hour,
	).ToSlice()
	is.Equal([]time.Time{
		time.Date(2023, 5, 20, 21, 59, 59, 0, time.UTC),
		time.Date(2023, 5, 20, 22, 59, 59, 0, time.UTC),
		time.Date(2023, 5, 20, 23, 59, 59, 0, time.UTC),
	}, r1)

	r2 := enu.NewDateRange(
		time.Date(2023, 5, 20, 21, 59, 59, 0, time.UTC),
		time.Date(2023, 5, 23, 21, 59, 59, 0, time.UTC),
		time.Hour*24,
	).ToSlice()
	is.Equal([]time.Time{
		time.Date(2023, 5, 20, 21, 59, 59, 0, time.UTC),
		time.Date(2023, 5, 21, 21, 59, 59, 0, time.UTC),
		time.Date(2023, 5, 22, 21, 59, 59, 0, time.UTC),
		time.Date(2023, 5, 23, 21, 59, 59, 0, time.UTC),
	}, r2)

	r3 := enu.NewDateRange(
		time.Date(2023, 5, 20, 21, 59, 59, 0, time.UTC),
		time.Date(2099, 5, 23, 21, 59, 59, 0, time.UTC),
		time.Hour*24,
	).Take(5).ToSlice()
	is.Equal([]time.Time{
		time.Date(2023, 5, 20, 21, 59, 59, 0, time.UTC),
		time.Date(2023, 5, 21, 21, 59, 59, 0, time.UTC),
		time.Date(2023, 5, 22, 21, 59, 59, 0, time.UTC),
		time.Date(2023, 5, 23, 21, 59, 59, 0, time.UTC),
		time.Date(2023, 5, 24, 21, 59, 59, 0, time.UTC),
	}, r3)

	r4 := enu.NewDateRange(
		time.Date(2023, 5, 20, 23, 59, 59, 0, time.UTC),
		time.Date(2023, 5, 20, 21, 59, 59, 0, time.UTC),
		time.Hour,
	).ToSlice()
	is.Equal([]time.Time{}, r4)
}

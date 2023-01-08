package enu_test

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enu"
)

type Month time.Month

func (m Month) Next(step int) enu.RangeValuer[time.Month, int] {
	val := int(m) + step
	if val > 12 {
		val -= 12
	}

	return Month(val)
}

func (m Month) Value() time.Month {
	return time.Month(m)
}

func (m Month) Compare(other time.Month) int {
	if int(m) < int(other) {
		return -1
	}
	if int(m) > int(other) {
		return 1
	}
	return 0
}

func TestRange(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := enu.New(enu.NewRange[time.Month, int](Month(time.January), Month(time.July), 1)).ToSlice()
	is.Equal([]time.Month{time.January, time.February, time.March, time.April, time.May, time.June, time.July}, r1)

	r2 := enu.New(enu.NewRange[time.Month, int](Month(time.January), Month(math.MaxInt), 1)).Take(13).ToSlice()
	is.Equal([]time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
		time.January,
	}, r2)
}

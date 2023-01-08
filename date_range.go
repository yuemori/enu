package enu

import (
	"time"
)

func NewDateRange(start, end time.Time, stepDuration time.Duration) *Enumerator[time.Time] {
	return &Enumerator[time.Time]{iter: NewDateRangeEnumerator(start, end, stepDuration)}
}

func NewDateRangeEnumerator(start, end time.Time, stepDuration time.Duration) *RangeEnumerator[time.Time, time.Duration] {
	return &RangeEnumerator[time.Time, time.Duration]{
		min:  Time{value: start},
		max:  Time{value: end},
		step: stepDuration,
	}
}

type Time struct {
	value time.Time
}

func (t Time) Compare(other time.Time) int {
	if t.value.Before(other) {
		return -1
	}
	if t.value.After(other) {
		return 1
	}
	return 0
}

func (t Time) Value() time.Time {
	return t.value
}

func (t Time) Next(duration time.Duration) RangeValuer[time.Time, time.Duration] {
	return Time{value: t.value.Add(duration)}
}

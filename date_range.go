package enu

import (
	"time"
)

func NewDateRange(start, end time.Time, stepDuration time.Duration) *Enumerator[time.Time] {
	return &Enumerator[time.Time]{iter: NewDateRangeEnumerator(start, end, stepDuration)}
}

func NewDateRangeEnumerator(start, end time.Time, stepDuration time.Duration) *RangeEnumerator[time.Time] {
	return &RangeEnumerator[time.Time]{
		min: ComparableTime{value: start, step: stepDuration},
		max: ComparableTime{value: end, step: stepDuration},
	}
}

type ComparableTime struct {
	value time.Time
	step  time.Duration
}

func compareTime(value time.Time, step time.Duration) ComparableTime {
	return ComparableTime{value: value, step: step}
}

func (t ComparableTime) Compare(other time.Time) int {
	if t.value.Before(other) {
		return -1
	}
	if t.value.After(other) {
		return 1
	}
	return 0
}

func (t ComparableTime) Value() time.Time {
	return t.value
}

func (t ComparableTime) Next() IComparable[time.Time] {
	return ComparableTime{value: t.value.Add(t.step), step: t.step}
}

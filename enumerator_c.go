package enumerator

import (
	"github.com/samber/lo"
)

func (e *EnumeratorC[T]) Uniq() *EnumeratorC[T] {
	e.iter = newSliceEnumerator(lo.Uniq(e.ToSlice()))
	return e
}

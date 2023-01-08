package enumerator

import (
	"sort"

	"github.com/samber/lo"
  {{.ImportPkg}}
)

type IEnumerable{{.Suffix}}[{{.TypeWithConstraint}}] interface {
	Next() ({{.ItemType}}, error)
	Stop()
}

type Enumerator{{.Suffix}}[{{.TypeWithConstraint}}] struct {
	iter IEnumerable{{.Suffix}}[{{.Type}}]
	err  error
}

func New{{.Suffix}}[{{.TypeWithConstraint}}](e IEnumerable{{.Suffix}}[{{.Type}}]) *Enumerator{{.Suffix}}[{{.Type}}] {
	return &Enumerator{{.Suffix}}[{{.Type}}]{iter: e}
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Error() error {
	return e.err
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Each(iteratee func(item {{.ItemType}}, index int)) *Enumerator{{.Suffix}}[{{.Type}}] {
	if e.err == nil {
		each{{.Suffix}}(e.iter, iteratee)
	}

	return e
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Count() int {
	v := 0
	if e.err != nil {
		return v
	}
	each{{.Suffix}}(e.iter, func(item {{.ItemType}}, _ int) {
		v += 1
	})
	return v
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) ToSlice() []{{.ItemType}} {
	result := make([]{{.ItemType}}, 0)
	if e.err != nil {
		return result
	}

	for {
		item, err := e.iter.Next()
		if err == Done {
			break
		}
		if err != nil {
			e.err = err
			break
		}
		result = append(result, item)
	}
	return result
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Filter(predicate func(item {{.ItemType}}, index int) bool) *Enumerator{{.Suffix}}[{{.Type}}] {
	if e.err != nil {
		return e
	}
	e.iter = newSliceEnumerator(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) First() ({{.ItemType}}, bool) {
	if e.err != nil {
		var empty {{.ItemType}}
		return empty, false
	}
	item, err := e.iter.Next()
	if err != nil {
		var empty {{.ItemType}}
		if err != Done {
			e.err = err
		}
		return empty, false
	}
	return item, true
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Last() ({{.ItemType}}, bool) {
	if e.err != nil {
		var empty {{.ItemType}}
		return empty, false
	}
	prev, err := e.iter.Next()
	if err == Done {
		var empty {{.ItemType}}
		return empty, false
	}
	for {
		item, err := e.iter.Next()
		if err == Done {
			return prev, true
		}
		prev = item
		if err != nil {
			var empty {{.ItemType}}
			e.err = err
			return empty, false
		}
	}
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Reverse() *Enumerator{{.Suffix}}[{{.Type}}] {
	e.iter = newSliceEnumerator(lo.Reverse(e.ToSlice()))
	return e
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) SortBy(sorter func(i, j {{.ItemType}}) bool) *Enumerator{{.Suffix}}[{{.Type}}] {
	res := e.ToSlice()
	sort.SliceStable(res, func(i, j int) bool {
		return sorter(res[i], res[j])
	})
	e.iter = newSliceEnumerator(res)
	return e
}

func each{{.Suffix}}[{{.TypeWithConstraint}}](iter IEnumerable{{.Suffix}}[{{.Type}}], iteratee func(item {{.ItemType}}, index int)) {
	index := 0
	for {
		item, err := iter.Next()
		if err == Done {
			break
		}
		iteratee(item, index)
		index += 1
	}
}

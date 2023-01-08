package enumerator

import (
	"sort"

	"github.com/samber/lo"
)

type IEnumerable{{.Suffix}}[{{.TypeWithConstraint}}] interface {
	Next() ({{.Type}}, error)
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

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Each(iteratee func(item {{.Type}}, index int)) *Enumerator{{.Suffix}}[{{.Type}}] {
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
	each{{.Suffix}}(e.iter, func(item {{.Type}}, _ int) {
		v += 1
	})
	return v
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) ToSlice() []{{.Type}} {
	result := make([]{{.Type}}, 0)
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

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Filter(predicate func(item {{.Type}}) bool) *Enumerator{{.Suffix}}[{{.Type}}] {
	if e.err != nil {
		return e
	}
	e.iter = newSliceEnumerator(lo.Filter(e.ToSlice(), func(item {{.Type}}, _ int) bool { return predicate(item) }))
	return e
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) First() ({{.Type}}, bool) {
	if e.err != nil {
		var empty {{.Type}}
		return empty, false
	}
	item, err := e.iter.Next()
	if err != nil {
		var empty {{.Type}}
		if err != Done {
			e.err = err
		}
		return empty, false
	}
	return item, true
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Last() ({{.Type}}, bool) {
	if e.err != nil {
		var empty {{.Type}}
		return empty, false
	}
	prev, err := e.iter.Next()
	if err == Done {
		var empty {{.Type}}
		return empty, false
	}
	for {
		item, err := e.iter.Next()
		if err == Done {
			return prev, true
		}
		prev = item
		if err != nil {
			var empty {{.Type}}
			e.err = err
			return empty, false
		}
	}
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) SortBy(sorter func(i, j {{.Type}}) bool) *Enumerator{{.Suffix}}[{{.Type}}] {
	res := e.ToSlice()
	sort.SliceStable(res, func(i, j int) bool {
		return sorter(res[i], res[j])
	})
	e.iter = newSliceEnumerator(res)
	return e
}

func each{{.Suffix}}[{{.TypeWithConstraint}}](iter IEnumerable{{.Suffix}}[{{.Type}}], iteratee func(item {{.Type}}, index int)) {
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

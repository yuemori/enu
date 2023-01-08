package enu

import (
	"sort"

	"github.com/samber/lo"
  {{.ImportPkg}}
)

type IEnumerable{{.Suffix}}[{{.TypeWithConstraint}}] interface {
	Next() ({{.ItemType}}, bool)
}

type Enumerator{{.Suffix}}[{{.TypeWithConstraint}}] struct {
	iter IEnumerable{{.Suffix}}[{{.Type}}]
}

func New{{.Suffix}}[{{.TypeWithConstraint}}](e IEnumerable{{.Suffix}}[{{.Type}}]) *Enumerator{{.Suffix}}[{{.Type}}] {
	return &Enumerator{{.Suffix}}[{{.Type}}]{iter: e}
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Each(iteratee func(item {{.ItemType}}, index int)) *Enumerator{{.Suffix}}[{{.Type}}] {
  each{{.Suffix}}(e.iter, iteratee)

	return e
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Count() int {
	v := 0
	each{{.Suffix}}(e.iter, func(item {{.ItemType}}, _ int) {
		v += 1
	})
	return v
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) ToSlice() []{{.ItemType}} {
	result := make([]{{.ItemType}}, 0)

	for {
		item, ok := e.iter.Next()
		if !ok {
			break
		}
		result = append(result, item)
	}
	return result
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Filter(predicate func(item {{.ItemType}}, index int) bool) *Enumerator{{.Suffix}}[{{.Type}}] {
	e.iter = newSliceEnumerator(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) First() ({{.ItemType}}, bool) {
	item, ok := e.iter.Next()
	if !ok {
		var empty {{.ItemType}}
		return empty, false
	}
	return item, true
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Last() ({{.ItemType}}, bool) {
	prev, ok := e.iter.Next()
	if !ok {
		var empty {{.ItemType}}
		return empty, false
	}
	for {
		item, ok := e.iter.Next()
		if !ok {
			return prev, true
		}
		prev = item
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
		item, ok := iter.Next()
		if !ok {
			break
		}
		iteratee(item, index)
		index += 1
	}
}


func (e *Enumerator{{.Suffix}}[{{.Type}}]) Reject(predicate func(item {{.ItemType}}, index int) bool) *Enumerator{{.Suffix}}[{{.Type}}] {
	e.iter = newSliceEnumerator(lo.Reject(e.ToSlice(), predicate))
	return e
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) IsAll(predicate func(item {{.ItemType}}) bool) bool {
	flag := true
	for {
		item, ok := e.iter.Next()
		if !ok {
			break
		}
		if !predicate(item) {
			flag = false
			break
		}
	}
	return flag
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) IsAny(predicate func(item {{.ItemType}}) bool) bool {
	flag := false
	for {
		item, ok := e.iter.Next()
		if !ok {
			break
		}
		if predicate(item) {
			flag = true
			break
		}
	}
	return flag
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Take(num int) *Enumerator{{.Suffix}}[{{.Type}}] {
	result := []{{.ItemType}}{}
	index := 0
	for {
		item, ok := e.iter.Next()
		if !ok {
			break
		}
		if index >= num {
			break
		}
		result = append(result, item)
		index += 1
	}
	e.iter = newSliceEnumerator(result)
	return e
}

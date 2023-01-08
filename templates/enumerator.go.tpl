package enu

import (
	"sort"

	"github.com/samber/lo"
  {{.ImportPkg}}
)

type IEnumerable{{.Suffix}}[{{.TypeWithConstraint}}] interface {
	Next() ({{.ItemType}}, bool)
	Stop()
	Reset()
}

type Enumerator{{.Suffix}}[{{.TypeWithConstraint}}] struct {
	iter      IEnumerable{{.Suffix}}[{{.Type}}]
	result    []{{.ItemType}}
	isStopped bool
}

func New{{.Suffix}}[{{.TypeWithConstraint}}](e IEnumerable{{.Suffix}}[{{.Type}}]) *Enumerator{{.Suffix}}[{{.Type}}] {
	return &Enumerator{{.Suffix}}[{{.Type}}]{iter: e}
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Each(iteratee func(item {{.ItemType}}, index int)) {
	if e.isStopped {
		lo.ForEach(e.result, iteratee)
		return
	}

	result := []{{.ItemType}}{}
	index := 0
	for {
		item, ok := e.iter.Next()
		if !ok {
			break
		}
		iteratee(item, index)
		index += 1
		result = append(result, item)
	}
	e.iter.Stop()
	e.isStopped = true
	e.iter = newSliceEnumerator(result)
	e.result = result
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) ToSlice() []{{.ItemType}} {
	if e.isStopped {
		return e.result
	}
	e.Each(func({{.ItemType}}, int) {})
	return e.result
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Count() int {
	v := 0
	e.Each(func(item {{.ItemType}}, _ int) {
		v += 1
	})
	return v
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Filter(predicate func(item {{.ItemType}}, index int) bool) *Enumerator{{.Suffix}}[{{.Type}}] {
	e.swap(lo.Filter(e.ToSlice(), predicate))
	return e
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Nth(index int) ({{.ItemType}}, bool) {
	item, err := lo.Nth(e.ToSlice(),index)
	if err != nil {
		return empty[{{.ItemType}}](), false
	}

	return item, true
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) First() ({{.ItemType}}, bool) {
	if e.isStopped {
		if len(e.result) == 0 {
			return empty[{{.ItemType}}](), false
		}
		return e.result[0], true
	}
	item, ok := e.iter.Next()
	if !ok {
		e.swap([]{{.ItemType}}{})
		e.iter.Stop()
		e.isStopped = true
		return empty[{{.ItemType}}](), false
	}
	e.iter.Reset()
	return item, true
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Last() ({{.ItemType}}, bool) {
	result := e.ToSlice()
	if len(result) == 0 {
		return empty[{{.ItemType}}](), false
	}
	return result[len(result)-1], true
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Reverse() *Enumerator{{.Suffix}}[{{.Type}}] {
	e.swap(lo.Reverse(e.ToSlice()))
	return e
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) SortBy(sorter func(i, j {{.ItemType}}) bool) *Enumerator{{.Suffix}}[{{.Type}}] {
	res := e.ToSlice()
	sort.SliceStable(res, func(i, j int) bool {
		return sorter(res[i], res[j])
	})
	e.swap(res)
	return e
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Reject(predicate func(item {{.ItemType}}, index int) bool) *Enumerator{{.Suffix}}[{{.Type}}] {
	e.swap(lo.Reject(e.ToSlice(), predicate))
	return e
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) IsAll(predicate func(item {{.ItemType}}) bool) bool {
	if e.isStopped {
		for _, item := range e.ToSlice() {
			if !predicate(item) {
				return false
			}
			return true
		}
	}

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
	e.iter.Reset()
	return flag
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) IsAny(predicate func(item {{.ItemType}}) bool) bool {
	if e.isStopped {
		for _, item := range e.ToSlice() {
			if predicate(item) {
				return true
			}
			return false
		}
	}

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
	e.iter.Reset()
	return flag
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) Take(num uint) *Enumerator{{.Suffix}}[{{.Type}}] {
	if e.isStopped {
		e.swap(lo.Subset(e.result, 0, num-1))
		return e
	}

	result := []{{.ItemType}}{}
	index := 0
	for {
		item, ok := e.iter.Next()
		if !ok {
			break
		}
		if uint(index) >= num {
			break
		}
		result = append(result, item)
		index += 1
	}
	e.swap(result)
	return e
}

func (e *Enumerator{{.Suffix}}[{{.Type}}]) swap(result []{{.ItemType}}) {
	if !e.isStopped {
		e.iter.Stop()
		e.isStopped = true
	}
	e.iter = newSliceEnumerator(result)
	e.result = result
}

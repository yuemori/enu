package enumerator

func (e *Enumerator[T]) ToMap() map[int]T {
	result := map[int]T{}
	_ = e.Each(func(item T, index int) {
		result[index] = item
	})
	return result
}

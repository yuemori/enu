package enu

type IEnumerator[T any] interface {
	Next() (T, bool)
	Reset()
	Stop()
}

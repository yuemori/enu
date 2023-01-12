package enu

type Enumerator[T any] interface {
	Next() (T, bool)
	Reset()
	Stop()
}

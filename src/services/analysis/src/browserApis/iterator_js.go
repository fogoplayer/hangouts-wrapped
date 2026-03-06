package browserApis

import "syscall/js"

type Iterator[T any] struct {
	value js.Value
}

func (iterator Iterator[T]) Done() bool {
	return iterator.value.Get("done").Bool()
}

func (iterator Iterator[T]) Value(jsToGoConverter func(js.Value) T) T {
	return jsToGoConverter(iterator.value.Get("value"))
}

func IteratorFromJs[T any](v js.Value) Iterator[T] { return Iterator[T]{v} }

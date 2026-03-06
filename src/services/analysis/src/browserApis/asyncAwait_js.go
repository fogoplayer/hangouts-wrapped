package browserApis

import "syscall/js"

// func await[T any](channel chan[T]){

// }

type Promise[T any] struct {
	value js.Value
}

func (p Promise[T]) ToChannel(jsToGoConverter func(js.Value) T) chan T {
	channel := make(chan T)
	p.value.Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
		promiseValue := args[0]
		channel <- jsToGoConverter(promiseValue)
		return nil
	}))

	return channel
}

package browserApis

import "syscall/js"

// func await[T any](channel chan[T]){

// }

type Promise[T any] struct {
	value js.Value
}

func (p Promise[T]) ToChannel(jsToGoConverter func(js.Value) T, channels ...chan T) chan T {

	var channel chan T
	switch len(channels) {
	case 0:
		channel = make(chan T)
	case 1:
		channel = channels[0]
	default:
		panic(nil)
	}
	p.value.Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
		promiseValue := args[0]
		channel <- jsToGoConverter(promiseValue)
		return nil
	}))

	return channel
}

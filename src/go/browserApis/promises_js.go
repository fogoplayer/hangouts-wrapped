package browserApis

import (
	"errors"
	"syscall/js"

	"zarinloosli.com/hangouts-wrapped/util"
)

type Promise[T any] struct {
	value      js.Value
	jsToGoFunc func(js.Value) (T, error)
}

func (p Promise[T]) ToChannel(channels ...chan PromiseResult[T]) chan PromiseResult[T] {

	var channel chan PromiseResult[T]
	switch len(channels) {
	case 0:
		channel = make(chan PromiseResult[T])
	case 1:
		channel = channels[0]
	default:
		util.WrongNumberOfArgumentsPanic(len(channels))
	}
	p.value.
		// Then
		Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
			promiseValue := args[0]
			goValue, err := p.jsToGoFunc(promiseValue)
			channel <- PromiseResult[T]{goValue, err}
			return nil
		})).
		// Catch
		Call("catch", js.FuncOf(func(this js.Value, args []js.Value) any {
			err := args[0]
			errorMessage := err.Get("reason").String()
			goError := errors.New(errorMessage)
			channel <- PromiseResult[T]{err: goError}
			return nil
		}))

	return channel
}

type PromiseResult[T any] struct {
	value T
	err   error
}

func (result PromiseResult[T]) Value() (T, error) {
	return result.value, result.err
}

func Await[T any](p Promise[T]) (T, error) {
	result := <-p.ToChannel()
	return result.Value()
}

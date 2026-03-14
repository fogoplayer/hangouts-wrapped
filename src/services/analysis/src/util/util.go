package util

import (
	"fmt"
	"reflect"
)

func WrongNumberOfArgumentsPanic(numberOfArguments int) {
	panic(fmt.Errorf("Wrong number of arguments passed: %d", numberOfArguments))
}

type UnableToCastHandleError error

func CreateUnableToCastFromAToBError[T any, U any](handle T) UnableToCastHandleError {
	return fmt.Errorf("Unable to cast %v to %s", handle, reflect.TypeFor[U]().Name())
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func UseVar(any) {}

func ListMap[T any, U any](array []T, converter func(T) U) []U {
	result := make([]U, len(array))
	for i, v := range array {
		result[i] = converter(v)
	}
	return result
}

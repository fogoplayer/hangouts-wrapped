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

func MapMap[
	InputKey comparable, InputValue any, OutputKey comparable, OutputValue any,
](
	value map[InputKey]InputValue,
	converter func(key InputKey, value InputValue) (OutputKey, OutputValue),
) map[OutputKey]OutputValue {
	result := make(map[OutputKey]OutputValue)
	for key, value := range value {
		newKey, newValue := converter(key, value)
		result[newKey] = newValue
	}
	return result
}

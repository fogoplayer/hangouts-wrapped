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

func UseVar(...any) {}

func StartsWithWords(candidate string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if len(candidate) < len(prefix) {
			continue
		}
		firstWord := candidate[0:len(prefix)]
		if firstWord == prefix {
			return true
		}
	}
	return false
}

func ToAny[T any](val T) any {
	return val
}

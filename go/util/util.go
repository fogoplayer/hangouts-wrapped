package util

import (
	"fmt"
	"reflect"
)

func WrongNumberOfArgumentsPanic(numberOfArguments int) {
	panic(fmt.Errorf("Wrong number of arguments passed: %d", numberOfArguments))
}

// Returns the first element of a variadic array if it exists, and a boolean describing whether it exists or not
func ExtractOptionalArgument[T any](variadicArgs []T) (T, bool) { // TODO use this everywhere
	var t T

	argsLength := len(variadicArgs)

	if argsLength > 0 {
		return variadicArgs[0], true
	}

	return t, false
}

func ExtractOptionalArgumentWithDefault[T any](variadicArgs []T, defaultValue T) T {
	result, exists := ExtractOptionalArgument(variadicArgs)

	if exists {
		return result
	}

	return defaultValue
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

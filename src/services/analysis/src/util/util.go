package util

import (
	"fmt"
	"reflect"
	"strings"
)

// TODO break this file up

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

func StartsWithWords(candidate string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		firstWord := strings.Split(candidate, " ")[0]
		if firstWord == prefix {
			return true
		}
	}
	return false
}

// /////// //
// Mapping //
// /////// //

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

func GetMapKeys[InputKey comparable, OutputKey any](m map[InputKey]OutputKey) []InputKey {
	keys := make([]InputKey, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func GetMapVals[InputKey comparable, OutputKey any](m map[InputKey]OutputKey) []OutputKey {
	values := make([]OutputKey, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

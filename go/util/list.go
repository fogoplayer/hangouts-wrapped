package util

import (
	"fmt"
	"math/rand"
	"reflect"
)

func ListForEach[T any](array []T, converter func(T)) {
	for _, v := range array {
		converter(v)
	}
}

func ListMap[T any, U any](array []T, converter func(T) U) []U {
	result := make([]U, len(array))
	for i, v := range array {
		result[i] = converter(v)
	}
	return result
}

func ListsAreEqual[T any](list1 []T, list2 []T, comparators ...func(a, b T) bool) bool {
	if len(list1) != len(list2) {
		return false
	}

	var elementsAreEqual func(a, b T) bool
	if len(comparators) > 1 {
		panic(fmt.Errorf("Passed in too many comparators (%d)", len(comparators)))
	} else if len(comparators) == 1 {
		elementsAreEqual = comparators[0]
	} else {
		elementsAreEqual = func(a, b T) bool {
			v := reflect.ValueOf(a)
			if v.Comparable() {
				u := reflect.ValueOf(b)
				return v.Equal(u)
			}
			panic("type cannot be compared")
		}
	}

	for i := range len(list1) {
		if !elementsAreEqual(list1[i], list2[i]) {
			return false
		}
	}

	return true
}

func CopyList[T any](list []T) []T {
	value := make([]T, len(list))
	copy(value, list)
	return value
}

func RandomFromList[T any](list []T) T {
	randomIndex := rand.Intn(len(list))
	return list[randomIndex]
}

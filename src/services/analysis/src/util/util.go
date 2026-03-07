package util

import (
	"fmt"
)

func WrongNumberOfArgumentsPanic(numberOfArguments int) {
	panic(fmt.Errorf("Wrong number of arguments passed: %d", numberOfArguments))
}

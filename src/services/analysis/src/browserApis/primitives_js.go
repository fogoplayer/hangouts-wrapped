package browserApis

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
	"syscall/js"
)

func GetIntFromJsValue(jsValue js.Value) (int64, error) {
	valueBytes := []byte{}
	js.CopyBytesToGo(valueBytes, jsValue)

	value, bytesRead := binary.Varint(valueBytes)
	if bytesRead <= 0 {
		return 0, errors.New("Couldn't read int")
	}
	return value, nil
}

func TypeMismatchError[T any](value js.Value) error {
	return fmt.Errorf("%v cannot be coerced into %s", value, reflect.TypeFor[T]().Name())
}

type JSWrapper interface {
	StoreAsGlobalVariable(string)
}

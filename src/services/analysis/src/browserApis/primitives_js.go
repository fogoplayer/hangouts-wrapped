package browserApis

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
	"syscall/js"
)

func TypeMismatchError[T any](value js.Value) error {
	return fmt.Errorf("%v cannot be coerced into %s", value, reflect.TypeFor[T]().Name())
}

type JSWrapper interface {
	StoreAsGlobalVariable(string)
}

var globalsSet = make(map[string]bool)

func SetGlobal(name string, x js.Value) {
	if globalsSet[name] {
		fmt.Println("overwriting", name)
	} else {
		fmt.Println("Setting", name)
	}
	js.Global().Set(name, x)
	globalsSet[name] = true
	fmt.Println("Set", name)
}

// ////////////////// //
// GoFromJSConverters //
// ////////////////// //

func JsFromJsReturnValueUnchanged(v js.Value) (js.Value, error) {
	return v, nil
}

// Converts as JS byte array into a Golang byte array
//
// Panics if value is of some other type, does not actually return an error.
// However, it needs to match the signature of other *FromJs methods
func ByteArrayFromJs(v js.Value) ([]byte, error) {
	data := make([]byte, v.Length())
	js.CopyBytesToGo(data, v)
	return data, nil
}

func GetIntFromJsValue(jsValue js.Value) (int64, error) {
	valueBytes := []byte{}
	js.CopyBytesToGo(valueBytes, jsValue)

	value, bytesRead := binary.Varint(valueBytes)
	if bytesRead <= 0 {
		return 0, errors.New("Couldn't read int")
	}
	return value, nil
}

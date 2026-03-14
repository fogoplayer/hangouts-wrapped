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

// /////////////////// //
// GoFromJS Converters //
// /////////////////// //

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

// /////////////////// //
// JSFromGo Converters //
// /////////////////// //

func CreateJsObject(name string, args ...any) js.Value { // TODO
	return js.Global().Get(name).New(args...)
}

// Uint8Array

type Uint8Array struct {
	js.Value
}

func Uint8ArrayFromGo(bytes []byte) js.Value /* Uint8Array */ {
	jsUint8Array := CreateUint8Array(len(bytes))
	js.CopyBytesToJS(jsUint8Array /* .Value */, bytes)
	return /* Uint8Array( */ jsUint8Array /* ) */
}

func CreateUint8Array(len int) js.Value /* Uint8Array */ {
	return /* Uint8Array{ */ CreateJsObject("Uint8Array", len) /* } */
}

// String

type JSString struct {
	js.Value
}

func StringFromGoBytes(bytes []byte) JSString {
	jsUint8Array := Uint8ArrayFromGo(bytes)
	textDecoder := CreateTextDecoder()
	return textDecoder.Decode(jsUint8Array)
}

func StringFromGoString(value string) JSString {
	return JSString{js.ValueOf(value)}
}

// TextDecoder

type TextDecoder struct {
	js.Value
}

func (textDecoder TextDecoder) Decode(bytes js.Value /* Uint8Array */) JSString {
	return JSString{textDecoder.Value.Call("decode", bytes)}
}

func CreateTextDecoder() TextDecoder {
	return TextDecoder{CreateJsObject("TextDecoder")}
}

// new TextDecoder().decode(uint8array);

// ////////////////////// //
// JS Object Constructors //
// ////////////////////// //

// TODO is there a way to use type aliases or embedded types to allow passing a JS type where a JS value is accepted?
// TODO make the types private so you have to use constructors, then use InstanceOf to make sure they are what they say they are?

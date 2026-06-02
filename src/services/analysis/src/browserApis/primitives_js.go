package browserApis

import (
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
// TODO add error handling using InstanceOf
func ByteArrayFromJs(v js.Value) ([]byte, error) {
	data := make([]byte, v.Length())
	js.CopyBytesToGo(data, v)
	return data, nil
}

func IntFromJs(jsValue js.Value) (int, error) {
	if jsValue.Type() != js.TypeNumber {
		return -1, TypeMismatchError[int](jsValue)
	}
	return jsValue.Int(), nil
}

func StringFromJs(jsValue js.Value) (string, error) {
	if jsValue.Type() != js.TypeString {
		return "", fmt.Errorf("&s is not a string", jsValue)
	}
	return jsValue.String(), nil
}

func ArrayFromJs[T any](jsValue js.Value, converter func(js.Value) (T, error)) ([]T, error) {
	if !isJsArray(jsValue) {
		return nil, TypeMismatchError[[]T](jsValue)
	}

	length := jsValue.Get("length").Int()
	result := make([]T, 0, length)

	jsValue.Call("forEach", js.FuncOf(func(this js.Value, args []js.Value) any {
		item := args[0]
		itemAsGo, _ := converter(item) // TODO error handling
		result = append(result, itemAsGo)
		return nil
	}))

	return result, nil
}

func isJsArray(jsValue js.Value) bool {
	return js.Global().Get("Array").Call("isArray", jsValue).Bool()
}

// /////////////////// //
// JSFromGo Converters //
// /////////////////// //

func CreateJsClassInstance(name string, args ...any) js.Value { // TODO
	return js.Global().Get(name).New(args...)
}

// Uint8Array

type Uint8Array struct {
	js.Value
}

func Uint8ArrayFromGo(bytes []byte) Uint8Array {
	jsUint8Array := CreateUint8Array(len(bytes))
	js.CopyBytesToJS(jsUint8Array.Value, bytes)
	return jsUint8Array
}

func CreateUint8Array(len int) Uint8Array {
	return Uint8Array{CreateJsClassInstance("Uint8Array", len)}
}

// String

type JSString struct {
	js.Value
}

func StringFromGoBytes(bytes []byte) JSString {
	jsUint8Array := Uint8ArrayFromGo(bytes)
	textDecoder := CreateTextDecoder()
	return textDecoder.Decode(jsUint8Array.Value)
}

func StringFromGoString(value string) JSString {
	return JSString{js.ValueOf(value)}
}

// TextDecoder

type textDecoder struct {
	js.Value
}

func (textDecoder textDecoder) Decode(bytes js.Value /* Uint8Array */) JSString {
	return JSString{textDecoder.Value.Call("decode", bytes)}
}

func CreateTextDecoder() textDecoder {
	return textDecoder{CreateJsClassInstance("TextDecoder")}
}

// Object
func ObjectFromGo[
	T js.Value | js.Func | bool | int | float32 | string | []any | map[string]any,
](goMap map[string]T) js.Value {
	result := make(map[string]any)
	for key, value := range goMap {
		result[key] = js.ValueOf(value)
	}
	return js.ValueOf(result)
}

// ////////////////////// //
// JS Object Constructors //
// ////////////////////// //

// TODO is there a way to use type aliases or embedded types to allow passing a JS type where a JS value is accepted?

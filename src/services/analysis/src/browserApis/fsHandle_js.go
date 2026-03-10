package browserApis

import (
	"fmt"
	"strings"
	"syscall/js"
)

type FSHandleInterface interface {
	JSWrapper
	Name() string
	Path() string
	IsDirectory() bool
	AsDirectoryHandle() (DirectoryHandle, error)
	AsFileHandle() (FileHandle, error)
	JsValue() js.Value
}

type FSHandle struct {
	jsValue    js.Value
	parentPath []string
}

func (handle FSHandle) IsDirectory() bool {
	type FSHandle_Kind string
	const (
		DIRECTORY FSHandle_Kind = "directory"
		FILE      FSHandle_Kind = "file"
	)

	switch FSHandle_Kind(handle.jsValue.Get("kind").String()) {
	case DIRECTORY:
		return true
	case FILE:
		return false
	default:
		fmt.Println("can't get kind")
		panic(TypeMismatchError[FSHandle_Kind](handle.jsValue.Get("kind")))
	}
}

func (handle FSHandle) JsValue() js.Value {
	return handle.jsValue
}

func (handle FSHandle) Name() string {
	return handle.jsValue.Get("name").String()
}

func (handle FSHandle) Path() string {
	return strings.Join(append(handle.parentPath, handle.Name()), "/")
}

func (handle FSHandle) AsDirectoryHandle() (DirectoryHandle, error) {
	if !handle.IsDirectory() {
		return DirectoryHandle{}, TypeMismatchError[DirectoryHandle](handle.jsValue)
	}
	return DirectoryHandle{handle}, nil
}

func (handle FSHandle) AsFileHandle() (FileHandle, error) {
	if handle.IsDirectory() {
		return FileHandle{}, TypeMismatchError[FileHandle](handle.jsValue)
	}
	return FileHandle{handle}, nil
}

func (handle FSHandle) StoreAsGlobalVariable(varName string) {
	js.Global().Set(varName, handle.jsValue)
}

var _ FSHandleInterface = FSHandle{} // Compile-time inheritance check

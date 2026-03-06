package browserApis

import (
	"fmt"
	"syscall/js"
)

func ShowDirectoryPicker() DirectoryHandle {
	fmt.Println("Calling browser directory picker")
	jsDirectoryHandle := js.Global().Call("showDirectoryPicker")
	directoryHandle := DirectoryHandle{jsDirectoryHandle}
	return directoryHandle
}

type DirectoryHandle struct {
	directoryHandle js.Value
}

func (handle DirectoryHandle) Entries() []FSHandle {
	jsHandleIter := handle.directoryHandle.Call("entries")
	jsHandleList := jsHandleIter.Call("toArray")

	jsHandleListLength, err := GetIntFromJsValue(jsHandleList.Get("length"))
	if err != nil {
		return []FSHandle{}
	}

	fmt.Println(jsHandleListLength)
	// handleList := []FSHandle{}
	// for i := range jsHandleListLength
	return []FSHandle{}
}

type FSHandle interface {
}

package browserApis

import (
	"fmt"
	"syscall/js"
)

func ShowDirectoryPicker() chan DirectoryHandle {
	jsDirectoryHandlePromise := js.Global().Call("showDirectoryPicker")
	directoryHandlePromise := Promise[DirectoryHandle]{jsDirectoryHandlePromise}
	return directoryHandlePromise.ToChannel(DirectoryHandleFromJs)
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

func DirectoryHandleFromJs(v js.Value) DirectoryHandle { return DirectoryHandle{v} }

type FSHandle interface {
}

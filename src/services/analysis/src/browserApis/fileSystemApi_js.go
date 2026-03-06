package browserApis

import (
	"fmt"
	"syscall/js"
)

func ShowDirectoryPicker(directoryChannel chan DirectoryHandle) {
	fmt.Println("Calling browser directory picker")
	jsDirectoryHandlePromise := js.Global().Call("showDirectoryPicker")
	jsDirectoryHandlePromise.Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("then")
		directoryChannel <- DirectoryHandle{args[0]}
		js.Global().Set("handle", args[0])
		fmt.Println("set handle")
		return nil
	}))
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

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
	js.Global().Set("iterator", jsHandleIter)
	fmt.Println("set iterator")

	entriesChannel := make(chan FSEntry)

	// TODO Loop
	nextFile := <-Promise[Iterator[FSEntry]]{jsHandleIter.Call("next")}.ToChannel(IteratorFromJs)
	if nextFile.Done() {
		close(entriesChannel)
	} else {
		fsEntry := nextFile.Value(func(v js.Value) FSEntry {
			return FSEntry{v.Get("0").String(), DirectoryHandleFromJs(v.Get("1"))}
		})
		entriesChannel <- fsEntry
	}

	for entry := range entriesChannel {
		fmt.Println(entry.Name)
	}

	// jsHandleList := jsHandleIter.Call("next")

	// jsHandleListLength, err := GetIntFromJsValue(jsHandleList.Get("length"))
	// if err != nil {
	// 	return []FSHandle{}
	// }

	// fmt.Println(jsHandleListLength)
	// handleList := []FSHandle{}
	// for i := range jsHandleListLength
	return []FSHandle{}
}

func (DirectoryHandle) IsDirectory() bool {
	return true
}

func DirectoryHandleFromJs(v js.Value) DirectoryHandle { return DirectoryHandle{v} }

type FSHandle interface {
	IsDirectory() bool
}

type FSEntry struct {
	Name   string
	Handle FSHandle
}

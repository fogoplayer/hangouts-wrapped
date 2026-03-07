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
	loopChannel := make(chan struct{}, 1)
	loopChannel <- struct{}{} // push one item for the equivalent of a do...while loop

	go func() {
		for range loopChannel {
			nextFile := <-Promise[Iterator[FSEntry]]{jsHandleIter.Call("next")}.ToChannel(IteratorFromJs)
			if nextFile.Done() {
				close(loopChannel)
				close(entriesChannel)
			} else {
				fsEntry := nextFile.Value(func(v js.Value) FSEntry {
					return FSEntry{v.Get("0").String(), DirectoryHandleFromJs(v.Get("1"))}
				})
				loopChannel <- struct{}{}
				entriesChannel <- fsEntry
			}
		}
	}()

	entriesList := []FSHandle{}
	for entry := range entriesChannel {
		entriesList = append(entriesList, entry.Handle)
	}

	return entriesList
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

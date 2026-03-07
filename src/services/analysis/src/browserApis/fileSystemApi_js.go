package browserApis

import (
	"fmt"
	"strings"
	"syscall/js"
)

func ShowDirectoryPicker(channels ...chan DirectoryHandle) chan DirectoryHandle {

	jsDirectoryHandlePromise := js.Global().Call("showDirectoryPicker")
	directoryHandlePromise := Promise[DirectoryHandle]{jsDirectoryHandlePromise}
	switch len(channels) {
	case 0:
		return directoryHandlePromise.ToChannel(DirectoryHandleFromJs)
	case 1:
		return directoryHandlePromise.ToChannel(DirectoryHandleFromJs, channels[0])
	default:
		panic(nil)
	}
}

// /////////////// //
// DirectoryHandle //
// /////////////// //
type DirectoryHandle struct {
	FSHandle
}

func (handle DirectoryHandle) Entries() []FSHandle {
	jsHandleIter := handle.jsValue.Call("entries")
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
					return FSEntry{
						v.Get("0").String(),
						FSHandle{v.Get("1"), append(handle.parentPath, handle.Name())},
					}
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

func DirectoryHandleFromJs(value js.Value) DirectoryHandle {
	candidate := DirectoryHandle{FSHandle{value, []string{}}}
	if !candidate.IsDirectory() {
		panic(nil)
	}
	return candidate
}

// //////// //
// FSHandle //
// //////// //
type FSHandle struct {
	jsValue    js.Value
	parentPath []string
}

func (handle FSHandle) IsDirectory() bool {
	FILE := "file"
	DIRECTORY := "directory"

	switch handle.jsValue.Get("kind").String() {
	case DIRECTORY:
		return true
	case FILE:
		return false
	default:
		panic(nil) // TODO not sure what the right value to panic on is here
	}
}

func (handle FSHandle) Name() string {
	return handle.jsValue.Get("name").String()
}

func (handle FSHandle) RelativePath() string {
	return strings.Join(append(handle.parentPath, handle.Name()), ",")
}

// /////// //
// FSEntry //
// /////// //
type FSEntry struct {
	Name   string
	Handle FSHandle
}

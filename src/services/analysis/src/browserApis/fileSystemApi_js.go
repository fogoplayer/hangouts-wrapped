package browserApis

import (
	"fmt"
	"strings"
	"syscall/js"
)

var PathToFSHandle map[string]*FSHandle = make(map[string]*FSHandle)

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
					name := v.Get("1")
					parentPath := append(handle.parentPath, handle.Name())
					return FSEntry{
						v.Get("0").String(),
						FSHandle{name, parentPath},
					}
				})
				PathToFSHandle[fsEntry.Handle.RelativePath()] = &fsEntry.Handle
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
	return FSHandle{value, []string{}}.AsDirectoryHandle()
}

// ////////// //
// FileHandle //
// ////////// //
type FileHandle struct {
	FSHandle
}

func FileHandleFromJs(value js.Value) FileHandle {
	return FSHandle{value, []string{}}.AsFileHandle()
}

func (handle FileHandle) Data() chan []byte {
	js.Global().Set("handle", handle.jsValue)
	jsFile := <-Promise[js.Value]{handle.jsValue.Call("getFile")}.ToChannel(func(v js.Value) js.Value { return v })
	return Promise[[]byte]{jsFile.Call("bytes")}.ToChannel(func(v js.Value) []byte {
		var data []byte
		js.CopyBytesToGo(data, v)
		return data
	})
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
		panic(fmt.Sprintf("%v cannot be coerced into %s", handle.jsValue, "\"directory\"|\"file\""))
	}
}

func (handle FSHandle) Name() string {
	return handle.jsValue.Get("name").String()
}

func (handle FSHandle) RelativePath() string {
	return strings.Join(append(handle.parentPath, handle.Name()), "/")
}

func (handle FSHandle) AsDirectoryHandle() DirectoryHandle {
	if !handle.IsDirectory() {
		TypeMismatchPanic[DirectoryHandle](handle.jsValue)
	}
	return DirectoryHandle{handle}
}

func (handle FSHandle) AsFileHandle() FileHandle {
	if handle.IsDirectory() {
		TypeMismatchPanic[FileHandle](handle.jsValue)
	}
	return FileHandle{handle}
}

// /////// //
// FSEntry //
// /////// //
type FSEntry struct {
	Name   string
	Handle FSHandle
}

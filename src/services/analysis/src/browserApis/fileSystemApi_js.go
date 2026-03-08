package browserApis

import (
	"errors"
	"fmt"
	"strings"
	"syscall/js"

	"zarinloosli.com/hangouts-wrapped/util"
)

func ShowDirectoryPicker(channels ...chan PromiseResult[DirectoryHandle]) chan PromiseResult[DirectoryHandle] {
	jsDirectoryHandlePromise := js.Global().Call("showDirectoryPicker")
	directoryHandlePromise := Promise[DirectoryHandle]{jsDirectoryHandlePromise}
	switch len(channels) {
	case 0:
		return directoryHandlePromise.ToChannel(getJsToDirectoryHandleFunctionForParent([]string{}))
	case 1:
		return directoryHandlePromise.ToChannel(getJsToDirectoryHandleFunctionForParent([]string{}), channels[0])
	default:
		util.WrongNumberOfArgumentsPanic(len(channels))
		return nil
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

	entriesChannel := make(chan FSEntry)
	loopChannel := make(chan struct{}, 1)
	loopChannel <- struct{}{} // push one item for the equivalent of a do...while loop

	go func() { // TODO is this goroutine necessary? Should it be moved inside?
		// I don't think so, but I'm not sure how next works
		for range loopChannel {
			nextFile, _ := Promise[Iterator[FSEntry]]{jsHandleIter.Call("next")}.ValueSync(IteratorFromJs)
			// TODO error handling
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

func (handle DirectoryHandle) GetEntry(name string) (FSHandleInterface, error) {
	parentPath := append(handle.parentPath, handle.Name())
	directoryChannel := Promise[DirectoryHandle]{handle.jsValue.Call("getDirectoryHandle", name)}.
		ToChannel(getJsToDirectoryHandleFunctionForParent(parentPath))
	fileChannel := Promise[FileHandle]{handle.jsValue.Call("getFileHandle", name)}.
		ToChannel(getJsToFileHandleFunctionForParent(parentPath))
	for range 2 {
		select {
		case directoryResult := <-directoryChannel:
			directoryHandle, err := directoryResult.Value()
			if err == nil {
				return directoryHandle, nil
			}
		case fileResult := <-fileChannel:
			fileHandle, err := fileResult.Value()
			if err == nil {
				return fileHandle, nil
			}
		}
	}
	return nil, errors.New("Entry does not exist")
}

func jsToDirectoryHandle(value js.Value, parentPath []string) (DirectoryHandle, error) {
	return FSHandle{value, parentPath}.AsDirectoryHandle()
}

func getJsToDirectoryHandleFunctionForParent(parentPath []string) func(value js.Value) (DirectoryHandle, error) {
	return func(value js.Value) (DirectoryHandle, error) {
		return jsToDirectoryHandle(value, parentPath)
	}
}

// ////////// //
// FileHandle //
// ////////// //
type FileHandle struct {
	FSHandle
}

// TODO pass in a channel and make the whole thing a goRoutine so it return instantaneously?
func (handle FileHandle) Bytes() chan []byte {
	bytesChannel := make(chan []byte)
	go func() {
		jsFile, _ := Promise[js.Value]{handle.jsValue.Call("getFile")}.ValueSync(func(v js.Value) (js.Value, error) { return v, nil })

		bytes, _ := Promise[[]byte]{jsFile.Call("bytes")}.ValueSync(func(v js.Value) ([]byte, error) {
			data := make([]byte, v.Length())
			js.CopyBytesToGo(data, v)
			return data, nil // TODO error handling for copyBytesToGo
		})
		// TODO error handling
		bytesChannel <- bytes
	}()
	return bytesChannel
}

func jsToFileHandle(value js.Value, parentPath []string) (FileHandle, error) {
	return FSHandle{value, parentPath}.AsFileHandle()
}

func getJsToFileHandleFunctionForParent(parentPath []string) func(value js.Value) (FileHandle, error) {
	return func(value js.Value) (FileHandle, error) {
		return jsToFileHandle(value, parentPath)
	}
}

// //////// //
// FSHandle //
// //////// //
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

// /////// //
// FSEntry //
// /////// //
type FSEntry struct {
	Name   string
	Handle FSHandle
}

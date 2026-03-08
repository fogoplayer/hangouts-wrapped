package browserApis

import (
	"errors"
	"syscall/js"
)

type DirectoryHandle struct {
	FSHandle
}

func (handle DirectoryHandle) Entries() []FSHandle {
	jsHandleIter := handle.jsValue.Call("entries")

	entriesChannel := make(chan FSEntry)
	loopChannel := make(chan struct{}, 1)
	loopChannel <- struct{}{} // push one item for the equivalent of a do...while loop

	go func() { // loop in a goroutine so that we can immediately start listening for entries, too
		for range loopChannel {
			// We need to know if we are `Done` to know if we are ending the loop or not... so there's no harm in getting the
			// value synchronously. Plus, we're in a goroutine already
			nextFile, _ := Await(Promise[Iterator[FSEntry]]{
				value:      jsHandleIter.Call("next"),
				jsToGoFunc: IteratorFromJs[FSEntry],
			})
			// no error handling needed - this is a wrapper for a JS async iterator, so it *will* have a `next`` method that
			// *will* return a value

			if nextFile.Done() {
				close(loopChannel)
				close(entriesChannel)
			} else {
				fsEntry := nextFile.Value(func(v js.Value) FSEntry {
					name := v.Get("0").String()
					fileHandle := v.Get("1")
					parentPath := append(handle.parentPath, handle.Name())
					return FSEntry{
						name,
						FSHandle{fileHandle, parentPath},
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
	directoryChannel := Promise[DirectoryHandle]{
		value:      handle.jsValue.Call("getDirectoryHandle", name),
		jsToGoFunc: getDirectoryHandleFromJsFunctionForParent(parentPath),
	}.ToChannel()
	fileChannel := Promise[FileHandle]{
		value:      handle.jsValue.Call("getFileHandle", name),
		jsToGoFunc: getFileHandleFromJsFunctionForParent(parentPath),
	}.ToChannel()
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

func directoryHandleFromJs(value js.Value, parentPath []string) (DirectoryHandle, error) {
	return FSHandle{value, parentPath}.AsDirectoryHandle()
}

func getDirectoryHandleFromJsFunctionForParent(parentPath []string) func(value js.Value) (DirectoryHandle, error) {
	return func(value js.Value) (DirectoryHandle, error) {
		return directoryHandleFromJs(value, parentPath)
	}
}

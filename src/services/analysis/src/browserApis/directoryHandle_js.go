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

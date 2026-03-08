package fsIo

import (
	"fmt"

	"zarinloosli.com/hangouts-wrapped/browserApis"
	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/util"
)

var PathToFSHandle map[string]model.FSAgnosticHandle = make(map[string]model.FSAgnosticHandle)

func ShowDirectoryPicker() {
	go func() {
		jsDirectoryResult := <-browserApis.ShowDirectoryPicker()
		jsDirectoryHandle, err := jsDirectoryResult.Value()
		if err != nil {
			fmt.Println("directory picker cancelled")
			return
		}

		fsHandle := FSHandle{jsDirectoryHandle}
		dirHandle := DirectoryHandle{fsHandle}
		PathToFSHandle[jsDirectoryHandle.Path()] = dirHandle

		model.FilePathsToIngestChannel <- jsDirectoryHandle.Path()
	}()
}

// ////////// //
// FileHandle //
// ////////// //
type FileHandle struct {
	FSHandle
}

func (handle FileHandle) Bytes() chan []byte {
	fileHandle, _ := handle.BrowserHandle.AsFileHandle()
	// TODO error handling
	return fileHandle.Bytes()
}

var _ model.FSAgnosticFileHandle = FileHandle{} // Compile-time inheritance check

// /////////////// //
// DirectoryHandle //
// /////////////// //
type DirectoryHandle struct {
	FSHandle
}

func (handle DirectoryHandle) Entries() []model.FSAgnosticHandle {
	fsEntries := []model.FSAgnosticHandle{}
	directoryHandle, _ := handle.BrowserHandle.AsDirectoryHandle()
	// TODO error handling
	for _, browserEntry := range directoryHandle.Entries() {
		PathToFSHandle[browserEntry.Path()] = FSHandle{browserEntry}
		fsEntries = append(fsEntries, FSHandle{browserEntry})
	}
	return fsEntries
}

func (handle DirectoryHandle) GetEntry(name string) (model.FSAgnosticHandle, error) {
	directoryHandle, _ := handle.BrowserHandle.AsDirectoryHandle()
	// TODO error handling
	entry, err := directoryHandle.GetEntry(name)
	return DirectoryHandle{FSHandle{entry}}, err
}

var _ model.FSAgnosticDirectoryHandle = DirectoryHandle{} // Compile-time inheritance check

// //////// //
// FSHandle //
// //////// //
type FSHandle struct {
	BrowserHandle browserApis.FSHandleInterface
}

func (handle FSHandle) Name() string {
	return handle.BrowserHandle.Name()
}

func (handle FSHandle) Path() string {
	return handle.BrowserHandle.Path()
}

func (handle FSHandle) IsDirectory() bool {
	return handle.BrowserHandle.IsDirectory()
}

func (handle FSHandle) AsDirectoryHandle() (model.FSAgnosticDirectoryHandle, error) {
	if handle.IsDirectory() {
		return DirectoryHandle{handle}, nil
	} else {
		return DirectoryHandle{}, util.CreateUnableToCastFromAToBError[FSHandle, DirectoryHandle](handle)
	}
}

func (handle FSHandle) AsFileHandle() (model.FSAgnosticFileHandle, error) {
	if !handle.IsDirectory() {
		return FileHandle{handle}, nil
	} else {
		return FileHandle{}, util.CreateUnableToCastFromAToBError[FSHandle, FileHandle](handle)
	}
}

var _ model.FSAgnosticHandle = FSHandle{} // Compile-time inheritance check

func GetFSHandleFromPath(path string) model.FSAgnosticHandle {
	return PathToFSHandle[path]
}

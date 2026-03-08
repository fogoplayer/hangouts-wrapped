package fsIo

import (
	"fmt"
	"path/filepath"

	"zarinloosli.com/hangouts-wrapped/browserApis"
	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/model/jsonSchema"
	"zarinloosli.com/hangouts-wrapped/util"
)

var PathToFSHandle map[string]model.FSAgnosticHandle = make(map[string]model.FSAgnosticHandle)

func ShowDirectoryPicker() {
	go func() {
		jsDirectoryResult := <-browserApis.ShowDirectoryPicker()
		jsDirectoryHandle, err := jsDirectoryResult.Value()
		if err != nil {
			return
		}

		fsHandle := FSHandle{jsDirectoryHandle}
		PathToFSHandle[jsDirectoryHandle.Path()] = DirectoryHandle{fsHandle}

		fmt.Println("Sending to channel")
		model.FilePathsToIngestChannel <- jsDirectoryHandle.Path()
	}()
}

func IngestDirectory(
	path string,
	userInfoJsonChannel chan<- jsonSchema.UserInfo_JsonSchema,
	groupInfoJsonChannel chan<- jsonSchema.GroupInfo_JsonSchema,
	messagesJsonChannel chan<- jsonSchema.Messages_JsonSchema,
) error {
	fsHandle := PathToFSHandle[path]
	if directoryHandle, err := fsHandle.AsDirectoryHandle(); err == nil {
		for _, v := range directoryHandle.Entries() {
			IngestDirectory(v.Path(), userInfoJsonChannel, groupInfoJsonChannel, messagesJsonChannel)
		}
	} else if fileHandle, err := fsHandle.AsFileHandle(); err == nil {
		if filepath.Ext(fileHandle.Name()) == ".json" {
			go func() {
				model.BytesChannel <- <-fileHandle.Bytes()
			}()
		}
	}
	return nil
}

// ////////// //
// FileHandle //
// ////////// //
type FileHandle struct {
	FSHandle
}

func (handle FileHandle) Bytes() chan []byte {
	return handle.browserHandle.AsFileHandle().Bytes()
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
	for _, browserEntry := range handle.browserHandle.AsDirectoryHandle().Entries() {
		PathToFSHandle[browserEntry.Path()] = FSHandle{browserEntry}
		fsEntries = append(fsEntries, FSHandle{browserEntry})
	}
	return fsEntries
}

func (handle DirectoryHandle) GetEntry(name string) (model.FSAgnosticHandle, error) {
	entry, err := handle.browserHandle.AsDirectoryHandle().GetEntry(name)
	return DirectoryHandle{FSHandle{entry}}, err
}

var _ model.FSAgnosticDirectoryHandle = DirectoryHandle{} // Compile-time inheritance check

// //////// //
// FSHandle //
// //////// //
type FSHandle struct {
	browserHandle browserApis.FSHandleInterface
}

func (handle FSHandle) Name() string {
	return handle.browserHandle.Name()
}

func (handle FSHandle) Path() string {
	return handle.browserHandle.Path()
}

func (handle FSHandle) IsDirectory() bool {
	return handle.browserHandle.IsDirectory()
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

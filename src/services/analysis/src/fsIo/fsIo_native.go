//go:build !(js && wasm)

package fsIo

import (
	"os"
	"path/filepath"

	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/util"
)

var GROUPS_DIRECTORY_NAME string = "Groups"
var USER_DATA_DIRECTORY_NAME string = "Users"

func ShowDirectoryPicker(channels ...chan DirectoryHandle) chan DirectoryHandle {
	// TODO implement
	return channels[0]
}

func getDirectoryContentsPaths(directoryPath string) ([]string, error) {
	EMPTY := []string{}

	contents, err := os.ReadDir(directoryPath)
	if err != nil {
		return EMPTY, err
	}

	filePaths := []string{}
	for _, file := range contents {
		filePaths = append(filePaths, filepath.Join(directoryPath, file.Name()))
	}
	return filePaths, nil
}

// ////////// //
// FileHandle //
// ////////// //
type FileHandle struct {
	FSHandle
}

func (handle FileHandle) Bytes() chan []byte {
	userFileBytes, _ := os.ReadFile(handle.path)
	// TODO handle error

	byteChannel := make(chan []byte, 1)
	byteChannel <- userFileBytes
	return byteChannel
}

var _ model.FSAgnosticFileHandle = FileHandle{} // Compile-time inheritance check

// /////////////// //
// DirectoryHandle //
// /////////////// //
type DirectoryHandle struct {
	FSHandle
}

func (handle DirectoryHandle) Entries() []model.FSAgnosticHandle {
	contents, _ := getDirectoryContentsPaths(handle.path)
	// TODO handle error
	result := []model.FSAgnosticHandle{}
	for _, entry := range contents {
		result = append(result, FSHandle{entry})
	}
	return result
}

func (handle DirectoryHandle) GetEntry(name string) (model.FSAgnosticHandle, error) {
	entryPath := filepath.Join(handle.Path(), name)

	_, err := os.Stat(entryPath)
	if err != nil {
		return nil, err
	}

	return FSHandle{entryPath}, nil
}

var _ model.FSAgnosticDirectoryHandle = DirectoryHandle{} // Compile-time inheritance check

// //////// //
// FSHandle //
// //////// //
type FSHandle struct {
	path string
}

func (handle FSHandle) Name() string {
	return filepath.Base(handle.path)
}

func (handle FSHandle) Path() string {
	return handle.path
}

func (handle FSHandle) IsDirectory() bool {
	fileInfo, err := os.Stat(handle.path)
	if err != nil {
		panic(err)
	}

	return fileInfo.IsDir()
}

func (handle FSHandle) AsDirectoryHandle() (model.FSAgnosticDirectoryHandle, error) {
	if !handle.IsDirectory() {
		return nil, util.CreateUnableToCastFromAToBError[FSHandle, model.FSAgnosticDirectoryHandle](handle)
	}
	return DirectoryHandle{handle}, nil
}

func (handle FSHandle) AsFileHandle() (model.FSAgnosticFileHandle, error) {
	if handle.IsDirectory() {
		return nil, util.CreateUnableToCastFromAToBError[FSHandle, model.FSAgnosticFileHandle](handle)
	}
	return FileHandle{handle}, nil
}

func GetFSHandleFromPath(path string) FSHandle {
	return FSHandle{path}
}

var _ model.FSAgnosticHandle = FSHandle{} // Compile-time inheritance check

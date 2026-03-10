//go:build !(js && wasm)

package fsIo

import (
	"os"
	"path/filepath"

	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/util"
)

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
	util.PanicIfError(err)

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

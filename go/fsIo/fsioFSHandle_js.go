package fsIo

import (
	"zarinloosli.com/hangouts-wrapped/browserApis"
	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/util"
)

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

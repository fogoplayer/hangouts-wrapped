package fsIo

import (
	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/util"
)

// ////////// //
// FileHandle //
// ////////// //
type FileHandle struct {
	FSHandle
}

func (handle FileHandle) Bytes() chan []byte {
	fileHandle, err := handle.BrowserHandle.AsFileHandle()
	util.PanicIfError(err)
	return fileHandle.Bytes()
}

var _ model.FSAgnosticFileHandle = FileHandle{} // Compile-time inheritance check

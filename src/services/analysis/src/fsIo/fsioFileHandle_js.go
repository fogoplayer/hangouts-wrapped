package fsIo

import "zarinloosli.com/hangouts-wrapped/model"

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

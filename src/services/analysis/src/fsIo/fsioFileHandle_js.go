package fsIo

import "zarinloosli.com/hangouts-wrapped/model"

// ////////// //
// FileHandle //
// ////////// //
type FileHandle struct {
	FSHandle
}

func (handle FileHandle) Bytes() chan []byte {
	fileHandle, err := handle.BrowserHandle.AsFileHandle()
	if err != nil {
		panic(err)
	}
	return fileHandle.Bytes()
}

var _ model.FSAgnosticFileHandle = FileHandle{} // Compile-time inheritance check

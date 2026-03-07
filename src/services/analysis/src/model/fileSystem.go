package model

type FSAgnosticHandle interface {
	Name() string
	Path() string
	IsDirectory() bool
	AsDirectoryHandle() (FSAgnosticDirectoryHandle, error)
	AsFileHandle() (FSAgnosticFileHandle, error)
}

type FSAgnosticDirectoryHandle interface {
	FSAgnosticHandle
	Entries() []FSAgnosticHandle
}

type FSAgnosticFileHandle interface {
	FSAgnosticHandle
	// reading bytes in WASM is an async operation
	// for wasm, we convert the promise to a channel
	// for native, we return a pre-populated buffered channel
	Bytes() chan []byte
}

package model

type ChatDirectoryHandle struct {
	DirectoryHandle FSAgnosticDirectoryHandle
	Messages        chan []byte
	GroupInfo       chan []byte
}

// Passthrough functions
//
// You can't embed an interface, and the concrete implementations are all in a higher layer
// so we give it a property that conforms to that interface and manually hook up all of the
// callers

func (handle ChatDirectoryHandle) AsDirectoryHandle() (FSAgnosticDirectoryHandle, error) {
	return handle.DirectoryHandle.AsDirectoryHandle()
}

func (handle ChatDirectoryHandle) AsFileHandle() (FSAgnosticFileHandle, error) {
	return handle.DirectoryHandle.AsFileHandle()
}

func (handle ChatDirectoryHandle) Entries() []FSAgnosticHandle {
	return handle.DirectoryHandle.Entries()
}

func (handle ChatDirectoryHandle) GetEntry(name string) (FSAgnosticHandle, error) {
	return handle.DirectoryHandle.GetEntry(name)
}

func (handle ChatDirectoryHandle) IsDirectory() bool {
	return handle.DirectoryHandle.IsDirectory()
}

func (handle ChatDirectoryHandle) Name() string {
	return handle.DirectoryHandle.Name()
}

func (handle ChatDirectoryHandle) Path() string {
	return handle.DirectoryHandle.Path()
}

var _ FSAgnosticDirectoryHandle = ChatDirectoryHandle{} // Compile-time inheritance check

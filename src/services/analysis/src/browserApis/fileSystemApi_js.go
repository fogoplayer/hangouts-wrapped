package browserApis

import (
	"syscall/js"

	"zarinloosli.com/hangouts-wrapped/util"
)

func ShowDirectoryPicker(channels ...chan PromiseResult[DirectoryHandle]) chan PromiseResult[DirectoryHandle] {
	jsDirectoryHandlePromise := js.Global().Call("showDirectoryPicker")
	directoryHandlePromise := Promise[DirectoryHandle]{
		jsDirectoryHandlePromise,
		getDirectoryHandleFromJsFunctionForParent([]string{}),
	}
	switch len(channels) {
	case 0:
		return directoryHandlePromise.ToChannel()
	case 1:
		return directoryHandlePromise.ToChannel(channels[0])
	default:
		util.WrongNumberOfArgumentsPanic(len(channels))
		return nil
	}
}

// /////// //
// FSEntry //
// /////// //
type FSEntry struct {
	Name   string
	Handle FSHandle
}

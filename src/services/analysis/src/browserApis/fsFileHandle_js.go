package browserApis

import "syscall/js"

type FileHandle struct {
	FSHandle
}

// TODO pass in a channel and make the whole thing a goRoutine so it return instantaneously?
func (handle FileHandle) Bytes() chan []byte {
	bytesChannel := make(chan []byte)
	go func() {
		jsFile, _ := Await(Promise[js.Value]{
			handle.jsValue.Call("getFile"),
			JsFromJsReturnValueUnchanged,
		})

		bytes, _ := Await(
			Promise[[]byte]{
				jsFile.Call("bytes"),
				// TODO pull out function
				ByteArrayFromJs,
			},
		)

		// TODO error handling
		bytesChannel <- bytes
	}()
	return bytesChannel
}

func fileHandleFromJs(value js.Value, parentPath []string) (FileHandle, error) {
	return FSHandle{value, parentPath}.AsFileHandle()
}

func getFileHandleFromJsFunctionForParent(parentPath []string) func(value js.Value) (FileHandle, error) {
	return func(value js.Value) (FileHandle, error) {
		return fileHandleFromJs(value, parentPath)
	}
}

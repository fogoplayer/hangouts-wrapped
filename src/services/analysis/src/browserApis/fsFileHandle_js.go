package browserApis

import "syscall/js"

type FileHandle struct {
	FSHandle
}

func (handle FileHandle) Bytes() chan []byte {
	bytesChannel := make(chan []byte)
	go func() {
		jsFile, _ := Await(Promise[js.Value]{
			value:      handle.jsValue.Call("getFile"),
			jsToGoFunc: JsFromJsReturnValueUnchanged,
		})

		bytes, err := Await(
			Promise[[]byte]{
				value:      jsFile.Call("bytes"),
				jsToGoFunc: ByteArrayFromJs,
			},
		)
		if err != nil {
			panic(TypeMismatchError[FileHandle](handle.JsValue()))
		}

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

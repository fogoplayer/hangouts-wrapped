package browserApis

import "syscall/js"

type FileHandle struct {
	FSHandle
}

// TODO pass in a channel and make the whole thing a goRoutine so it return instantaneously?
func (handle FileHandle) Bytes() chan []byte {
	bytesChannel := make(chan []byte)
	go func() {
		jsFile, _ := Promise[js.Value]{handle.jsValue.Call("getFile")}.ValueSync(func(v js.Value) (js.Value, error) { return v, nil })

		bytes, _ := Promise[[]byte]{jsFile.Call("bytes")}.ValueSync(func(v js.Value) ([]byte, error) {
			data := make([]byte, v.Length())
			js.CopyBytesToGo(data, v)
			return data, nil // TODO error handling for copyBytesToGo
		})
		// TODO error handling
		bytesChannel <- bytes
	}()
	return bytesChannel
}

func jsToFileHandle(value js.Value, parentPath []string) (FileHandle, error) {
	return FSHandle{value, parentPath}.AsFileHandle()
}

func getJsToFileHandleFunctionForParent(parentPath []string) func(value js.Value) (FileHandle, error) {
	return func(value js.Value) (FileHandle, error) {
		return jsToFileHandle(value, parentPath)
	}
}

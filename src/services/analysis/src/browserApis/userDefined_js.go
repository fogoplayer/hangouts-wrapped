package browserApis

import (
	"syscall/js"

	"zarinloosli.com/hangouts-wrapped/model"
)

func GenerateSchema(fileHandle model.FSAgnosticFileHandle) {
	go func() {
		json := StringFromGoBytes(<-fileHandle.Bytes())
		key := StringFromGoString(fileHandle.Name())
		documentJson(json, key)
	}()
}

func documentJson(json JSString, key JSString) {
	js.Global().Call("documentJson", json.Value, key.Value)
}

package jsInterface

import (
	"syscall/js"

	"zarinloosli.com/hangouts-wrapped/browserApis"
	"zarinloosli.com/hangouts-wrapped/model"
)

func Initialize() {
	js.Global().Set("showWasmDirectoryPicker", showDirectoryPicker)
}

var showDirectoryPicker js.Func = js.FuncOf(func(this js.Value, args []js.Value) any {
	browserApis.ShowDirectoryPicker(model.ChatDataDirectoryChannel)
	return nil
})

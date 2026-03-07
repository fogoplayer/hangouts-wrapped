package jsInterface

import (
	"syscall/js"

	"zarinloosli.com/hangouts-wrapped/fsIo"
)

func Initialize() {
	js.Global().Set("showWasmDirectoryPicker", showDirectoryPicker)
}

var showDirectoryPicker js.Func = js.FuncOf(func(this js.Value, args []js.Value) any {
	fsIo.ShowDirectoryPicker()
	return nil
})

package jsInterface

import (
	"syscall/js"

	"zarinloosli.com/hangouts-wrapped/browserApis"
	"zarinloosli.com/hangouts-wrapped/fsIo"
	"zarinloosli.com/hangouts-wrapped/model"
)

func Initialize() {
	js.Global().Set("showWasmDirectoryPicker", showDirectoryPicker)
	js.Global().Set("getIngestStats", getIngestStats)
}

var showDirectoryPicker js.Func = js.FuncOf(func(this js.Value, args []js.Value) any {
	fsIo.ShowDirectoryPicker()
	return nil
})

var getIngestStats js.Func = js.FuncOf(func(this js.Value, args []js.Value) any {
	return browserApis.ObjectFromGo(model.IngestStats)
})

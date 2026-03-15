package jsInterface

import (
	"syscall/js"

	"zarinloosli.com/hangouts-wrapped/browserApis"
	"zarinloosli.com/hangouts-wrapped/fsIo"
	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/util"
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
	jsReadyMap := util.MapMap(
		model.GetIngestStats(),
		func(k model.IngestStatsKey, v int) (string, int) {
			return string(k), v
		})
	return browserApis.ObjectFromGo(jsReadyMap)
})

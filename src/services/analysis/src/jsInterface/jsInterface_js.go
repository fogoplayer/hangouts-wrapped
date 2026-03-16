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
	convertToStringKeys := func(k model.IngestStatsKey, v int) (string, int) {
		return string(k), v
	}

	jsReadyMap := util.MapMap(model.GetIngestStats(), convertToStringKeys)
	jsObject := browserApis.ObjectFromGo(jsReadyMap)
	jsObject.Set("toString", js.FuncOf(func(this js.Value, args []js.Value) any {
		// TODO this only works because we mark the return value as Readonly
		// Might be worth using code gen to get a more robus solution
		return js.ValueOf(model.GetIngestStats().String())
	}))
	return jsObject
})

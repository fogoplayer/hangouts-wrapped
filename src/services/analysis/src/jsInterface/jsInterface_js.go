package jsInterface

import (
	"syscall/js"

	"zarinloosli.com/hangouts-wrapped/browserApis"
	"zarinloosli.com/hangouts-wrapped/fsIo"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/util"
)

// The use of IIFEs in this file provide type safety. The return type of each function is documented in analysis.d.ts.
// If a function return type is changed here, it should also be changed there

func Initialize() {
	js.Global().Set("showWasmDirectoryPicker", showDirectoryPicker)
	js.Global().Set("getIngestStats", getIngestStats)
	js.Global().Set("getApplicationPhase", getApplicationPhase)
}

var showDirectoryPicker js.Func = js.FuncOf(func(this js.Value, args []js.Value) any {
	func() {
		fsIo.ShowDirectoryPicker()
	}()
	return nil
})

var getIngestStats js.Func = js.FuncOf(func(this js.Value, args []js.Value) any {
	jsReadyMap := func() map[string]int {
		convertToStringKeys := func(k state.IngestStatsKey, v int) (string, int) {
			return string(k), v
		}

		return util.MapMap(state.GetIngestStats(), convertToStringKeys)
	}()
	jsObject := browserApis.ObjectFromGo(jsReadyMap)
	jsObject.Set("toString", js.FuncOf(func(this js.Value, args []js.Value) any {
		// TODO this only works because we mark the return value as Readonly
		// Might be worth using code gen to get a more robust solution
		return js.ValueOf(state.GetIngestStats().String())
	}))
	return jsObject
})

var getApplicationPhase js.Func = js.FuncOf(func(this js.Value, args []js.Value) any {
	return js.ValueOf(map[string]any{
		"value": string(state.ApplicationPhase.Value()),
		"onChange": js.FuncOf(func(this js.Value, args []js.Value) any {
			for _, arg := range args {
				isFunction := arg.InstanceOf(js.Global().Get("Function"))
				if isFunction {
					state.ApplicationPhase.OnChange(func() {
						arg.Call("call")
					})
				}
			}
			return nil
		}),
	})
})

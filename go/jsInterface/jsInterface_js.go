package jsInterface

import (
	"syscall/js"

	"zarinloosli.com/hangouts-wrapped/browserApis"
	"zarinloosli.com/hangouts-wrapped/fsIo"
	"zarinloosli.com/hangouts-wrapped/model/reports"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/state/stats"
	"zarinloosli.com/hangouts-wrapped/subroutines"
	"zarinloosli.com/hangouts-wrapped/subroutines/filters"
	"zarinloosli.com/hangouts-wrapped/util"
)

// The use of IIFEs in this file provide type safety. The return type of each function is documented in analysis.d.ts.
// If a function return type is changed here, it should also be changed there
// TODO finish this project

// TODO aapparently you don't need to ValueOf return values

func Initialize() {
	js.Global().Set("showWasmDirectoryPicker", showDirectoryPicker)
	js.Global().Set("getIngestStats", getIngestStats)
	js.Global().Set("getApplicationPhase", getApplicationPhase)
	js.Global().Set("getReportsList", getReportsList)
	js.Global().Set("runReport", runReport)
	js.Global().Set("getStableChatsList", getStableChatsList)
	js.Global().Set("setChatFilter", setChatFilter)
}

var showDirectoryPicker js.Func = js.FuncOf(func(this js.Value, args []js.Value) any {
	func() {
		fsIo.ShowDirectoryPicker()
	}()
	return nil
})

var getIngestStats js.Func = js.FuncOf(func(this js.Value, args []js.Value) any {
	jsReadyMap := func() map[string]int {
		convertToStringKeys := func(k stats.IngestStatsKey, v int) (string, int) {
			return string(k), v
		}

		return util.MapMap(stats.GetIngestStats(), convertToStringKeys)
	}()
	jsObject := browserApis.ObjectFromGo(jsReadyMap)
	jsObject.Set("toString", js.FuncOf(func(this js.Value, args []js.Value) any {
		// TODO this only works because we mark the return value as Readonly
		// Might be worth using code gen to get a more robust solution
		return js.ValueOf(stats.GetIngestStats().String())
	}))
	return jsObject
})

var getApplicationPhase js.Func = js.FuncOf(func(this js.Value, args []js.Value) any {
	// TODO don't recreate object each time
	return js.ValueOf(map[string]any{
		"value": string(state.ApplicationPhase.Value()),
		"onChange": js.FuncOf(func(this js.Value, args []js.Value) any {
			for _, arg := range args {
				isFunction := arg.Type() == js.TypeFunction
				if isFunction {
					state.ApplicationPhase.OnChange(func(state.ApplicationPhaseType) {
						arg.Call("call")
					})
				}
			}
			return nil
		}),
	})
})

var getReportsList js.Func = js.FuncOf(func(this js.Value, args []js.Value) any {
	reportsList := reports.GetReportDescriptionsAsList()
	return util.ListMap(reportsList, func(el string) any { return el }) // TODO we have a standard function for this now
})

var setChatFilter js.Func = js.FuncOf(func(this js.Value, args []js.Value) any {
	selectedChatIndexes := util.ListMap(args, func(jsValue js.Value) int {
		num, err := browserApis.IntFromJs(jsValue)
		if err != nil {
			panic(err)
		}
		return num
	})

	filters.SetChatFilterFromIndexes(selectedChatIndexes)
	return nil
})

var getStableChatsList = js.FuncOf(func(this js.Value, args []js.Value) any {
	return util.ListMap(state.GetStableChatNamesList(), util.ToAny)
})

var runReport js.Func = js.FuncOf(func(this js.Value, args []js.Value) any {
	reportEnum, _ := browserApis.IntFromJs(args[0])
	// TODO error handling
	subroutines.SelectReport(reports.ReportName(reportEnum))
	results := subroutines.GetResults()

	return results.ToJsReadyMap()
})

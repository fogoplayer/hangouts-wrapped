package jsInterface

import (
	"fmt"
	"syscall/js"

	"zarinloosli.com/hangouts-wrapped/browserApis"
)

func Initialize() {
	js.Global().Set("ingestDirectory", ingestDirectory)
	fmt.Println("Set ingestDirectory")
}

var ingestDirectory js.Func = js.FuncOf(func(this js.Value, args []js.Value) any {
	channel := browserApis.ShowDirectoryPicker()
	go func() {
		directoryEntry := <-channel
		for _, v := range directoryEntry.Entries() {
			fmt.Println(v.Name(), v.IsDirectory(), v)
			// js.Global().Set("handle_"+strconv.Itoa(i), v)
		}
	}()
	return nil
})

package jsInterface

import (
	"fmt"
	"syscall/js"
)

func Initialize() {
	js.Global().Set("analysis", GetDirectory())
}

func GetDirectory() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println(args)
		return nil
	})
}

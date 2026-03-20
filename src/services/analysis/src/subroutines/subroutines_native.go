//go:build !(js && wasm)

package subroutines

import (
	"fmt"
	"runtime"
	"time"

	"zarinloosli.com/hangouts-wrapped/state"
)

func Setup() {}

func WhileIngesting() {
	if runtime.GOOS != "js" {
		go func() {
			for state.ApplicationPhase.Value() == state.Ingesting {
				time.Sleep(time.Millisecond * 100)
				fmt.Println(state.GetIngestStats())
			}
		}()
	}
}

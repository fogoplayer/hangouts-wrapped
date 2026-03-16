//go:build !(js && wasm)

package subroutines

import (
	"fmt"
	"time"

	"zarinloosli.com/hangouts-wrapped/state"
)

func Setup() {}

func WhileIngesting() {
	progressPrintingWaitGroup.Go(func() {
		for {
			// do...
			time.Sleep(time.Millisecond * 100)
			fmt.Println(state.GetIngestStats())
			// ..while
			if state.ApplicationPhase.Value() != state.Ingesting {
				return
			}
		}
	})
}

//go:build !(js && wasm)

package subroutines

import (
	"fmt"
	"time"

	"zarinloosli.com/hangouts-wrapped/state"
)

func Setup() {}

func WhileIngesting() {
	go func() {
		for state.ApplicationPhase.Value() == state.Ingesting {
			time.Sleep(time.Millisecond * 100)
			fmt.Println(state.GetIngestStats())
		}
	}()
}

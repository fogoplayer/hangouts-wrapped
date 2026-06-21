//go:build !(js && wasm)

package filters

import (
	"time"

	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/userInteractionIo"
	"zarinloosli.com/hangouts-wrapped/util"
)

func SetLowerDateFilter() {
	date := userInteractionIo.Prompt("Date that messages must be sent after to appear in results (MM/DD/YYYY)", parseDateInput)
	state.MinDateFilter.Set(date)
}

func SetUpperDateFilter() {
	date := userInteractionIo.Prompt("Date that messages must be sent before to appear in results (MM/DD/YYYY)", parseDateInput)
	state.MinDateFilter.Set(date)
}

func parseDateInput(input string) (time.Time, error) {
	return time.ParseInLocation(util.MM_DD_YYYY, input, time.Local)
}

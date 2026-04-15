//go:build !(js && wasm)

package subroutines

import (
	"fmt"
	"time"

	"zarinloosli.com/hangouts-wrapped/model/reports"
	reportoutputs "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/userInteractionIo"
	"zarinloosli.com/hangouts-wrapped/util"
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

func PromptForReport() reports.ReportName { // TODO is this the right package for this function?
	keys := util.GetMapKeys(reports.ReportDescriptions)
	values := util.GetMapVals(reports.ReportDescriptions)

	selection := userInteractionIo.Prompt("Choose a report by typing a number:", values)
	if !(selection >= 0 && selection < len(keys)) {
		panic(fmt.Errorf("Prompting for report returned an invalid value: %d", selection))
	}
	selectedReport := keys[selection]
	return selectedReport
}

func OutputReport(results reportoutputs.ReportOutputInterface) {
	fmt.Println(results.String())
}

//go:build !(js && wasm)

package subroutines

import (
	"fmt"
	"time"

	"zarinloosli.com/hangouts-wrapped/model/reports"
	reportoutputs "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/userInteractionIo"
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
	values := reports.GetReportDescriptionsAsList()
	selection := userInteractionIo.Prompt("Choose a report by typing a number:", values)
	return reports.ReportName(selection)
}

func OutputReport(results reportoutputs.ReportOutputInterface) {
	// jsReadyMap := results.ToJsReadyMap()
	// util.UseVar(jsReadyMap)
	fmt.Println(results.String())
}

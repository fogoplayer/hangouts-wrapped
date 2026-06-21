//go:build !(js && wasm)

package subroutines

import (
	"fmt"
	"time"

	"zarinloosli.com/hangouts-wrapped/model/reports"
	reportoutputs "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/state/stats"
	"zarinloosli.com/hangouts-wrapped/userInteractionIo"
)

func Setup() {}

func WhileIngesting() {
	progressPrintingWaitGroup.Go(func() {
		for {
			// do...
			time.Sleep(time.Millisecond * 100)
			fmt.Println(stats.GetIngestStats())
			// ..while
			if state.ApplicationPhase.Value() != state.Ingesting {
				return
			}
		}
	})
}

func PromptForAction() action {
	selection := userInteractionIo.SelectPrompt("Choose an action:", GetActionDescriptionsAsList())
	return action(selection)
}

func PromptForReport() reports.ReportName {
	values := reports.GetReportDescriptionsAsList()
	selection := userInteractionIo.SelectPrompt("Choose a report by typing a number:", values)
	return reports.ReportName(selection)
}

func OutputReport(results reportoutputs.ReportOutputInterface) {
	// jsReadyMap := results.ToJsReadyMap()
	// util.UseVar(jsReadyMap)
	fmt.Println(results.String())
}

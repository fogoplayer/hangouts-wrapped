package subroutines

import (
	"zarinloosli.com/hangouts-wrapped/model/reports"
)

func WhileIngesting() {}

var reportToRunChannel = make(chan reports.ReportName)

func PromptForReport() reports.ReportName {
	return <-reportToRunChannel
}

func SelectReport(selectedReport reports.ReportName) {
	reportToRunChannel <- selectedReport
}

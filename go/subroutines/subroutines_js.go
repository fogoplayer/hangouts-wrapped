package subroutines

import (
	"zarinloosli.com/hangouts-wrapped/model/reports"
	reportoutputs "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
)

func WhileIngesting() {}

var reportToRunChannel = make(chan reports.ReportName)
var reportResultsChannel = make(chan reportoutputs.ReportOutputInterface)

func PromptForReport() reports.ReportName {
	return <-reportToRunChannel
}

func SelectReport(selectedReport reports.ReportName) {
	reportToRunChannel <- selectedReport
}

func OutputReport(results reportoutputs.ReportOutputInterface) {
	reportResultsChannel <- results
}

func GetResults() reportoutputs.ReportOutputInterface {
	return <-reportResultsChannel
}

func PromptForAction() action {
	return RunReport
}

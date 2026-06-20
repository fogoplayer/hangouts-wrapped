package reports

import (
	. "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
)

func countByPerson() *BarOutput {
	messagesByUser := state.CountMessagesByUser()

	output := CreateBarOutput()
	for user, count := range messagesByUser {
		output.Push(ReportOutputEntry[string, int]{
			Label: user,
			Value: count,
		})
	}

	return &output
}

package reports

import (
	. "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
)

func countByPerson() *BarOutput {
	allChats := state.AllChats.Value()
	messagesByUser := make(map[string]int)
	for _, chat := range allChats {
		for _, message := range chat.Messages.Values() {
			messagesByUser[message.Creator.String()] += 1
		}
	}

	output := CreateBarOutput()
	for user, count := range messagesByUser {
		output.Push(ReportOutputEntry[string, int]{
			Label: user,
			Value: count,
		})
	}

	return &output
}

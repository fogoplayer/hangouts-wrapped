package reports

import (
	. "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
)

func countByChat() *BarOutput {
	allChats := state.GetFilteredChats()

	output := CreateBarOutput()
	for _, chat := range allChats {
		output.Push(ReportOutputEntry[string, int]{
			Label: chat.Name,
			Value: len(chat.Messages.Values()),
		})
	}

	return &output
}

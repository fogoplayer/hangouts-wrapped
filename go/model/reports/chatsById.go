package reports

import (
	. "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
)

func chatsById() *TextOutput {
	allChats := state.GetFilteredChats()

	output := CreateTextOutput()
	for _, chat := range allChats {
		output.Push(ReportOutputEntry[string, string]{ // TODO reports should just have a "push" method
			Label: chat.Name,
			Value: chat.Id,
		})
	}

	return &output
}

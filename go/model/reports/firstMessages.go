package reports

import (
	. "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/util"
)

func firstMessages() *TextOutput {
	messagesByUser := state.GetMessagesByUser()

	output := CreateTextOutput()

	for user, messageList := range messagesByUser {
		if len(messageList) == 0 {
			continue
		}

		messageText := messageList[0].Text_
		util.UseVar(messageText)

		output.Push(ReportOutputEntry[string, string]{
			Label: user,
			Value: messageList[0].Text_,
		})
	}

	return &output
}

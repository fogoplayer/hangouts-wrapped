package reports

import (
	. "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/util"
)

func randomMessage() *TextOutput {
	allChats := state.GetFilteredChats()

	randomChat := *util.RandomFromList(allChats)
	randomMessage := util.RandomFromList(randomChat.Messages.Values())

	output := CreateTextOutput(func(a, b ReportOutputEntry[string, string]) int {
		aOrder := getEntryOrder(a)
		bOrder := getEntryOrder(b)

		return aOrder - bOrder
	})

	output.Push(ReportOutputEntry[string, string]{
		Label: "Text",
		Value: randomMessage.Text_,
	})

	output.Push(ReportOutputEntry[string, string]{
		Label: "Post time",
		Value: randomMessage.CreatedDate_.Format(util.HANGOUTS_LOCAL),
	})

	output.Push(ReportOutputEntry[string, string]{
		Label: "Chat",
		Value: randomChat.Name,
	})

	return &output
}

func getEntryOrder(x ReportOutputEntry[string, string]) int {
	switch x.Label {
	case "Chat": // TODO enumerate these
		return 1
	case "Post time":
		return 2
	case "Text":
		return 3
	default:
		return 999
	}
}

package reports

import (
	"zarinloosli.com/hangouts-wrapped/model/parsed"
	. "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/util"
)

func randomAttachment() *FileOutput {
	allChats := state.GetFilteredChats()

	var randomChat *parsed.Chat
	var fileName string

	for {
		randomChat = util.RandomFromList(allChats)
		messagesWithImages := util.ListFilter(randomChat.Messages.Values(), func(message parsed.Message) bool {
			return len(message.Attached_Files_) != 0
		})

		if len(messagesWithImages) == 0 {
			continue
		}
		randomMessage := util.RandomFromList(messagesWithImages)
		attachment := util.RandomFromList(randomMessage.Attached_Files_)
		fileName = attachment.ExportName
		break
	}

	output := CreateFileOutput()
	output.Push(ReportOutputEntry[string, string]{
		Label: randomChat.Name,
		Value: fileName,
	})

	return &output
}

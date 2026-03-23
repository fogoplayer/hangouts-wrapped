package main

import (
	"zarinloosli.com/hangouts-wrapped/model/reports"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/subroutines"
	"zarinloosli.com/hangouts-wrapped/subroutines/setup"
)

func main() {
	setup.Setup()
	state.ApplicationPhase.Set(state.WaitingForDirectory)

	chatDataDirectory := promptForChatDataDirectory() // TODO move to subroutines directory
	subroutines.IngestChatDirectory(chatDataDirectory)
	subroutines.ParseIngestedFiles()

	state.ApplicationPhase.Set(state.Ingesting)
	subroutines.WhileIngesting()
	state.IngestWaitGroup.Wait()

	state.ApplicationPhase.Set(state.WaitingForReport)
	subroutines.PostIngest()

	for true {
		selectedReport := subroutines.PromptForReport()

		state.ApplicationPhase.Set(state.GeneratingReport)
		reports.RunReport(selectedReport)

		state.ApplicationPhase.Set(state.WaitingForReport)
	}
}

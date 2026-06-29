package main

import (
	"os"

	"zarinloosli.com/hangouts-wrapped/model/reports"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/subroutines"
	"zarinloosli.com/hangouts-wrapped/subroutines/setup"
)

func main() {
	// TODO make the entire thing a loop so users can go back and select a new directory
	setup.Setup()
	state.ApplicationPhase.Set(state.WaitingForDirectory)

	chatDataDirectory := promptForChatDataDirectory() // TODO move to subroutines directory
	subroutines.IngestChatDirectory(chatDataDirectory)
	subroutines.ParseIngestedFiles()

	state.ApplicationPhase.Set(state.Ingesting)
	subroutines.WhileIngesting()
	state.IngestWaitGroup.Wait()

	state.ApplicationPhase.Set(state.WaitingForInput)
	subroutines.PostIngest()

	for true {
		switch subroutines.PromptForAction() {
		case subroutines.RunReport:
			selectedReport := subroutines.PromptForReport()

			state.ApplicationPhase.Set(state.GeneratingReport)
			results := reports.RunReport(selectedReport)
			subroutines.OutputReport(results)
		case subroutines.SetIncludedChats:
			subroutines.SetChatFilter()
		case subroutines.Exit:
			os.Exit(0)
		}
		state.ApplicationPhase.Set(state.WaitingForInput)
	}
}

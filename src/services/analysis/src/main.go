package main

import (
	"fmt"

	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/subroutines"
)

func main() {
	subroutines.Setup()
	state.ApplicationPhase.Set(state.WaitingForDirectory)

	chatDataDirectory := promptForChatDataDirectory()
	subroutines.IngestChatDirectory(chatDataDirectory)
	subroutines.ParseIngestedFiles()

	state.ApplicationPhase.Set(state.Ingesting)
	subroutines.WhileIngesting()
	state.IngestWaitGroup.Wait()

	state.ApplicationPhase.Set(state.WaitingForReport)
	subroutines.PostIngest()
	fmt.Println(state.GetIngestStats())
}

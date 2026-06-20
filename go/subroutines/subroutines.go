package subroutines

import (
	"fmt"
	"sync"

	"zarinloosli.com/hangouts-wrapped/state"
)

var progressPrintingWaitGroup = sync.WaitGroup{}

func PostIngest() {
	close(state.FilePathsToIngestChannel)
	close(state.ChatDirectoryHandleChannel)
	progressPrintingWaitGroup.Wait()
}

type action int

const (
	RunReport action = iota
	SetIncludedChats
	Exit
)

var ActionDescriptions = map[action]string{
	RunReport:        "Run report",
	SetIncludedChats: "Set chat filter",
	Exit:             "Exit",
}

func GetActionDescriptionsAsList() []string {
	result := make([]string, len(ActionDescriptions))
	for reportEnumAsInt := range len(ActionDescriptions) {
		description := ActionDescriptions[action(reportEnumAsInt)]
		if reportEnumAsInt >= 0 && reportEnumAsInt < len(result) {
			result[reportEnumAsInt] = description
		} else {
			panic(fmt.Errorf("%s is an out-of-order enum", description))
		}
	}
	return result
}

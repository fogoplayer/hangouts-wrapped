package reports

import (
	"fmt"

	. "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
)

func countByPerson() BarOutput {
	allChats := state.AllChats.Value()
	messagesByUser := make(map[string]int)
	for _, chat := range allChats {
		for _, message := range chat.Messages.Values() {
			messagesByUser[message.Creator.String()] += 1
		}
	}

	output := CreateBarOutput()
	for range messagesByUser {
		maxCount := 0
		maxUser := ""

		for user, count := range messagesByUser {
			if count > maxCount {
				maxCount = count
				maxUser = user
			}
		}
		delete(messagesByUser, maxUser)

		fmt.Println(maxUser, maxCount)
		// output.values.Push(ReportOutputEntry[int]{maxUser, maxCount})

		// output.Labels = append(output.Labels, user)
		// output.Values = append(output.Values, count)
	}

	return output
}

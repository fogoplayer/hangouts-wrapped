package reports

import (
	"time"

	. "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/util"
)

func countByHour() *LineOutput {
	allChats := state.AllChats.Value()

	output := CreateLineOutput(func(t time.Time) string { return t.Format(util.HOUR_ONLY) })

	countsByHour := make([]int, 24)

	for _, chat := range allChats {
		for _, monthTreeList := range chat.Messages {
			for _, dayTreeList := range monthTreeList {
				for _, hourTreeList := range dayTreeList {
					for hour, minuteTreeList := range hourTreeList {
						countsByHour[hour] += len(minuteTreeList.Values())
					}
				}
			}
		}
	}

	for hour, count := range countsByHour {
		output.Push(ReportOutputEntry[time.Time, int]{
			Label: time.Date(0, 0, 0, hour, 0, 0, 0, time.UTC),
			Value: count,
		})
	}

	return &output
}

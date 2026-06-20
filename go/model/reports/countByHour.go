package reports

import (
	"time"

	. "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/util"
)

func countByHour() *LineOutput {
	countsByHour := state.CountMessagesByHour()

	output := CreateLineOutput(func(t time.Time) string { return t.Format(util.HOUR_ONLY) })

	for hour, count := range countsByHour {
		output.Push(ReportOutputEntry[time.Time, int]{
			Label: time.Date(0, 0, 0, hour, 0, 0, 0, time.UTC),
			Value: count,
		})
	}

	return &output
}

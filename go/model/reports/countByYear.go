package reports

import (
	"slices"
	"time"

	. "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/util"
)

func countByYear() *LineOutput {
	countsByYear := state.CountMessagesByYear()

	years := util.GetMapKeys(countsByYear)
	slices.SortFunc(years, func(i, j time.Time) int {
		if i.Before(j) {
			return -1
		} else if i.After(j) {
			return 1
		} else {
			return 0
		}
	})

	output := CreateLineOutput(func(t time.Time) string {
		return t.Format(util.YEAR_ONLY)
	})

	for _, year := range years {
		output.Push(ReportOutputEntry[time.Time, int]{
			Label: year,
			Value: countsByYear[year],
		})
	}

	return &output
}

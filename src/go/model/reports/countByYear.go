package reports

import (
	"slices"
	"time"

	. "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/util"
)

func countByYear() *LineOutput {
	allChats := state.AllChats.Value()

	output := CreateLineOutput(func(t time.Time) string {
		return t.Format(util.YEAR_ONLY)
	})

	countsByYear := make(map[time.Time]int)

	for _, chat := range allChats {
		for year, monthTreeList := range chat.Messages {
			countsByYear[time.Date(int(year), 1, 1, 0, 0, 0, 0, time.UTC)] += len(monthTreeList.Values())
		}
	}

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

	for _, year := range years {
		output.Push(ReportOutputEntry[time.Time, int]{
			Label: year,
			Value: countsByYear[year],
		})
	}

	return &output
}

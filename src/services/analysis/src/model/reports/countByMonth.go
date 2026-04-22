package reports

import (
	"slices"
	"time"

	. "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/util"
)

func countByMonth() *LineOutput {

	allChats := state.AllChats.Value()

	output := CreateLineOutput(func(t time.Time) string { return t.Format(util.MONTH_YEAR) })

	countsByMonth := make(map[time.Time]int)

	for _, chat := range allChats {
		// util.MapMap(map[]chat.Messages)
		for year, monthTreeList := range chat.Messages {
			for month, dayTreeList := range monthTreeList {
				countsByMonth[time.Date(int(year), time.Month(month), 1, 0, 0, 0, 0, time.UTC)] += len(dayTreeList.Values())
			}
		}
	}

	months := util.GetMapKeys(countsByMonth)
	slices.SortFunc(months, func(i, j time.Time) int {
		if i.Before(j) {
			return -1
		} else if i.After(j) {
			return 1
		} else {
			return 0
		}
	})

	for _, month := range months {
		output.Push(ReportOutputEntry[time.Time, int]{
			Label: month,
			Value: countsByMonth[month],
		})
	}

	return &output
}

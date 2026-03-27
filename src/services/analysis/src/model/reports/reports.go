package reports

import (
	"fmt"

	"zarinloosli.com/hangouts-wrapped/state"
)

type report func() ReportOutput // TODO narrow this

type ReportName int

const (
	CountByPerson ReportName = iota //Post count by person
	// CountByDay reportName // Post count by day (month?)
	CountByMonth
	CountByYear
	CountByHour
	CountByChat        // Post count by chat (which is the biggest in the timeframe)
	RandomMessage      // Get a random message in selection
	RandomImage        // Get a random image in selection
	MostUsedEmoji      // Emoji frequency?
	MostReactions      // Most reacted-to message
	MembershipTimeline // Every time someone joined and left one the selected chats in the selected timeframe
	// Pipe dreams
	WordFrequency // Google Books-style word frequency (Also check for common misspellings?)
	// Line graphs of how often people were posting / emoji were used / words were used over time to see the ebb and flow
	LongestThreads  // Longest thread
	FirstMessages   // Each user's First message
	WordCountByChat // Get the number of words that have been written
	WordCountByUser
)

var Reports = map[ReportName]report{
	CountByPerson: countByPerson,
}

var ReportDescriptions = map[ReportName]string{
	CountByPerson: "Number of messages sent by each user",
}

func GetReportDescriptionsAsList() []string {
	result := make([]string, len(ReportDescriptions))
	for reportEnum, description := range ReportDescriptions {
		reportEnumAsInt := int(reportEnum)
		if reportEnum >= 0 && reportEnumAsInt < len(result) {
			result[reportEnumAsInt] = description
		} else {
			panic(fmt.Errorf("%s is an out-of-order enum", description))
		}
	}
	return result
}

func RunReport(reportName ReportName) any { // TODO what is the format of report results? Some sort of table?
	report := Reports[reportName]
	if report == nil {
		// TODO eventually this should be a panic
		fmt.Printf("%d does not exist (%s)\n", reportName, ReportDescriptions[reportName])
	}

	return report()
}

// /////////// //
// Return Type //
// /////////// //

type ReportKind string

const (
	Bar         ReportKind = "bar"
	Line        ReportKind = "line"
	SingleValue ReportKind = "singlevalue"
)

type ReportOutput struct {
	Kind   ReportKind
	Labels []string
	Values []any
}

func (reportOutput *ReportOutput) ToJsReadyMap() map[string]any {
	return map[string]any{
		"kind":   reportOutput.Kind,
		"labels": reportOutput.Labels,
		"Values": reportOutput.Values,
	}
}

// /////// //
// Reports //
// /////// //

func countByPerson() ReportOutput {
	allChats := state.AllChats.Value()
	messagesByUser := make(map[string]int)
	for _, chat := range allChats {
		for _, message := range chat.Messages.Values() {
			messagesByUser[message.Creator.String()] += 1
		}
	}

	output := ReportOutput{}
	output.Kind = Bar
	for user, count := range messagesByUser {
		output.Labels = append(output.Labels, user)
		output.Values = append(output.Values, count)
	}

	return output
}

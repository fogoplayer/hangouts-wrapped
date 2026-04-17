package reports

import (
	"fmt"

	. "zarinloosli.com/hangouts-wrapped/model/reports/reportOutputs"
)

type report func() ReportOutputInterface // TODO narrow this

type ReportName int

const (
	CountByPerson ReportName = iota //Post count by person
	CountByChat                     // Post count by chat (which is the biggest in the timeframe)
	ChatsById                       // Map chat names to IDs
	CountByDay                      // Post count by day (month?)
	CountByMonth
	CountByYear
	CountByHour
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

var ReportDescriptions = map[ReportName]string{
	CountByPerson: "Number of messages sent by each user",
	CountByChat:   "Number of messages sent in each chat",
	ChatsById:     "Map chat names to IDs",
}

func GetReportDescriptionsAsList() []string {
	result := make([]string, len(ReportDescriptions))
	for reportEnumAsInt := range len(ReportDescriptions) {
		description := ReportDescriptions[ReportName(reportEnumAsInt)]
		if reportEnumAsInt >= 0 && reportEnumAsInt < len(result) {
			result[reportEnumAsInt] = description
		} else {
			panic(fmt.Errorf("%s is an out-of-order enum", description))
		}
	}
	return result
}

func RunReport(reportName ReportName) ReportOutputInterface {
	var output ReportOutputInterface

	switch reportName {
	case CountByPerson:
		return countByPerson()
	case CountByChat:
		return countByChat()
	case ChatsById:
		return chatsById()
	default:
		// TODO eventually this should be a panic
		fmt.Printf("%d does not exist (%s)\n", reportName, ReportDescriptions[reportName])
	}

	return output

}

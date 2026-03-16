package reports

type report func(...[]any) any // TODO narrow this

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

var Reports = map[ReportName]report{}

var ReportDescriptions = map[ReportName]string{
	CountByPerson: "Number of messages sent by each user",
}

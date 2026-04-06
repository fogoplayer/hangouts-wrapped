package reports

import (
	"fmt"
	"math"
	"strings"

	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/util"
)

type report func() ReportOutputInterface // TODO narrow this

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

func RunReport(reportName ReportName) ReportOutputInterface {
	switch reportName {
	case CountByPerson:
		return countByPerson()

	default:
		// TODO eventually this should be a panic
		fmt.Printf("%d does not exist (%s)\n", reportName, ReportDescriptions[reportName])
		return ReportOutput[any]{}
	}
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

type ReportOutputInterface interface {
	Labels() []string
	Values() []any
	String() string
}

type ReportOutput[T any] struct {
	Kind   ReportKind
	values util.Heap[ReportOutputEntry[T]]
}

type ReportOutputEntry[T any] struct {
	Label string
	Value T
}

func (reportOutput ReportOutput[T]) Labels() []string {
	labels := make([]string, 0, reportOutput.values.Len())

	for reportOutput.values.Len() > 0 {
		v := reportOutput.values.Pop()
		labels = append(labels, v.Label)
	}

	return labels
}

func (reportOutput ReportOutput[T]) Values() []any {
	labels := make([]any, 0, reportOutput.values.Len())

	for reportOutput.values.Len() > 0 {
		v := reportOutput.values.Pop()
		labels = append(labels, v.Value)
	}

	return labels
}

func (reportOutput *ReportOutput[T]) ToJsReadyMap() map[string]any {
	return map[string]any{
		"kind":   reportOutput.Kind,
		"labels": reportOutput.Labels(),
		"values": reportOutput.Values(),
	}
}

func (reportOutput ReportOutput[T]) String() string {
	return reportOutput.toString()
}

func (reportOutput *ReportOutput[T]) toString(builders ...*strings.Builder) string {
	var builder *strings.Builder
	if len(builders) > 0 {
		builder = builders[0]
	} else {
		builder = &strings.Builder{}
	}

	if len(reportOutput.Labels()) != len(reportOutput.Values()) {
		panic(fmt.Errorf("Report has %d labels and %d values!", len(reportOutput.Labels()), len(reportOutput.Values())))
	}

	for i := range reportOutput.Labels() {
		fmt.Fprintf(builder, "%s: %s\n", reportOutput.Labels()[i], reportOutput.Values()[i])
	}

	return builder.String()
}

var _ ReportOutputInterface = ReportOutput[any]{} // Compile-time inheritance check

// //////////////// //
// Specific Outputs //
// //////////////// //

type BarOutput struct {
	ReportOutput[int]
}

func (barOutput BarOutput) String() string {
	return barOutput.toString()
}

func (barOutput *BarOutput) toString(builders ...*strings.Builder) string {
	var builder *strings.Builder
	if len(builders) > 0 {
		builder = builders[0]
	} else {
		builder = &strings.Builder{}
	}

	COLUMNS := 40.0
	max := -1.0
	for _, value := range barOutput.Values() {
		valueAsFloat := float64(value.(int)) // TODO simplify casting
		if valueAsFloat > max {
			max = valueAsFloat
		}
	}

	// TODO calling these methods is not stable, find a way to stablize them
	values := barOutput.Values()
	labels := barOutput.Labels()

	for i := range len(labels) {
		fmt.Fprintf(builder, "%s: ", labels[i])
		value := float64(values[i].(int)) // TODO store type somehow
		chars := float64(value) / max * COLUMNS
		roundedChars := int(math.Round(chars))
		for range roundedChars {
			fmt.Fprintf(builder, "%c", 0x2588)
		}
		fmt.Fprintln(builder, "\t", value)
	}
	return barOutput.ReportOutput.toString(builder)
}

func CreateBarOutput() BarOutput {
	return BarOutput{
		ReportOutput[int]{
			Kind: Bar,
			values: util.CreateHeap(func(a, b ReportOutputEntry[int]) int {
				return b.Value - a.Value
			}),
		},
	}
}

var _ ReportOutputInterface = BarOutput{} // Compile-time inheritance check

// /////// //
// Reports //
// /////// //

func countByPerson() BarOutput {
	allChats := state.AllChats.Value()
	messagesByUser := make(map[string]int)
	for _, chat := range allChats {
		for _, message := range chat.Messages.Values() {
			messagesByUser[message.Creator.String()] += 1
		}
	}

	output := CreateBarOutput()
	for user, count := range messagesByUser {
		output.values.Push(ReportOutputEntry[int]{user, count})
		// output.Labels = append(output.Labels, user)
		// output.Values = append(output.Values, count)
	}

	return output
}

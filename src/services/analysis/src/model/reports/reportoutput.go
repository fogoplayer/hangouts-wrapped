package reports

import (
	"fmt"
	"math"
	"strings"

	"zarinloosli.com/hangouts-wrapped/util"
)

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

type ReportOutputEntry[T any] struct {
	Label string
	Value T
}

type ReportOutput[T any] struct {
	Kind   ReportKind
	values util.Heap[ReportOutputEntry[T]]
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

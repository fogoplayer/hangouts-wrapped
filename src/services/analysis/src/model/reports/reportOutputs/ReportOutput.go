package reportoutputs

import (
	"fmt"
	"strings"

	"zarinloosli.com/hangouts-wrapped/util"
)

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

func (reportOutput ReportOutput[T]) TypedValues() []T {
	labels := make([]T, 0, reportOutput.values.Len())

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

var _ ReportOutputInterface = &ReportOutput[any]{} // Compile-time inheritance check

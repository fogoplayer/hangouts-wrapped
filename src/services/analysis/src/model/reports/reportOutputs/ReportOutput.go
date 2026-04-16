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

// TODO delete
// replace all uses with TypedValues
// rename TypedValues to Values
func (reportOutput ReportOutput[T]) Values() []any {
	return util.ListMap(reportOutput.TypedValues(), func(value T) any { return value })
}

func (reportOutput ReportOutput[T]) TypedValues() []T {
	values := make([]T, 0, reportOutput.values.Len())

	for _, v := range reportOutput.values.Values() {
		values = append(values, v.Value)
	}

	return values
}

func (reportOutput *ReportOutput[T]) ToJsReadyMap() map[string]any {
	// for some reason, when these are inlined into the map, the order isn't stable.
	// But it seems to work properly like this
	labels := util.ListMap(reportOutput.Labels(), util.ToAny)
	data := util.ListMap(reportOutput.TypedValues(), util.ToAny)
	return map[string]any{
		"type": string(reportOutput.Kind),
		"data": map[string]any{
			"labels": labels,
			"datasets": []any{map[string]any{
				"data": data,
			}},
		},
	}
}

func (reportOutput *ReportOutput[T]) Push(vals ...ReportOutputEntry[T]) {
	util.ListForEach(vals, func(val ReportOutputEntry[T]) {
		reportOutput.values.Push(val)
	})
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

func CreateReportOutput[T any](comparator func(a, b ReportOutputEntry[T]) int) ReportOutput[T] {
	return ReportOutput[T]{
		Kind:   Bar,
		values: util.CreateHeap(comparator),
	}
}

var _ ReportOutputInterface = &ReportOutput[any]{} // Compile-time inheritance check

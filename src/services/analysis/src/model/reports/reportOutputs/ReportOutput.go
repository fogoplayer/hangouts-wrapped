package reportoutputs

import (
	"fmt"
	"strconv"
	"strings"

	"zarinloosli.com/hangouts-wrapped/util"
)

type ReportOutput[L comparable, V any] struct {
	Kind   ReportKind
	values util.Heap[ReportOutputEntry[L, V]]
}

func (reportOutput ReportOutput[L, V]) Labels() []L {
	labels := make([]L, 0, reportOutput.values.Len())

	for _, v := range reportOutput.values.Values() {
		labels = append(labels, v.Label)
	}

	return labels
}

func (reportOutput ReportOutput[L, V]) LabelsAsStrings() []string {
	return util.ListMap(reportOutput.Labels(), func(label L) string {
		labelAsAny := util.ToAny(label)

		str, isString := labelAsAny.(string)
		if isString {
			return str
		}
		stringer, isStringer := labelAsAny.(fmt.Stringer)
		if isStringer {
			return stringer.String()
		}
		i, isInt := labelAsAny.(int)
		if isInt {
			return strconv.Itoa(i)
		}
		panic(fmt.Errorf("%v is not a string or an int", label))
	})
}

func (reportOutput ReportOutput[L, V]) ValuesAsAny() []any {
	return util.ListMap(reportOutput.Values(), func(value V) any { return value })
}

func (reportOutput ReportOutput[L, V]) Values() []V {
	values := make([]V, 0, reportOutput.values.Len())

	for _, v := range reportOutput.values.Values() {
		values = append(values, v.Value)
	}

	return values
}

func (reportOutput *ReportOutput[L, V]) ToJsReadyMap() map[string]any {
	labels := util.ListMap(reportOutput.Labels(), util.ToAny)
	data := util.ListMap(reportOutput.Values(), util.ToAny)

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

func (reportOutput *ReportOutput[L, V]) Push(vals ...ReportOutputEntry[L, V]) {
	util.ListForEach(vals, func(val ReportOutputEntry[L, V]) {
		reportOutput.values.Push(val)
	})
}

func (reportOutput ReportOutput[L, V]) String() string {
	builder := &strings.Builder{}

	if len(reportOutput.Labels()) != len(reportOutput.Values()) {
		panic(fmt.Errorf("Report has %d labels and %d values!", len(reportOutput.Labels()), len(reportOutput.Values())))
	}

	for i := range reportOutput.Labels() {
		fmt.Fprintf(builder, "%s: %v\n", reportOutput.LabelsAsStrings()[i], reportOutput.Values()[i])
	}

	return builder.String()
}

func CreateReportOutput[L comparable, V any](comparator func(a, b ReportOutputEntry[L, V]) int) ReportOutput[L, V] {
	return ReportOutput[L, V]{
		Kind:   Bar,
		values: util.CreateHeap(comparator),
	}
}

var _ ReportOutputInterface = &ReportOutput[*any, any]{} // Compile-time inheritance check

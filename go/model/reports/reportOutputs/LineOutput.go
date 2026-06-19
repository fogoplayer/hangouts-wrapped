package reportoutputs

import (
	"fmt"
	"math"
	"strings"
	"text/tabwriter"
	"time"

	"zarinloosli.com/hangouts-wrapped/util"
)

type LineOutput struct {
	ReportOutput[time.Time, int]
	keyToString func(time.Time) string
}

// TODO maybe we move all of these ToJsReady methods into a a reportOutputs_js.go file
// and can I alias map[string]any to JSObject without making ValueOf mad?

// TODO hide legend

func (output *LineOutput) LabelsAsStrings() []string {
	return util.ListMap(output.Labels(), output.keyToString)
}

func (output *LineOutput) ToJsReadyMap() map[string]any {
	chartConfig := output.ReportOutput.ToJsReadyMap()
	// format labels for display
	labels := util.ListMap(output.Labels(), output.keyToString)
	chartConfig["data"].(map[string]any)["labels"] = util.ListMap(labels, util.ToAny)
	return chartConfig
}

func (output LineOutput) String() string {
	builder := &strings.Builder{}
	tabWriter := tabwriter.NewWriter(builder, 0, 0, 1, ' ', 0)

	COLUMNS := 40.0
	max := -1.0
	for _, value := range output.Values() {
		valueAsFloat := float64(value)
		if valueAsFloat > max {
			max = valueAsFloat
		}
	}

	values := output.Values()
	labels := output.LabelsAsStrings()

	for i := range len(labels) {
		fmt.Fprintf(tabWriter, "%s:", labels[i])
		fmt.Fprint(tabWriter, "\t")

		value := float64(values[i])
		fmt.Fprint(tabWriter, value)
		fmt.Fprint(tabWriter, "\t")

		chars := float64(value) / max * COLUMNS
		roundedChars := int(math.Round(chars))
		for range roundedChars {
			fmt.Fprintf(tabWriter, "%c", 0x2588)
		}
		fmt.Fprintln(tabWriter)
	}

	tabWriter.Flush()
	return builder.String()
}

func CreateLineOutput(keyToString func(time.Time) string) LineOutput {
	return LineOutput{
		ReportOutput[time.Time, int]{
			Kind:   Line,
			values: util.CreateHeap(CompareLineOutputEntries),
		},
		keyToString,
	}
}

func CompareLineOutputEntries(a, b ReportOutputEntry[time.Time, int]) int {
	if a.Label.Before(b.Label) {
		return -1
	} else if a.Label.After(b.Label) {
		return 1
	} else {
		return 0
	}
}

var _ ReportOutputInterface = &LineOutput{} // Compile-time inheritance check

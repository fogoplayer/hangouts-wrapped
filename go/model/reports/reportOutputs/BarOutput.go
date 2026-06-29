package reportoutputs

import (
	"fmt"
	"math"
	"strings"
	"text/tabwriter"

	"zarinloosli.com/hangouts-wrapped/util"
)

type BarOutput struct {
	ReportOutput[string, int]
}

// TODO maybe we move all of these ToJsReady methods into a a reportOutputs_js.go file
// and can I alias map[string]any to JSObject without making ValueOf mad?

// TODO hide legend

func (barOutput *BarOutput) ToJsReadyMap() map[string]any {
	chartConfig := barOutput.ReportOutput.ToJsReadyMap()
	chartConfig["options"] = map[string]any{ // TODO union options instead of re-implementing base options here
		"indexAxis": "y",
		"plugins": map[string]any{
			"legend": map[string]any{
				"display": false,
			},
		},
	}
	return chartConfig
}

func (barOutput BarOutput) String() string {
	builder := &strings.Builder{}

	tabWriter := tabwriter.NewWriter(builder, 0, 0, 1, ' ', 0)

	COLUMNS := 40.0
	max := -1.0
	for _, value := range barOutput.Values() {
		valueAsFloat := float64(value)
		if valueAsFloat > max {
			max = valueAsFloat
		}
	}

	values := barOutput.Values()
	labels := barOutput.Labels()

	for i := range len(labels) {
		fmt.Fprintf(tabWriter, "%d. %s:", i+1, labels[i])
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

func CreateBarOutput() BarOutput {
	return BarOutput{
		ReportOutput[string, int]{
			Kind:   Bar,
			values: util.CreateHeap(CompareBarOutputEntries),
		},
	}
}

func CompareBarOutputEntries(a, b ReportOutputEntry[string, int]) int {
	return b.Value - a.Value
}

var _ ReportOutputInterface = &BarOutput{} // Compile-time inheritance check

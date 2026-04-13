package reportoutputs

import (
	"fmt"
	"math"
	"strings"

	"zarinloosli.com/hangouts-wrapped/util"
)

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

var _ ReportOutputInterface = &BarOutput{} // Compile-time inheritance check

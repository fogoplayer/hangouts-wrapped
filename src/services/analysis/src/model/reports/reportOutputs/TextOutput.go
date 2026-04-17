package reportoutputs

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"zarinloosli.com/hangouts-wrapped/util"
)

type TextOutput struct {
	ReportOutput[string]
}

func (output *TextOutput) String() string {
	builder := &strings.Builder{}
	tabWriter := tabwriter.NewWriter(builder, 0, 0, 1, ' ', 0)

	if len(output.Labels()) != len(output.Values()) {
		panic(fmt.Errorf("Report has %d labels and %d values!", len(output.Labels()), len(output.Values())))
	}

	for i := range output.Labels() {
		fmt.Fprintf(tabWriter, "%s:\t%s\n", output.Labels()[i], output.TypedValues()[i])
	}

	tabWriter.Flush()
	return builder.String()
}

func CreateTextOutput() TextOutput {
	return TextOutput{
		ReportOutput[string]{
			Kind:   Text,
			values: util.CreateHeap(CompareTextOutputEntries),
		},
	}
}

func CompareTextOutputEntries(a, b ReportOutputEntry[string]) int {
	if a.Value < b.Value {
		return 1
	} else if a.Value > b.Value {
		return -1
	} else {
		return 0
	}
}

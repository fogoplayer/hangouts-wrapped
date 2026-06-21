package reportoutputs

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"zarinloosli.com/hangouts-wrapped/util"
)

type FileOutput struct {
	ReportOutput[string, string]
}

func (output *FileOutput) String() string {
	builder := &strings.Builder{}
	tabWriter := tabwriter.NewWriter(builder, 0, 0, 1, ' ', 0)

	if len(output.Labels()) != len(output.ValuesAsAny()) {
		panic(fmt.Errorf("Report has %d labels and %d values!", len(output.Labels()), len(output.ValuesAsAny())))
	}

	for i := range output.Labels() {
		fmt.Fprintf(tabWriter, "%s:\t%s\n", output.Labels()[i], output.Values()[i])
	}

	tabWriter.Flush()
	return builder.String()
}

func CreateFileOutput(comparators ...func(a, b ReportOutputEntry[string, string]) int) FileOutput {
	comparator := util.ExtractOptionalArgumentWithDefault(comparators, CompareFileOutputEntries)

	return FileOutput{
		ReportOutput[string, string]{
			Kind:   Text,
			values: util.CreateHeap(comparator),
		},
	}
}

func CompareFileOutputEntries(a, b ReportOutputEntry[string, string]) int {
	if a.Label < b.Label {
		return -1
	} else if a.Label > b.Label {
		return 1
	} else {
		return 0
	}
}

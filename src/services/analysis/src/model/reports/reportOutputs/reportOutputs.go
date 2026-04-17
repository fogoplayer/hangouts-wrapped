package reportoutputs

type ReportKind string

const (
	Bar         ReportKind = "bar"
	Line        ReportKind = "line"
	SingleValue ReportKind = "singlevalue"
	Text        ReportKind = "text"
)

type ReportOutputInterface interface {
	LabelStrings() []string
	Values() []any
	String() string
	ToJsReadyMap() map[string]any
}

type ReportOutputEntry[L comparable, V any] struct {
	Label L
	Value V
}

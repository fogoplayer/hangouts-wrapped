package reportoutputs

type ReportKind string

const (
	Bar         ReportKind = "bar"
	Line        ReportKind = "line"
	SingleValue ReportKind = "singlevalue"
	Text        ReportKind = "text"
)

type ReportOutputInterface interface {
	Labels() []string
	Values() []any
	String() string
	ToJsReadyMap() map[string]any
}

type ReportOutputEntry[T any] struct {
	Label string
	Value T
}

package reportoutputs

type ReportKind string

const (
	Bar         ReportKind = "bar"
	Line        ReportKind = "line"
	SingleValue ReportKind = "singlevalue"
)

type ReportOutputInterface interface {
	Labels() []string
	Values() []any
	String() string
}

type ReportOutputEntry[T any] struct {
	Label string
	Value T
}

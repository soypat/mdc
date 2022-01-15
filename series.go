package mdc

import "github.com/hexops/vecty"

type Series interface {
	// Head is the title of the data.
	Head() string
	// Kind
	Kind() DataKind
	AtRow(i int) vecty.MarkupOrChild
}

type DataKind int

const (
	defaultDataKind DataKind = iota
	DataString
	DataNumeric
	DataCheckbox
)

func (dk DataKind) ClassName() (class string) {
	switch dk {
	case DataString, defaultDataKind:
		// no class.
	case DataNumeric:
		class = "mdc-data-table__cell--numeric"
	case DataCheckbox:
		class = "mdc-data-table__header-row-checkbox"
	default:
		panic("unknown DataKind")
	}
	return class
}

// Compile-time check of interface implementation.
var (
	_ Series = (*StringSeries)(nil)
)

type StringSeries struct {
	Label string
	Data  []string
}

func (ss *StringSeries) Head() string   { return ss.Label }
func (ss *StringSeries) Kind() DataKind { return DataString }

func (ss *StringSeries) AtRow(i int) vecty.MarkupOrChild {
	return vecty.Text(ss.Data[i])
}

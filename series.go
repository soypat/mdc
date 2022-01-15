package mdc

import (
	"strconv"

	"github.com/hexops/vecty"
)

// Series interface is meant to be used with the DataTable
// component for displaying structured data.
type Series interface {
	// Head is the title of the data.
	Head() string
	// Kind
	Kind() DataKind
	AtRow(i int) vecty.MarkupOrChild
}

type DataKind int

const (
	// Keep in sync with DataTable's heads() method.
	defaultDataKind DataKind = iota
	DataString
	DataNumeric
	DataCheckbox
)

// CellClassName returns the cell class.
func (dk DataKind) CellClassName() (class string) {
	switch dk {
	case DataString, defaultDataKind:
		// no class.
	case DataNumeric:
		class = "mdc-data-table__cell--numeric"
	case DataCheckbox:
		class = "mdc-data-table__cell-checkbox"
	default:
		panic("unknown DataKind")
	}
	return class
}

// Compile-time check of interface implementation.
var (
	_ Series = (*StringSeries)(nil)
	_ Series = (*IntSeries)(nil)
	_ Series = (*FloatSeries)(nil)
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

type IntSeries struct {
	Label string
	Data  []int
}

func (ss *IntSeries) Head() string   { return ss.Label }
func (ss *IntSeries) Kind() DataKind { return DataNumeric }

func (ss *IntSeries) AtRow(i int) vecty.MarkupOrChild {
	return vecty.Text(strconv.Itoa(ss.Data[i]))
}

type FloatSeries struct {
	Label string
	Data  []float64
	Prec  int
	// Floating point format verbs (see https://pkg.go.dev/fmt)
	//  'e': scientific notation, e.g. -1.234456e+78
	//  'E': scientific notation, e.g. -1.234456E+78
	//  'f': decimal point but no exponent, e.g. 123.456
	//  'g': %e for large exponents, %f otherwise. Precision is discussed below.
	//  'G': %E for large exponents, %f otherwise
	//  'x': hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
	//  'X': upper-case hexadecimal notation, e.g. -0X1.23ABCP+20
	//  'b': decimalless scientific notation with exponent a power of two, in the manner of strconv.FormatFloat with the 'b' format e.g. -123456p-78
	Fmt byte
}

func (ss *FloatSeries) Head() string   { return ss.Label }
func (ss *FloatSeries) Kind() DataKind { return DataNumeric }

func (ss *FloatSeries) AtRow(i int) vecty.MarkupOrChild {
	const (
		defaultFmt  = 'g'
		defaultPrec = 6
	)
	prec := ss.Prec
	if prec == 0 {
		ss.Prec = defaultPrec
	}
	if ss.Fmt == 0 {
		ss.Fmt = defaultFmt
	}
	return vecty.Text(strconv.FormatFloat(ss.Data[i], defaultFmt, prec, 64))
}

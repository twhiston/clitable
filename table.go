package clitable

import (
	"io"
	"os"
	"fmt"
)

type colSize map[int]int
type row []string

// Table represents the table to be printed
type Table struct {
	columnSizes    colSize
	maximumColumns int
	rows           []row
	Fmt            TableFormat
	Writer         io.Writer
	ErrWriter      io.Writer
}

type TableFormat struct {
	Corner string
	Row    string
	Column string
	Pad    string
	Blank  string
}

func (c colSize) setIfBigger(k, v int) {
	cv, e := c[k]
	if (e && cv < v) || !e {
		c[k] = v
	}
}

// New creates a new table
func New() *Table {
	t := new(Table)
	t.columnSizes = make(colSize)
	t.Fmt.Corner = "+"
	t.Fmt.Row = "-"
	t.Fmt.Column = "| "
	t.Fmt.Pad = " "
	t.Fmt.Blank = " "
	t.Writer = os.Stdout
	t.ErrWriter = os.Stderr
	return t
}

// AddRow adds a row to the table
func (t *Table) AddRow(cols ...string) *Table {
	colCount := len(cols)
	var r row
	for i, col := range cols {
		t.columnSizes.setIfBigger(i, len(col))
		r = append(r, col)

	}
	t.rows = append(t.rows, r)

	if colCount < t.maximumColumns {
		t.ColumnPadding()
	} else if colCount > t.maximumColumns {
		t.maximumColumns = colCount
		t.ColumnPadding()
	}
	return t
}

func (t *Table) ColumnPadding() {
	for k, v := range t.rows {
		colCount := len(v)
		for colCount < t.maximumColumns {
			v = append(v, t.Fmt.Blank)
			colCount += 1
		}
		t.rows[k] = v
	}
}

// Print prints the table to standard output
func (t *Table) Print() {
	t.Fprint(t.Writer)
}

// Errprint prints the table to error output
func (t *Table) Errprint() {
	t.Fprint(t.ErrWriter)
}

// Fprint prints the table to any io.Writer of your choice
func (t *Table) Fprint(w io.Writer) {
	t.printEmpty(w)
	for _, row := range t.rows {
		t.printRow(row, w)
		t.printEmpty(w)
	}
}

func (t *Table) printRow(row row, w io.Writer) {
	colCount := 0
	for k, col := range row {
		c := string(col)
		fmt.Fprint(w, t.Fmt.Column, pad(c+"", t.columnSizes[k]+1, t.Fmt.Pad))
		colCount += 1
	}
	fmt.Fprintf(w, t.Fmt.Column+"\n")
}

func (t *Table) printEmpty(w io.Writer) {
	for i := 0; i < len(t.columnSizes); i ++ {
		fmt.Fprint(w, t.Fmt.Corner+pad("", t.columnSizes[i]+len(t.Fmt.Column), t.Fmt.Row))
	}
	fmt.Fprintf(w, t.Fmt.Corner+"\n")
}

func pad(str string, dlen int, padchar string) string {
	if len(str) < dlen {
		diff := dlen - len(str)
		app := str
		for i := 0; i < diff; i++ {
			app += string(padchar)
		}
		return app
	}
	return str
}

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
}

//TableFormat describes the characters that will be used when rendering the table
type TableFormat struct {
	Corner string
	Row    string
	Column string
	Pad    string
	Blank  string
}

func (c colSize) setIfBigger(k, v int) {
	current, exists := c[k]
	if (exists && current < v) || !exists {
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
		t.columnPadding()
	} else if colCount > t.maximumColumns {
		t.maximumColumns = colCount
		t.columnPadding()
	}
	return t
}

func (t *Table) columnPadding() {
	for k, v := range t.rows {
		colCount := len(v)
		for colCount < t.maximumColumns {
			v = append(v, t.Fmt.Blank)
			colCount++
		}
		t.rows[k] = v
	}
}

// Print prints the table to standard output
func (t *Table) Print() {
	t.Fprint(os.Stdout)
}

// Errprint prints the table to error output
func (t *Table) Errprint() {
	t.Fprint(os.Stderr)
}

// Fprint prints the table to any io.Writer of your choice
func (t *Table) Fprint(w io.Writer) {
	t.printDivider(w)
	for _, row := range t.rows {
		t.printRow(row, w)
		t.printDivider(w)
	}
}

func (t *Table) printRow(row row, w io.Writer) {
	colCount := 0
	for k, col := range row {
		c := col
		fmt.Fprint(w, t.Fmt.Column, pad(c, t.columnSizes[k]+len(t.Fmt.Corner), t.Fmt.Pad))
		colCount++
	}
	fmt.Fprintf(w, t.Fmt.Column+"\n")
}

func (t *Table) printDivider(w io.Writer) {
	for i := 0; i < len(t.columnSizes); i ++ {
		fmt.Fprint(w, t.Fmt.Corner+pad("", t.columnSizes[i]+len(t.Fmt.Column), t.Fmt.Row))
	}
	fmt.Fprintf(w, t.Fmt.Corner+"\n")
}

func pad(str string, dlen int, padchar string) string {
	if padchar == "" {
		panic("pad character cannot be null")
	}
	inputLength := len(str)
	if inputLength < dlen {
		app := str
		i := inputLength
		for ;i < dlen; {
			app += padchar
			i += len(padchar)
		}
		// Ensure correct length if pad is larger that 1 character
		app = app[:dlen]
		return app
	}
	return str
}

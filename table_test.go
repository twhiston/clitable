package clitable

import (
	"testing"
	"bytes"
)

func TestTable(t *testing.T) {
	table := New()
	if len(table.columnSizes) != 0 {
		t.Error("initialization error")
	}

	table.AddRow("One", "Two", "Three")
	if table.columnSizes[0] != 3 || table.columnSizes[1] != 3 || table.columnSizes[2] != 5 {
		t.Error("Counting error")
	}

	table.AddRow("One", "Two", "MoreThanThree")
	if table.columnSizes[0] != 3 || table.columnSizes[1] != 3 || table.columnSizes[2] != 13 {
		t.Error("re-Counting error")
	}

	table.AddRow("a", "a", "a")
	if table.columnSizes[0] != 3 || table.columnSizes[1] != 3 || table.columnSizes[2] != 13 {
		t.Error("re-Counting error")
	}
}

func TestTableOther(t *testing.T) {

	table := New()
	table.AddRow("One", "Two", "Three")
	table.AddRow("xxxxxxx", "rr")
	table.AddRow("a", "b", "c", "d", "e")

	expectedOutput := `+---------+-----+-------------------+
| One     | Two | Three             |
+---------+-----+-------------------+
| xxxxxxx | rr  | Thrffffffffffffee |
+---------+-----+-------------------+
`

	buffer := new(bytes.Buffer)
	table.Writer = buffer
	table.Print()
	if buffer.String() != expectedOutput {
		t.Fatal("incorrect output.\nExpected:\n", expectedOutput, "\nActual:\n", buffer.String())
	}
}

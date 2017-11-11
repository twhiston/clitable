package clitable

import (
	"testing"
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

	if table.maximumColumns != 3 {
		t.Error("maximum columns error")
	}

	table.AddRow("1", "2", "3", "4", "5", "6")

	if table.maximumColumns != 6 {
		t.Error("maximum columns error")
	}
}


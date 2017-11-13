package clitable

import (
	"testing"
	"github.com/Pallinder/go-randomdata"
)

//nolint
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

func TestInvalidPadChar(t *testing.T){
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("An invalid pad character, which could cause an infinite loop did not panic")
		}
	}()

	inString := "fail"
	expectedLength := 10
	pad(inString, expectedLength, "")
}


func TestPadWithLongPad(t *testing.T) {

	//Test that it pads as expected
	inString := "test"
	expectedLength := 10
	outString := pad(inString, expectedLength, "--->")
	if len(outString) != expectedLength {
		t.Fatal("pad did not output the expected length. Expected: ", expectedLength, " Actual: ", len(outString))
	}

}

func TestPad(t *testing.T) {

	for i := 0; i < 1000; i++ {
		input := randomdata.SillyName()
		strleng := randomdata.Number(len(input), len(input)+randomdata.Number(1,100))
		padchar := randomdata.StringSample(" ","->","._-!-_.", randomdata.SillyName(),randomdata.City())

		outString := pad(input, strleng, padchar)
		if len(outString) != strleng {
			t.Error("pad did not output the expected length. Expected: ", strleng, " Actual: ", len(outString))
		}
	}

}
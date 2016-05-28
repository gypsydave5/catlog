package main

import (
	"bytes"
	"testing"
)

func TestLibrary(t *testing.T) {
	var testLibrary = makeTestLibrary()

	if testLibrary[0].PublicationDate != 1847 {
		t.Error("Expected 1847, but got", testLibrary[0].PublicationDate)
	}
}

func TestWriteCSV(t *testing.T) {
	var testLibrary = makeTestLibrary()
	var out = new(bytes.Buffer)
	testLibrary.WriteCSV(out)

	result := out.String()
	expected := "Wuthering Heights,Emily Bronte,1847,Thomas Cautley Newbury,1,Kate Bush\nTess of the d'Urbervilles,Thomas Hardy,1892,James R. Osgood,1,\"Wessex,19th Century\"\n"

	if result != expected {
		t.Error("Expected", expected, "but got", result)
	}
}

func makeTestLibrary() library {

	var testBooks = []book{
		book{
			"Wuthering Heights",
			"Emily Bronte",
			1847,
			"Thomas Cautley Newbury",
			1,
			[]string{"Kate Bush"},
		},
		book{
			"Tess of the d'Urbervilles",
			"Thomas Hardy",
			1892,
			"James R. Osgood",
			1,
			[]string{"Wessex", "19th Century"},
		},
	}

	var testLibrary = library(testBooks)

	return testLibrary
}

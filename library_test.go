package main

import (
	"bytes"
	"strings"
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
	expected := "1,Wuthering Heights,Emily Bronte,1847,Thomas Cautley Newbury,1,Kate Bush\n2,Tess of the d'Urbervilles,Thomas Hardy,1892,James R. Osgood,1,\"Wessex,19th Century\"\n"

	if result != expected {
		t.Error("Expected", expected, "but got", result)
	}
}

func TestWriteJSON(t *testing.T) {
	var testLibrary = makeTestLibrary()
	var out = new(bytes.Buffer)
	testLibrary.WriteJSON(out)

	result := out.String()
	expected := `[{"ID":1,"Title":"Wuthering Heights","Author":"Emily Bronte","PublicationDate":1847,"Publisher":"Thomas Cautley Newbury","Edition":1,"Keywords":["Kate Bush"]},{"ID":2,"Title":"Tess of the d'Urbervilles","Author":"Thomas Hardy","PublicationDate":1892,"Publisher":"James R. Osgood","Edition":1,"Keywords":["Wessex","19th Century"]}]`

	if result != expected {
		t.Error("Expected", expected, "but got", result)
	}
}

func TestReadCSV(t *testing.T) {
	var lib library
	csv := "Wuthering Heights,Emily Bronte,1847,Thomas Cautley Newbury,1,Kate Bush\nTess of the d'Urbervilles,Thomas Hardy,1892,James R. Osgood,1,\"Wessex,19th Century\"\n"
	r := strings.NewReader(csv)

	lib = newLibraryFromCSV(r)
	if lib[0].Author != "Emily Bronte" {
		t.Errorf("Expected Emily Bronte, but got %v", lib[0].Author)
	}
}

func makeTestLibrary() library {

	var testBooks = []book{
		book{
			1,
			"Wuthering Heights",
			"Emily Bronte",
			1847,
			"Thomas Cautley Newbury",
			1,
			[]string{"Kate Bush"},
		},
		book{
			2,
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

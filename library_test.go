package main

import "testing"

func TestLibrary(t *testing.T) {
	b := []book{
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
			[]string{"Women making poor life choices"},
		},
	}
	lib := library(b)

	if lib[0].PublicationDate != 1847 {
		t.Error("Expected 1847, but got", lib[0].PublicationDate)
	}
}

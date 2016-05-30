package main

import "os"

type catalogue struct {
	library  library
	filename string
}

func (cat *catalogue) FetchBookByID(id int) book {
	for _, b := range cat.library {
		if b.ID == id {
			return b
		}
	}
	return book{}
}

func newJSONCatalogue(filename string) catalogue {
	catFile, _ := os.Open(filename)
	lib, _ := newLibraryFromJSON(catFile)
	return catalogue{
		lib,
		filename,
	}
}

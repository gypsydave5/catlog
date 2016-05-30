package main

import "os"

type catalogue struct {
	library  library
	filename string
	nextID   int
}

func (cat *catalogue) FetchBookByID(id int) book {
	for _, b := range cat.library {
		if b.ID == id {
			return b
		}
	}
	return book{}
}

func (cat *catalogue) FetchBookByTitle(title string) book {
	for _, b := range cat.library {
		if b.Title == title {
			return b
		}
	}
	return book{}
}

func (cat *catalogue) CreateBook(b book) {
	b.ID = cat.nextID
	cat.nextID++
	cat.library = append(cat.library, b)
}

func newJSONCatalogue(filename string) catalogue {
	catFile, _ := os.Open(filename)
	lib, _ := newLibraryFromJSON(catFile)
	nextID := lib[len(lib)-1].ID + 1
	return catalogue{
		lib,
		filename,
		nextID,
	}
}

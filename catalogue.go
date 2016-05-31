package main

import (
	"errors"
	"os"
)

var (
	errBookNotFound = errors.New("Book not found in catalogue")
)

type catalogue struct {
	library  library
	filename string
	nextID   int
}

func (cat *catalogue) FetchBookByID(id int) (book, error) {
	for _, b := range cat.library {
		if b.ID == id {
			return b, nil
		}
	}
	return book{}, errBookNotFound
}

func (cat *catalogue) FetchBookByTitle(title string) book {
	for _, b := range cat.library {
		if b.Title == title {
			return b
		}
	}
	return book{}
}

func (cat *catalogue) UpdateBook(ub book) {
	for i, b := range cat.library {
		if b.ID == ub.ID {
			cat.library[i] = ub
		}
	}
}

func (cat *catalogue) CreateBook(b book) {
	b.ID = cat.nextID
	cat.nextID++
	cat.library = append(cat.library, b)
}

func (cat *catalogue) DeleteBookWithID(id int) {
	for i, b := range cat.library {
		if b.ID == id {
			cat.library[i] = book{}
		}
	}
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

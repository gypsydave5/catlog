package main

import (
	"errors"
	"io"
)

var (
	errBookNotFound = errors.New("Book not found in catalogue")
)

type catalogue struct {
	library         library
	catalogueReader io.Reader
	catalogueWriter io.Writer
	nextID          int
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

func newJSONCatalogue(catReader io.Reader, catWriter io.Writer) catalogue {
	lib, _ := newLibraryFromJSON(catReader)
	nextID := lib[len(lib)-1].ID + 1
	return catalogue{
		lib,
		catReader,
		catWriter,
		nextID,
	}
}

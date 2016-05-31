package main

import (
	"errors"
	"io"
)

var (
	errBookNotFound = errors.New("Book not found in catalogue")
)

type catalogue interface {
	FetchBookByID(int) (book, error)
	FetchBookByTitle(string) book
	UpdateBook(book)
	CreateBook(book)
	DeleteBookWithID(int)
}

type fileCatalogue struct {
	library         library
	catalogueReader io.Reader
	catalogueWriter io.Writer
	nextID          int
}

func (cat *fileCatalogue) FetchBookByID(id int) (book, error) {
	for _, b := range cat.library {
		if b.ID == id {
			return b, nil
		}
	}
	return book{}, errBookNotFound
}

func (cat *fileCatalogue) FetchBookByTitle(title string) (book, error) {
	for _, b := range cat.library {
		if b.Title == title {
			return b, nil
		}
	}
	return book{}, errBookNotFound
}

func (cat *fileCatalogue) UpdateBook(ub book) {
	for i, b := range cat.library {
		if b.ID == ub.ID {
			cat.library[i] = ub
		}
	}
}

func (cat *fileCatalogue) CreateBook(b book) {
	b.ID = cat.nextID
	cat.nextID++
	cat.library = append(cat.library, b)
	cat.library.WriteJSON(cat.catalogueWriter)
}

func (cat *fileCatalogue) DeleteBookWithID(id int) {
	for i, b := range cat.library {
		if b.ID == id {
			cat.library[i] = book{}
		}
	}
}

func newJSONCatalogue(catReader io.Reader, catWriter io.Writer) fileCatalogue {
	lib, _ := newLibraryFromJSON(catReader)
	nextID := lib[len(lib)-1].ID + 1
	return fileCatalogue{
		lib,
		catReader,
		catWriter,
		nextID,
	}
}

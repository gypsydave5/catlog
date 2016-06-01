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
	writeCallback   func()
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

func (cat *fileCatalogue) UpdateBook(ub book) error {
	for i, b := range cat.library {
		if b.ID == ub.ID {
			cat.library[i] = ub
			cat.library.WriteJSON(cat.catalogueWriter)
			return nil
		}
	}
	return errBookNotFound
}

func (cat *fileCatalogue) CreateBook(b book) {
	b.ID = cat.nextID
	cat.nextID++
	cat.library = append(cat.library, b)
  cat.writeCallback()
	cat.library.WriteJSON(cat.catalogueWriter)
}

func (cat *fileCatalogue) DeleteBookWithID(id int) {
	for i, b := range cat.library {
		if b.ID == id {
			cat.library = append(cat.library[:i], cat.library[i+1:]...)
			cat.writeCallback()
			cat.library.WriteJSON(cat.catalogueWriter)
		}
	}
}

func newJSONCatalogue(catReader io.Reader, catWriter io.Writer, cb func()) (fileCatalogue, error) {
	lib, err := newLibraryFromJSON(catReader)
	if err != nil {
		return fileCatalogue{}, err
	}
	nextID := lib[len(lib)-1].ID + 1
	return fileCatalogue{
		library:         lib,
		catalogueReader: catReader,
		catalogueWriter: catWriter,
		nextID:          nextID,
		writeCallback:   cb,
	}, nil
}

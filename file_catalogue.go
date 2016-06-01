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

type file interface {
	io.ReadWriteSeeker
	Truncate(int64) error
}

type fileCatalogue struct {
	library library
	file    file
	nextID  int
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
			wipeFile(cat.file)
			cat.library.WriteJSON(cat.file)
			return nil
		}
	}
	return errBookNotFound
}

func (cat *fileCatalogue) CreateBook(b book) {
	b.ID = cat.nextID
	cat.nextID++
	cat.library = append(cat.library, b)
	wipeFile(cat.file)
	cat.library.WriteJSON(cat.file)
}

func (cat *fileCatalogue) DeleteBookWithID(id int) {
	for i, b := range cat.library {
		if b.ID == id {
			cat.library = append(cat.library[:i], cat.library[i+1:]...)
			wipeFile(cat.file)
			cat.library.WriteJSON(cat.file)
		}
	}
}

func newJSONCatalogue(f file) (fileCatalogue, error) {
	lib, err := newLibraryFromJSON(f)
	if err != nil {
		return fileCatalogue{}, err
	}
	nextID := lib[len(lib)-1].ID + 1
	return fileCatalogue{
		library: lib,
		file:    f,
		nextID:  nextID,
	}, nil
}

func wipeFile(f file) error {
	f.Seek(0, 0)
	err := f.Truncate(0)
	if err != nil {
		return err
	}

	return err
}

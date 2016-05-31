package main

import (
	"bytes"
	"testing"
)

var (
	dune = book{
		Title:           "Dune",
		Author:          "Frank Herbert",
		PublicationDate: 1965,
		Publisher:       "Chilton Books",
		Edition:         1,
		Keywords:        []string{"Desert", "Science Fiction"},
	}
)

func TestFileCatalogue(t *testing.T) {
	cat := newTestCatalogue()
	book, _ := cat.FetchBookByID(1)
	if book.Title != "Wuthering Heights" {
		t.Error("Expected Wuthering Heights, got", book.Title)
	}
}

func TestFetchBookByTitleError(t *testing.T) {
	cat := newTestCatalogue()
	_, err := cat.FetchBookByID(6)
	if err != errBookNotFound {
		t.Error("Expected errBookNotFound to be returned")
	}
}

func TestFetchBookByID(t *testing.T) {
	cat := newTestCatalogue()
	book, _ := cat.FetchBookByTitle("Tess of the d'Urbervilles")
	if book.ID != 2 {
		t.Error("Expected 1, got", book.ID)
	}
}

func TestFetchBookByIDError(t *testing.T) {
	cat := newTestCatalogue()
	_, err := cat.FetchBookByTitle("Cardenio")
	if err != errBookNotFound {
		t.Error("Expected errBookNotFound to be returned")
	}
}

func TestAddToCatalogue(t *testing.T) {
	testCatalogueBuffer := newTestCatalogueBuffer()
	cat := newJSONCatalogue(testCatalogueBuffer, testCatalogueBuffer)

	cat.CreateBook(dune)

	book, err := cat.FetchBookByID(3)

	if err != nil {
		t.Error("Error reading 'Dune' from catalogue")
	}

	if book.Title != "Dune" {
		t.Error("Expected the title of 'Dune' to be 'Dune', but instead it was", dune.Title)
	}

	expectedBuffer := `[{"ID":1,"Title":"Wuthering Heights","Author":"Emily Bronte","PublicationDate":1847,"Publisher":"Thomas Cautley Newbury","Edition":1,"Keywords":["Kate Bush"]},{"ID":2,"Title":"Tess of the d'Urbervilles","Author":"Thomas Hardy","PublicationDate":1892,"Publisher":"James R. Osgood","Edition":1,"Keywords":["Wessex","19th Century"]},{"ID":3,"Title":"Dune","Author":"Frank Herbert","PublicationDate":1965,"Publisher":"Chilton Books","Edition":1,"Keywords":["Desert","Science Fiction"]}]`

	if testCatalogueBuffer.String() != expectedBuffer {
		t.Error("'Dune' was not written to the catalogue buffer")
	}
}

func TestUpdateBookInCatalogue(t *testing.T) {
	testCatalogueBuffer := newTestCatalogueBuffer()
	cat := newJSONCatalogue(testCatalogueBuffer, testCatalogueBuffer)

	book, _ := cat.FetchBookByTitle("Wuthering Heights")
	book.Title = "Heathcliff!"
	book.Author = "Cliff Richard"
	cat.UpdateBook(book)
	updatedBook, _ := cat.FetchBookByID(1)

	if updatedBook.Title != "Heathcliff!" {
		t.Error("Expected Heathcliff!, but got", updatedBook.Title)
	}
}

func TestDeleteBookInCatalogue(t *testing.T) {
	testCatalogueBuffer := newTestCatalogueBuffer()
	cat := newJSONCatalogue(testCatalogueBuffer, testCatalogueBuffer)
	cat.DeleteBookWithID(1)

	_, err := cat.FetchBookByID(1)

	if err != errBookNotFound {
		t.Error("Expected book not found error, but got:", err)
	}
}

func newTestCatalogueBuffer() *bytes.Buffer {
	catalogueString := `[{"ID":1,"Title":"Wuthering Heights","Author":"Emily Bronte","PublicationDate":1847,"Publisher":"Thomas Cautley Newbury","Edition":1,"Keywords":["Kate Bush"]},{"ID":2,"Title":"Tess of the d'Urbervilles","Author":"Thomas Hardy","PublicationDate":1892,"Publisher":"James R. Osgood","Edition":1,"Keywords":["Wessex","19th Century"]}]`
	return bytes.NewBufferString(catalogueString)
}

func newTestCatalogue() fileCatalogue {
	catBuffer := newTestCatalogueBuffer()
	return newJSONCatalogue(catBuffer, catBuffer)
}

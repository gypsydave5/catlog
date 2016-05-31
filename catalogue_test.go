package main

import (
	"bytes"
	"os"
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
	catalogueString = `[{"ID":1,"Title":"Wuthering Heights","Author":"Emily Bronte","PublicationDate":1847,"Publisher":"Thomas Cautley Newbury","Edition":1,"Keywords":["Kate Bush"]},{"ID":2,"Title":"Tess of the d'Urbervilles","Author":"Thomas Hardy","PublicationDate":1892,"Publisher":"James R. Osgood","Edition":1,"Keywords":["Wessex","19th Century"]}]`
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
	cat, err := newJSONCatalogue(testCatalogueBuffer, testCatalogueBuffer, func() {})

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
	cat, _ := newJSONCatalogue(testCatalogueBuffer, testCatalogueBuffer, func() {})

	book, _ := cat.FetchBookByTitle("Wuthering Heights")
	book.Title = "Heathcliff!"
	book.Author = "Cliff Richard"
	cat.UpdateBook(book)
	updatedBook, _ := cat.FetchBookByID(1)

	if updatedBook.Title != "Heathcliff!" {
		t.Error("Expected Heathcliff!, but got", updatedBook.Title)
	}

	fileLib, _ := newLibraryFromJSON(cat.catalogueReader)
	bookTitleFromFile := fileLib[0].Title

	if bookTitleFromFile != "Heathcliff!" {
		t.Error("Expected Heathcliff!, but got", bookTitleFromFile)
	}
}

func TestUpdateBookInCatalogueError(t *testing.T) {
	cat := newTestCatalogue()
	err := cat.UpdateBook(dune)
	if err != errBookNotFound {
		t.Error("Expected an errBookNotFound")
	}
}

func TestDeleteBookInCatalogue(t *testing.T) {
	testCatalogueBuffer := newTestCatalogueBuffer()
	cat, _ := newJSONCatalogue(testCatalogueBuffer, testCatalogueBuffer, func() {})
	cat.DeleteBookWithID(1)

	_, err := cat.FetchBookByID(1)

	if err != errBookNotFound {
		t.Error("Expected book not found error, but got:", err)
	}

	fileLib, err := newLibraryFromJSON(cat.catalogueReader)
	if err != nil {
		t.Error("Unexpected error on book deletion", err)
	}
	bookTitleFromFile := fileLib[0].Title
	if bookTitleFromFile != "Tess of the d'Urbervilles" {
		t.Error("Expected Tess of the d'Urbervilles, but got", bookTitleFromFile)
	}
}

func TestCatalogueFileGetsWrittenTo(t *testing.T) {
	catalogueFile := setUpTestCatalogueFile()
	cat, _ := newJSONCatalogue(catalogueFile, catalogueFile, func() {
		catalogueFile.Truncate(0)
		catalogueFile.Seek(0, 0)
	})
	cat.DeleteBookWithID(1)

	_, err := cat.FetchBookByID(1)
	if err != errBookNotFound {
		t.Error("Expected book not found error, but got:", err)
	}

	testCatalogueFile, err := os.Open("testCatalogue.json")
	if err != nil {
		t.Error("Error opening testCatalogue.json:", err)
	}
	fileLib, err := newLibraryFromJSON(testCatalogueFile)
	if err != nil {
		t.Error("Error parsing testCatalogue.json", err)
	}

	if len(fileLib) != 1 {
		t.Errorf("Expected only one book record, but got %d", len(fileLib))
	}
	bookTitleFromFile := fileLib[0].Title
	if bookTitleFromFile != "Tess of the d'Urbervilles" {
		t.Error("Expected Tess of the d'Urbervilles, but got", bookTitleFromFile)
	}
	bookIDFromFile := fileLib[0].ID
	if bookIDFromFile != 2 {
		t.Errorf("Expected 2, but got %d", bookIDFromFile)
	}
}

func newTestCatalogueBuffer() *bytes.Buffer {
	return bytes.NewBufferString(catalogueString)
}

func newTestCatalogue() fileCatalogue {
	catBuffer := newTestCatalogueBuffer()
	cat, _ := newJSONCatalogue(catBuffer, catBuffer, func() {})
	return cat
}

func setUpTestCatalogueFile() *os.File {
	catalogueFile, _ := os.Create("testCatalogue.json")
	catalogueFile.WriteString(catalogueString)
	catalogueFile.Close()
	catalogueFile, _ = os.OpenFile("testCatalogue.json", os.O_RDWR, 0666)
	return catalogueFile
}

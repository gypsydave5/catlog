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

type mockFiletype struct {
	buffer *bytes.Buffer
}

func (mft mockFiletype) Read(b []byte) (int, error)         { return mft.buffer.Read(b) }
func (mft mockFiletype) Write(b []byte) (int, error)        { return mft.buffer.Write(b) }
func (mft mockFiletype) Truncate(i int64) error             { return nil }
func (mft mockFiletype) Seek(i int64, j int) (int64, error) { return 0, nil }

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
	mockFile := newMockFileType(testCatalogueBuffer)
	cat, _ := newJSONCatalogue(mockFile)

	cat.CreateBook(dune)

	expectedBuffer := `[{"ID":1,"Title":"Wuthering Heights","Author":"Emily Bronte","PublicationDate":1847,"Publisher":"Thomas Cautley Newbury","Edition":1,"Keywords":["Kate Bush"]},{"ID":2,"Title":"Tess of the d'Urbervilles","Author":"Thomas Hardy","PublicationDate":1892,"Publisher":"James R. Osgood","Edition":1,"Keywords":["Wessex","19th Century"]},{"ID":3,"Title":"Dune","Author":"Frank Herbert","PublicationDate":1965,"Publisher":"Chilton Books","Edition":1,"Keywords":["Desert","Science Fiction"]}]`

	if testCatalogueBuffer.String() != expectedBuffer {
		t.Error("'Dune' was not written to the catalogue buffer")
	}
}

func TestUpdateBookInCatalogue(t *testing.T) {
	testCatalogueBuffer := newTestCatalogueBuffer()
	mockFile := newMockFileType(testCatalogueBuffer)
	cat, _ := newJSONCatalogue(mockFile)
	book, _ := cat.FetchBookByTitle("Wuthering Heights")
	book.Title = "Heathcliff!"
	book.Author = "Cliff Richard"

	cat.UpdateBook(book)

	fileLib, _ := newLibraryFromJSON(cat.file)
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
	mockFile := newMockFileType(testCatalogueBuffer)
	cat, _ := newJSONCatalogue(mockFile)

	cat.DeleteBookWithID(1)

	fileLib, _ := newLibraryFromJSON(cat.file)
	bookTitleFromFile := fileLib[0].Title
	if bookTitleFromFile != "Tess of the d'Urbervilles" {
		t.Error("Expected Tess of the d'Urbervilles, but got", bookTitleFromFile)
	}
}

func TestCatalogueFileGetsWrittenTo(t *testing.T) {
	catalogueFile := setUpTestCatalogueFile()
	cat, _ := newJSONCatalogue(catalogueFile)
	cat.DeleteBookWithID(1)

	_, err := cat.FetchBookByID(1)
	if err != errBookNotFound {
		t.Error("Expected book not found error, but got:", err)
	}

	testCatalogueFile, _ := os.Open("testCatalogue.json")
	fileLib, _ := newLibraryFromJSON(testCatalogueFile)

	bookIDFromFile := fileLib[0].ID
	if bookIDFromFile != 2 {
		t.Errorf("Expected 2, but got %d", bookIDFromFile)
	}
}

func newTestCatalogueBuffer() *bytes.Buffer {
	return bytes.NewBufferString(catalogueString)
}

func newTestCatalogue() fileCatalogue {
	var mft file
	catBuffer := newTestCatalogueBuffer()
	mft = newMockFileType(catBuffer)
	cat, _ := newJSONCatalogue(mft)
	return cat
}

func newMockFileType(b *bytes.Buffer) file {
	return mockFiletype{buffer: b}
}

func setUpTestCatalogueFile() *os.File {
	fileName := "testCatalogue.json"
	catalogueFile, _ := os.Create(fileName)
	catalogueFile.WriteString(catalogueString)
	catalogueFile.Close()
	catalogueFile, _ = os.OpenFile("testCatalogue.json", os.O_RDWR, 0666)
	return catalogueFile
}

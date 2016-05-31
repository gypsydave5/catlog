package main

import (
	"bytes"
	"os"
	"testing"
)

func TestFileCatalogue(t *testing.T) {
	SetUpFileCatalogue(t)
	cat := newJSONCatalogue("test_catalogue.json")
	book, _ := cat.FetchBookByID(1)
	if book.Title != "Wuthering Heights" {
		t.Error("Expected Wuthering Heights, got", book.Title)
	}

	TearDownFileCatalogue(t)
}

func TestAddToCatalogue(t *testing.T) {
	SetUpFileCatalogue(t)
	cat := newJSONCatalogue("test_catalogue.json")
	bk := book{
		Title:           "Dune",
		Author:          "Frank Herbert",
		PublicationDate: 1965,
		Publisher:       "Chilton Books",
		Edition:         1,
		Keywords:        []string{"Desert", "Science Fiction"},
	}

	cat.CreateBook(bk)
	book := cat.FetchBookByTitle("Dune")

	if book.Author != "Frank Herbert" {
		t.Error("Expected Frank Herbert, got", book.Author)
	}
	if book.ID != 3 {
		t.Error("Expected 3, got", book.ID)
	}
	TearDownFileCatalogue(t)
}

func TestUpdateBookInCatalogue(t *testing.T) {
	SetUpFileCatalogue(t)
	cat := newJSONCatalogue("test_catalogue.json")
	book := cat.FetchBookByTitle("Wuthering Heights")
	book.Title = "Heathcliff!"
	book.Author = "Cliff Richard"
	cat.UpdateBook(book)
	updatedBook, _ := cat.FetchBookByID(1)

	if updatedBook.Title != "Heathcliff!" {
		t.Error("Expected Heathcliff!, but got", updatedBook.Title)
	}

	TearDownFileCatalogue(t)
}

func TestDeleteBookInCatalogue(t *testing.T) {
	SetUpFileCatalogue(t)
	cat := newJSONCatalogue("test_catalogue.json")
	cat.DeleteBookWithID(1)

	_, err := cat.FetchBookByID(1)

	if err != errBookNotFound {
		t.Error("Expected book not found error, but got:", err)
	}

	TearDownFileCatalogue(t)
}

func SetUpFileCatalogue(t *testing.T) {
	file, err := os.Create("test_catalogue.json")
	if err != nil {
		t.Error("Unexpected error on file creation:", err)
	}
	file.WriteString(`[{"ID":1,"Title":"Wuthering Heights","Author":"Emily Bronte","PublicationDate":1847,"Publisher":"Thomas Cautley Newbury","Edition":1,"Keywords":["Kate Bush"]},{"ID":2,"Title":"Tess of the d'Urbervilles","Author":"Thomas Hardy","PublicationDate":1892,"Publisher":"James R. Osgood","Edition":1,"Keywords":["Wessex","19th Century"]}]`)
	file.Close()
}

func TearDownFileCatalogue(t *testing.T) {
	err := os.Remove("test_catalogue.json")
	if err != nil {
		t.Error("Unexpected error on file deletion:", err)
	}
}

func newTestCatalogue() *bytes.Buffer {
	catalogueString := `[{"ID":1,"Title":"Wuthering Heights","Author":"Emily Bronte","PublicationDate":1847,"Publisher":"Thomas Cautley Newbury","Edition":1,"Keywords":["Kate Bush"]},{"ID":2,"Title":"Tess of the d'Urbervilles","Author":"Thomas Hardy","PublicationDate":1892,"Publisher":"James R. Osgood","Edition":1,"Keywords":["Wessex","19th Century"]}]`
	return bytes.NewBufferString(catalogueString)
}

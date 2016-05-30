package main

import (
	"os"
	"testing"
)

func TestFileCatalogue(t *testing.T) {
	SetUpFileCatalogue(t)

	cat := newJSONCatalogue("test_catalogue.json")
	book := cat.FetchBookByID(1)
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

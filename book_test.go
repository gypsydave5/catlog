package main

import "testing"

var b = book{
	Title:           "Northanger Abbey",
	Author:          "Jane Austen",
	PublicationDate: 1817,
	Publisher:       "John Murray",
	Edition:         1,
	Keywords:        []string{"19th Century", "Ghost"},
}

func TestBookProperties(t *testing.T) {
	if b.Author != "Jane Austen" {
		t.Error("Expected Jane Austen, but got", b.Author)
	}
}

func TestBookToStringSlice(t *testing.T) {
	record := b.ToStringSlice()
	if record[3] != "John Murray" {
		t.Error("Expected John Murray, but got", record[3])
	}
	if record[4] != "1" {
		t.Error("Expected 1, but got", record[4])
	}
	if record[5] != "19th Century,Ghost" {
		t.Error("Expected 19th Century,Ghost, but got", record[5])
	}
}

func TestStringSliceToBook(t *testing.T) {
	var b book
	stringSliceBook := []string{
		"The Shockwave Rider",
		"John Brunner",
		"1975",
		"Harper & Row",
		"1",
		"Science Fiction,Cyberpunk",
	}

	b = NewBookFromStringSlice(stringSliceBook)

	if b.Author != "John Brunner" {
		t.Error("Expected John Brunner, but got", b.Author)
	}

	if b.Keywords[1] != "Cyberpunk" {
		t.Error("Expected Cyberpunk, but got", b.Keywords[1])
	}

}

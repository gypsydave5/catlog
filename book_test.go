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
	record := b.ToStrings()
	if record[3] != "John Murray" {
		t.Error("Expected John Murray, but got", b.Author)
	}
	if record[5] != "19th Century,Ghost" {
		t.Error("Expected 19th Century,Ghost, but got", record[5])
	}
}

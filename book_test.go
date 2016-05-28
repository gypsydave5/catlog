package main

import "testing"

func TestBookProperties(t *testing.T) {
	b := book{
		Title:           "Northanger Abbey",
		Author:          "Jane Austen",
		PublicationDate: "1817",
		Publisher:       "John Murray",
		Edition:         "1st",
		Keywords:        []string{"19th Century", "Ghost"},
	}
}

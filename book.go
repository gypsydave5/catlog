package main

import "strings"

type book struct {
	Title           string
	Author          string
	PublicationDate int
	Publisher       string
	Edition         int
	Keywords        []string
}

func (b *book) ToStrings() []string {
	result := make([]string, 6)
	result[0] = b.Title
	result[1] = b.Author
	result[2] = string(b.PublicationDate)
	result[3] = b.Publisher
	result[4] = string(b.Edition)
	result[5] = strings.Join(b.Keywords, ",")
	return result
}

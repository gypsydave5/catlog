package main

import "strings"
import "strconv"

type book struct {
	Title           string
	Author          string
	PublicationDate int
	Publisher       string
	Edition         int
	Keywords        []string
}

func (b *book) ToStringSlice() []string {
	result := make([]string, 6)
	result[0] = b.Title
	result[1] = b.Author
	result[2] = strconv.Itoa(b.PublicationDate)
	result[3] = b.Publisher
	result[4] = strconv.Itoa(b.Edition)
	result[5] = strings.Join(b.Keywords, ",")
	return result
}

func NewBookFromStringSlice(ss []string) (b book) {
	b.Title = ss[0]
	b.Author = ss[1]
	b.PublicationDate, _ = strconv.Atoi(ss[2])
	b.Publisher = ss[3]
	b.Edition, _ = strconv.Atoi(ss[4])
	b.Keywords = strings.Split(ss[5], ",")
	return
}

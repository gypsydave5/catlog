package main

import "strings"
import "strconv"

type book struct {
	ID              int
	Title           string
	Author          string
	PublicationDate int
	Publisher       string
	Edition         int
	Keywords        []string
}

func (b *book) ToStringSlice() []string {
	result := make([]string, 7)
	result[0] = strconv.Itoa(b.ID)
	result[1] = b.Title
	result[2] = b.Author
	result[3] = strconv.Itoa(b.PublicationDate)
	result[4] = b.Publisher
	result[5] = strconv.Itoa(b.Edition)
	result[6] = strings.Join(b.Keywords, ",")
	return result
}

func newBookFromStringSlice(ss []string) (b book) {
	b.Title = ss[0]
	b.Author = ss[1]
	b.PublicationDate, _ = strconv.Atoi(ss[2])
	b.Publisher = ss[3]
	b.Edition, _ = strconv.Atoi(ss[4])
	b.Keywords = strings.Split(ss[5], ",")
	return
}

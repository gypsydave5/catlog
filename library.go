package main

import "io"

import "encoding/csv"
import "encoding/json"

type library []book

func (l library) WriteCSV(w io.Writer) {
	csvWriter := csv.NewWriter(w)

	for _, book := range l {
		csvWriter.Write(book.ToStringSlice())
	}

	csvWriter.Flush()
}

func (l library) WriteJSON(w io.Writer) {
	b, _ := json.Marshal(l)
	w.Write(b)
}

// NewLibraryFromCSV returns a new library by reading in a CSV file. Each row
// represents a book. Column order is as follows:
//   Title
//   Author
//   PublicationDate
//   Publisher
//   Edition
//   Tags
func newLibraryFromCSV(r io.Reader) library {
	var lib library
	csvReader := csv.NewReader(r)

	for {
		bookSlice, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		lib = append(lib, newBookFromStringSlice(bookSlice))
	}

	return lib
}

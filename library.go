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

func NewLibraryFromCSV(r io.Reader) library {
	csvReader := csv.NewReader(r)
	var lib library
	for {
		bookSlice, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		lib = append(lib, NewBookFromStringSlice(bookSlice))
	}
	return lib
}

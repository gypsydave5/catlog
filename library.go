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

//func NewLibraryFromCSV(r io.Reader) library {
//csvBook, _ := csv.NewReader(r).Read()
//}

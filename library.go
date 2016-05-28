package main

import (
	"fmt"
	"io"
)
import "encoding/csv"

type library []book

func (l library) WriteCSV(w io.Writer) {
	csvWriter := csv.NewWriter(w)
	for _, book := range l {
		fmt.Printf("book title: %s\n", book.Title)
		csvWriter.Write(book.ToStringSlice())
	}
	csvWriter.Flush()
}

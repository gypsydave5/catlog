package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
)

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

func newLibraryFromJSON(r io.Reader) (library, error) {
	var lib library
	b, _ := ioutil.ReadAll(r)
	err := json.Unmarshal(b, &lib)
	if err != nil {
		err := fmt.Errorf("Error reading catalogue: %v", err)
		return library{}, err
	}
	return lib, nil
}

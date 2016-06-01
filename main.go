package main

import (
	"os"
	"log"
)

func main() {
	myFile, err := os.OpenFile("my_cats.json", os.O_RDWR, 0660)
  if err != nil {
		log.Fatal(err)
	}

  myFirstCatalogue, err := newJSONCatalogue(myFile, myFile, func () {
		myFile.Truncate(0)
		myFile.Seek(0, 0)
	})

	if err != nil {
		log.Fatal(err)
	}

	myNewBook := book{
		Title: "MY BOOK!!!!!!!!!!!!",
		Author: "Sam",
		PublicationDate: 2016,
		Publisher: "The Go Team",
		Edition: 1,
		Keywords: []string {"Do", "Not", "Read"},
	}

	myFirstCatalogue.CreateBook(myNewBook)

	log.Println(myFirstCatalogue.library)
}

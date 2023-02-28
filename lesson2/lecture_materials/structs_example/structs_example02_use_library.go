package main

import (
	"fmt"
	"structs_example/books"
)

func DemonstrateStructConstructors() {
	return
	book, err := books.New("me", "title", nil)
	if err != nil {
		fmt.Println("could not create book", err)
		return
	}
	fmt.Println(book)
}

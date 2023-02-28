package books

import (
	"fmt"
)

type Book struct {
	Title      string
	Author     string
	pagesCount int
}

func DemonstrateInitializers() {
	return
	b1 := Book{Title: "Don Quixote", Author: "Miguel de Cervantes", pagesCount: 120}
	fmt.Println(b1)

	b2 := Book{Title: "Celtic Mythos"}
	fmt.Println(b2)

	b3 := Book{}
	fmt.Println(b3)

	b4 := Book{"The Hobbit", "J. R. R. Tolkien", 300}
	fmt.Println(b4)

	fmt.Println("b1 == b2: ", b1 == b2)
}

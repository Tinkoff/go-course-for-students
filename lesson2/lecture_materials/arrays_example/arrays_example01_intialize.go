package main

import "fmt"

func DemonstrateArrayInitializers() {
	return
	var a [3]int
	fmt.Println("var a [3]int:", a)

	fmt.Println("[3]int{1, 2, 3}:", [3]int{1, 2, 3})

	fmt.Println("[3]int{1}:", [3]int{1})

	fmt.Println("[...]int{1, 2, 3, 4, 5}:", [...]int{1, 2, 3, 4, 5})
}

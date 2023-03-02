package main

import "fmt"

func DemonstrateLinebreaks() {
	return
	a := "hel \n lo"
	b := `hel \n lo`

	fmt.Println(a)
	fmt.Println(b)
}

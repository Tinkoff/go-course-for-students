package main

import "fmt"

type counter struct {
	count int
}

func DemonstrateMapContentModification() {
	return
	a := map[int]counter{}

	a[0] = counter{}
	fmt.Println("a:", a, "a[0]:", a[0])

	//a[0].count++
	//fmt.Println("a:", a, "a[0]:", a[0])
}

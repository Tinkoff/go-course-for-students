package main

import "fmt"

func DemonstrateMapCreate() {
	return
	var a map[string]int

	b := map[string]int{
		"test":   0,
		"test-2": 2,
	}

	c := make(map[string]int, 10)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	//var m map[ []int ]int
}

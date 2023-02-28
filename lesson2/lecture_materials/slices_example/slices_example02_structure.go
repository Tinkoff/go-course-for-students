package main

import "fmt"

func DemonstrateSlicesInternalStructure() {
	return
	a := []int{100, 200, 300}
	b := a

	fmt.Println("a:", a, "b:", b)

	//fmt.Println("b[0] = -1")
	//b[0] = -1
	//fmt.Println("a:", a, "b:", b)

	//fmt.Println("\nb = append(b, 400);b[0] = -2")
	//b = append(b, 400)
	//b[0] = -2
	//fmt.Println("a:", a, "b:", b)

	//fmt.Println("\na = b;b = append(b, 400);b[0] = -3")
	//a = b
	//b = append(b, 500)
	//b[0] = -3
	//fmt.Println("a:", a, "b:", b)
}

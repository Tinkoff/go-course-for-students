package main

import "fmt"

func DemonstrateSubslices() {
	return
	a := []int{0, 1, 2, 3, 4}

	fmt.Println("a[:]", a[:])
	fmt.Println("a[2:3]", a[2:3])
	fmt.Println("a[2:]", a[2:])
	fmt.Println("a[:3]", a[:3])

	//b := a[2:4]
	//b[0] = -1
	//fmt.Println("a:", a, "b:", b)
}

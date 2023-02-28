package main

import "fmt"

func DemonstrateOperations() {
	return
	a := "test"

	fmt.Println("a:", a)
	fmt.Println("a[0]:", a[0])
	fmt.Println("[]byte(a):", []byte(a))

	//a[0] = 'r'
	//fmt.Printf("pointers: %p %p\n", &[]byte(a)[0], &[]byte(a)[0])
}

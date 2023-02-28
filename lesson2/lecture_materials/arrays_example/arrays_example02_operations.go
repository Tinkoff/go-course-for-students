package main

import (
	"fmt"
)

func ConsumeArray(arr [3]int) {
	fmt.Println(arr)
}

func DemonstrateArrayOperations() {
	return
	a := [...]int{100, 200, 300}
	b := [...]int{300, 400, 500}

	fmt.Println("a:", a, "b:", b)
	fmt.Println("len(a):", len(a))
	fmt.Println("a == b:", a == b)

	//b = a
	//fmt.Println("b[0] before:", b[0])
	//b[0] = -1
	//fmt.Println("b[0] after:", b[0], "a[0] after:", a[0])

	//ConsumeArray(a)

	//for i, v := range a {
	//	fmt.Println("i:", i, "v:", v)
	//}

	//{
	//	var i int
	//	var v int
	//	for ; i < len(a); i++ {
	//		v = a[i]
	//		fmt.Println("i:", i, "v:", v)
	//	}
	//}

	//for x := range a {
	//	fmt.Println("x:", x)
	//}
}

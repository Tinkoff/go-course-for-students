package main

import (
	"encoding/hex"
	"fmt"
)

func DemonstrateRussianText() {
	return
	a := "ัะตัั!"

	fmt.Println(a, len(a))

	//fmt.Println(a[1:5])

	//printHex(a)
}

func printHex(s string) {
	fmt.Println(hex.Dump([]byte(s)))
}

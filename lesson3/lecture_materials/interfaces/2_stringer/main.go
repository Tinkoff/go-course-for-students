package main

import (
	"fmt"
	"strconv"
)

// common interfaces: https://gist.github.com/asukakenji/ac8a05644a2e98f1d5ea8c299541fce9

type superInt int

func (i superInt) String() string {
	return strconv.Itoa(int(i) + 1)
}

func main() {
	var i superInt = 4
	fmt.Printf("raw superInt : %d\n", int(i))
	fmt.Printf("superInt: %s\n", i)
	fmt.Println("superInt:", i)

	// s := fmt.Stringer(i)
	// fmt.Printf("stringer: %s\n", s)
}

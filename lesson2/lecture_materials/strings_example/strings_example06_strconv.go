package main

import (
	"fmt"
	"strconv"
)

func DemonstrateStrconv() {
	return
	i := 42

	s1 := strconv.FormatInt(int64(i), 10)
	s2 := strconv.Itoa(i)
	//s3 := fmt.Sprintf("%d", i)

	fmt.Println(s1, s2)

	//parsed, err := strconv.ParseInt("42", 10, strconv.IntSize)
	//fmt.Println(parsed, err)
}

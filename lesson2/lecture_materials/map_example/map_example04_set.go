package main

import "fmt"

func UseSets() {
	return
	v := map[int]struct{}{}

	v[0] = struct{}{}
	_, contains := v[0]
	fmt.Println(contains)
}

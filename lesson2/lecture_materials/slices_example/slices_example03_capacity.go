package main

import (
	"fmt"
	"strconv"
)

func IntsToStrings(ints []int) []string {
	result := make([]string, len(ints))
	for i, v := range ints {
		result[i] = strconv.Itoa(v)
	}
	return result
}

func DemonstrateSlicesCapacity() {
	return

	stringsSlice := IntsToStrings([]int{3, 4, 5})
	fmt.Println(stringsSlice, "len:", len(stringsSlice), "cap:", cap(stringsSlice))
}

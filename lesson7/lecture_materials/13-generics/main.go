package main

import "fmt"

func main() {

	exist := ContainString([]string{"one", "two", "three"}, "two")
	fmt.Printf("is two exist: %v\n", exist)

	exist = ContainsInt([]int{1, 2, 3}, 3)
	fmt.Printf("is 3 exist: %v\n", exist)

	//fmt.Printf("a != b: %v\n", IsEqual("a", "b"))
	//fmt.Printf("a != 1: %v\n", IsEqual("a", 1)) // cant be compiled
	//fmt.Printf("1 == 1: %v\n", IsEqual(1, 1))

}

func ContainString(s []string, needle string) bool {
	for _, v := range s {
		if v == needle {
			return true
		}
	}
	return false
}

func ContainsInt(s []int, needle int) bool {
	for _, v := range s {
		if v == needle {
			return true
		}
	}
	return false
}

func Contains[T comparable](t []T, needle T) bool {
	for _, v := range t {
		if v == needle {
			return true
		}
	}
	return false
}

func IsEqual[T comparable](a, b T) bool {
	return a == b
}

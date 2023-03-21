package main

import (
	"fmt"
)

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func decorator(fn func(string)) func(string) {
	fmt.Println("before func")

	return func(arg string) {
		arg = "Sr. " + arg

		fn(arg)

		fmt.Println("after func")
	}
}

func main() {
	f := sayHello
	f = decorator(f)
	f("John")
}

package main

import "fmt"

func main() {
	var b interface{}

	b = "Hello, Go!"
	_ = any("Hello, Go!")
	fmt.Printf("%T, %+v\n", b, b)

	b = []int{1, 2, 3}
	_ = any([]int{1, 2, 3})
	fmt.Printf("%T, %+v\n", b, b)

	b = func() { fmt.Println("Hello, Go!") }
	_ = any(func() { fmt.Println("Hello, Go!") })
	fmt.Printf("%T, %+v\n", b, b)

	// slice of anything!
	b = []any{1, "a", map[string]string{}, struct{}{}, func() {}}
	fmt.Printf("%T, %+v\n", b, b)

	// map of anything!
	b = map[string]any{"int": 1, "string": "Hello, Go!", "func": func() {}}
	fmt.Printf("%T, %+v\n", b, b)

	// struct with anything!
	b = struct {
		A any
		B any
	}{
		A: [10]int{},
		B: 15.456,
	}
	fmt.Printf("%T, %+v\n", b, b)
}

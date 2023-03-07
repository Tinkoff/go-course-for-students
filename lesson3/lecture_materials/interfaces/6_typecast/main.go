package main

import (
	"fmt"
	"strconv"
)

type superInt int

func (i *superInt) String() string {
	return strconv.Itoa(int(*i) + 1)
}

type myStringer interface {
	String() string
}

func main() {
	var b any = "Hi, Go!"

	// assert to simple type
	if s, ok := b.(string); ok {
		fmt.Printf("successful typecast: %s\n", s)
	} else {
		fmt.Printf("unsuccessful typecast: '%s', %T\n", s, b)
	}

	/*
		si := superInt(5)
		b = &si
		// b = myStringer(&si)
		// b = superInt(5)
		if i, ok := b.(int); ok {
			fmt.Printf("successful typecast to int: %d\n", i)
		} else if s, ok := b.(fmt.Stringer); ok {
			fmt.Printf("successful typecast to Stringer: %v\n", s)
		} else if s, ok := b.(myStringer); ok {
			fmt.Printf("successful typecast to myStringer: %v\n", s)
		} else if mi, ok := b.(superInt); ok {
			fmt.Printf("successful typecast to myInt: %s\n", mi)
		} else {
			fmt.Printf("unsuccessful typecast for %T\n", b)
		}
	*/

	/*
		switch a := b.(type) {
		case string: // basic types
			fmt.Println("b is string", a, len(a))
		case superInt: // structs
			fmt.Println("b is superInt", a)
		case fmt.Stringer: // interfaces
			fmt.Println("b is Stringer", a, a.String())
		default:
			fmt.Printf("b is %T %v\n", b, a)
		}
	*/
}

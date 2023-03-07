package main

import (
	"fmt"
	"math"
)

type Spacer interface {
	Space() float64
}

// ------------------------------------------
type circle struct {
	radius float64
}

func (s circle) Space() float64 {
	return math.Pi * math.Pow(s.radius, 2)
}

type rectangle struct {
	width  float64
	height float64
}

func (s rectangle) Space() float64 {
	return s.width * s.height
}

// ------------------------------------------

func main() {
	var s1 Spacer = circle{radius: 10}
	var s2 Spacer = &rectangle{width: 12, height: 10}
	fmt.Println("s1 == s2", s1 == s2)
	// var f fmt.Stringer
	// fmt.Println("s1 == f", s1 == f)

	fmt.Println("s1 == raw", s1 == circle{radius: 10})

	var s3 Spacer = circle{radius: 10}
	fmt.Println("s1 == s3", s1 == s3)

	/*
		    // c := &circle{radius: 10}
			var s4 Spacer = &circle{radius: 10}
			var s5 Spacer = &circle{radius: 10}
			fmt.Println("s1 == s4", s1 == s4)
			fmt.Println("s4 == s5", s4 == s5)
	*/

	//
	//var s6 *circle
	//fmt.Println("s1 == ?", s1 == getMaxBySpace(s1, s6))
}

// ------------------------------------------

func getMaxBySpace(s1, s2 Spacer) Spacer {
	if s1 == nil && s2 != nil {
		return s2
	}

	if s1 != nil && s2 == nil {
		return s1
	}

	if s1 == nil && s2 == nil {
		return nil
	}

	if s1.Space() > s2.Space() {
		return s1
	}

	return s2
}

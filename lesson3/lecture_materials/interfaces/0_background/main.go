package main

import (
	"fmt"
	"math"
)

// ------------------------------------------
type circle struct {
	radius float64
}

func (s *circle) Space() float64 {
	return math.Pi * math.Pow(s.radius, 2)
}

func getMaxCircleBySpace(s1, s2 *circle) *circle {
	if s1.Space() > s2.Space() {
		return s1
	}

	return s2
}

func main() {
	c1 := &circle{radius: 10}
	c2 := &circle{radius: 12}
	fmt.Println(getMaxCircleBySpace(c1, c2))

	// r1 := &rectangle{width: 10, height: 10}
	// r2 := &rectangle{width: 12, height: 10}
	// fmt.Println(getMaxRectangleBySpace(r1, r2))
}

// ------------------------------------------
type rectangle struct {
	width  float64
	height float64
}

func (s *rectangle) Space() float64 {
	return s.width * s.height
}

func getMaxRectangleBySpace(s1, s2 *rectangle) *rectangle {
	if s1.Space() > s2.Space() {
		return s1
	}

	return s2
}

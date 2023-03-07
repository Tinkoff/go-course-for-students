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

func (s *circle) Space() float64 {
	return math.Pi * math.Pow(s.radius, 2)
}

type rectangle struct {
	width  float64
	height float64
}

func (s *rectangle) Space() float64 {
	return s.width * s.height
}

func (s *rectangle) Perimeter() float64 {
	return 2 * (s.width + s.height)
}

// ------------------------------------------
func getMaxBySpace(s1, s2 Spacer) Spacer {
	if s1.Space() > s2.Space() {
		return s1
	}

	return s2
}

func main() {
	s1 := &circle{radius: 10}
	var s2 Spacer = &rectangle{width: 12, height: 10}
	fmt.Println(getMaxBySpace(s1, s2))
}

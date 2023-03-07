package main

import (
	"fmt"
)

type Spacer interface {
	Space() float64
}

type StringSpacer interface {
	Spacer
	fmt.Stringer
}

// ------------------------------------------
type rectangle struct {
	width  float64
	height float64
}

func (s *rectangle) Space() float64 {
	return s.width * s.height
}

func (s rectangle) String() string {
	return fmt.Sprintf("width: %.2f, height: %.2f", s.width, s.height)
}

// ------------------------------------------
func getMaxBySpace(s1, s2 Spacer) Spacer {
	if s1.Space() > s2.Space() {
		return s1
	}

	return s2
}

func main() {
	var s1 StringSpacer = &rectangle{width: 10, height: 10}
	var s2 StringSpacer = &rectangle{width: 12.12345, height: 10.54321}
	fmt.Println(getMaxBySpace(s1, s2))
}

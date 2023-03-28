package main

//go:generate mockgen -source main.go -package main -destination area_calculator_mock.go
type AreaCalculator interface {
	GetArea() float64
}

type Shape struct {
	r *float64

	areaCalculator AreaCalculator
}

func (s *Shape) Square() float64 {
	return s.areaCalculator.GetArea()
}

func main() {
	// ...
}

package golang_united_school_homework

import (
	"math"
)

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (r Circle) CalcPerimeter() float64 {
	return 2 * math.Pi * r.Radius
}

func (r Circle) CalcArea() float64 {
	return math.Pi * r.Radius * r.Radius
}

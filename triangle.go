package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (r Triangle) CalcPerimeter() float64 {
	return 3 * r.Side
}

func (r Triangle) CalcArea() float64 {
	return math.Sqrt(3) / 4 * r.Side * r.Side
}

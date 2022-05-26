package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Shape
	Radius float64
}

func (c Circle) CalcPerimeter() float64 {
	return math.Pi * c.Radius * 2
}

func (c Circle) CalcArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

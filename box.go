package golang_united_school_homework

import "errors"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return errors.New("full")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("out of range")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("out of range")
	}
	removed := b.shapes[i]
	newShapes := b.shapes[:i]
	for _, shape := range b.shapes[i+1:] {
		newShapes = append(newShapes, shape)
	}
	b.shapes = newShapes
	return removed, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("out of range")
	}
	removed := b.shapes[i]
	b.shapes[i] = shape
	return removed, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() (sum float64) {
	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}
	return
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() (sum float64) {
	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}
	return

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	newShapes := make([]Shape, 0)
	exists := false
	for _, shape := range b.shapes {
		_, ok1 := shape.(Circle)
		_, ok2 := shape.(*Circle)
		if !ok2 && !ok1 {
			newShapes = append(newShapes, shape)
		} else {
			exists = true
		}
	}
	if !exists {
		return errors.New("no circles")
	} else {
		b.shapes = newShapes
		return nil
	}
}

package golang_united_school_homework

import "errors"

var (
	outOfRange = errors.New("out of the shapesCapacity range")
	notExist   = errors.New("shape by index doesn't exist")
)

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
	if b.shapesCapacity <= len(b.shapes) {
		return outOfRange
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= b.shapesCapacity {
		return nil, outOfRange
	}
	if i < 0 || i >= len(b.shapes) {
		return nil, outOfRange
	}
	if b.shapes[i] == nil {
		return nil, notExist
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape, err := b.GetByIndex(i)
	if err == nil {
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	}
	return shape, err
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	shapeToRemove, err := b.GetByIndex(i)
	if err == nil {
		b.shapes[i] = shape
		return shapeToRemove, nil
	}
	return nil, err
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	sum := float64(0)
	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	sum := float64(0)
	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}
	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	shapesWithoutCircles := []Shape{}
	for _, shape := range b.shapes {
		_, ok := shape.(*Circle)
		if !ok {
			shapesWithoutCircles = append(shapesWithoutCircles, shape)
		}
	}
	if len(shapesWithoutCircles) != len(b.shapes) {
		b.shapes = shapesWithoutCircles
		return nil
	}
	return notExist
}

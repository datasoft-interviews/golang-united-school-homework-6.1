package golang_united_school_homework

import (
	"testing"
)

func TestAddShapeToZeroBox(t *testing.T) {
	box := NewBox(0)
	shape := Circle{2.1}
	want_err := outOfRange
	if got_err := box.AddShape(&shape); got_err != want_err {
		t.Errorf("got %t, want %t", got_err, want_err)
	}
}

func TestAddShape(t *testing.T) {
	box := NewBox(1)
	shape := Circle{2.1}
	want_err := error(nil)
	if got_err := box.AddShape(&shape); got_err != want_err {
		t.Errorf("got %t, want %t", got_err, want_err)
	}
}

func TestGetByIndex(t *testing.T) {
	box := NewBox(2)
	var shape1 Shape = Circle{2.1}
	box.AddShape(shape1)
	var shape2 Shape = &Rectangle{2.1, 3.4}
	box.AddShape(shape2)
	want_shape, want_err := shape1, error(nil)
	if got_shape, got_err := box.GetByIndex(0); got_shape != want_shape || got_err != want_err {
		t.Errorf("got %t (%t), want %t (%t)", got_shape, got_err, want_shape, want_err)
	}
	want_shape, want_err = shape2, error(nil)
	if got_shape, got_err := box.GetByIndex(1); got_shape != want_shape || got_err != want_err {
		t.Errorf("got %t (%t), want %t (%t)", got_shape, got_err, want_shape, want_err)
	}
}

func TestGetByNegativeIndex(t *testing.T) {
	box := NewBox(1)
	want_shape, want_err := Shape(nil), outOfRange
	if got_shape, got_err := box.GetByIndex(-1); got_shape != want_shape || got_err != want_err {
		t.Errorf("got %t (%t), want %t (%t)", got_shape, got_err, want_shape, want_err)
	}
}

func TestGetByWrongIndex(t *testing.T) {
	box := NewBox(1)
	want_shape, want_err := Shape(nil), outOfRange
	if got_shape, got_err := box.GetByIndex(1); got_shape != want_shape || got_err != want_err {
		t.Errorf("got %t (%t), want %t (%t)", got_shape, got_err, want_shape, want_err)
	}
}

func TestExtractByIndex(t *testing.T) {
	box := NewBox(1)
	var shape Shape = Circle{2.1}
	box.AddShape(shape)
	want_shape, want_err := shape, error(nil)
	if got_shape, got_err := box.ExtractByIndex(0); got_shape != want_shape && got_err != want_err {
		t.Errorf("got %t (%t), want %t (%t)", got_shape, got_err, want_shape, want_err)
	}
	want_shape, want_err = Shape(nil), outOfRange
	if got_shape, got_err := box.ExtractByIndex(0); got_shape != want_shape || got_err != want_err {
		t.Errorf("got %t (%t), want %t (%t)", got_shape, got_err, want_shape, want_err)
	}
}

func TestGetByIndexAfterExtractByIndex(t *testing.T) {
	box := NewBox(1)
	var shape Shape = Circle{2.1}
	box.AddShape(shape)
	want_shape, want_err := shape, error(nil)
	if got_shape, got_err := box.ExtractByIndex(0); got_shape != want_shape || got_err != want_err {
		t.Errorf("got %t (%t), want %t (%t)", got_shape, got_err, want_shape, want_err)
	}
	want_shape, want_err = Shape(nil), outOfRange
	if got_shape, got_err := box.GetByIndex(0); got_shape != want_shape || got_err != want_err {
		t.Errorf("got %t (%t), want %t (%t)", got_shape, got_err, want_shape, want_err)
	}
}

func TestReplaceByIndex(t *testing.T) {
	box := NewBox(1)
	var shape1 Shape = Circle{2.1}
	var shape2 Shape = Rectangle{2.1, 3.1}
	box.AddShape(shape1)
	want_shape, want_err := shape1, error(nil)
	if got_shape, got_err := box.ReplaceByIndex(0, shape2); got_shape != want_shape || got_err != want_err {
		t.Errorf("got %t (%t), want %t (%t)", got_shape, got_err, want_shape, want_err)
	}
	want_shape, want_err = shape2, nil
	if got_shape, got_err := box.GetByIndex(0); got_shape != want_shape || got_err != want_err {
		t.Errorf("got %t (%t), want %t (%t)", got_shape, got_err, want_shape, want_err)
	}
}

func TestSumPerimeter(t *testing.T) {
	box := NewBox(2)
	var shape1 Shape = Rectangle{4, 5}
	var shape2 Shape = Rectangle{2, 3}
	box.AddShape(shape1)
	box.AddShape(shape2)
	want := float64(18 + 10)
	if got := box.SumPerimeter(); got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

func TestSumArea(t *testing.T) {
	box := NewBox(2)
	var shape1 Shape = Rectangle{4, 5}
	var shape2 Shape = Rectangle{2, 3}
	box.AddShape(shape1)
	box.AddShape(shape2)
	want := float64(20 + 6)
	if got := box.SumArea(); got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

func TestRemoveAllCircles(t *testing.T) {
	box := NewBox(2)
	var shape1 Shape = &Circle{4}
	var shape2 Shape = &Rectangle{2, 3}
	box.AddShape(shape1)
	box.AddShape(shape2)
	want := error(nil)
	if got := box.RemoveAllCircles(); got != want {
		t.Errorf("got %f, want %f", got, want)
	}
	want = notExist
	if got := box.RemoveAllCircles(); got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

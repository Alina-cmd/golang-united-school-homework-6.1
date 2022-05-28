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
		shapes:         make([]Shape, 0, shapesCapacity),
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	}

	return errors.New("it goes out of the shapesCapacity range")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if b.shapesCapacity <= i {
		return nil, errors.New("index went out of the range")
	} else if i > len(b.shapes)-1 {
		return nil, errors.New("shape by index doesn't exist")
	}
	return b.shapes[i], nil

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if b.shapesCapacity <= i {
		return nil, errors.New("index went out of the range")
	} else if i > len(b.shapes)-1 {
		return nil, errors.New("shape by index doesn't exist")
	}
	sh := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return sh, nil

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if b.shapesCapacity <= i {
		return nil, errors.New("index went out of the range")
	} else if i > len(b.shapes)-1 {
		return nil, errors.New("shape by index doesn't exist")
	}
	sh := b.shapes[i]
	b.shapes[i] = shape
	return sh, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, val := range b.shapes {
		sum += val.CalcPerimeter()
	}
	return sum

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, val := range b.shapes {
		sum += val.CalcArea()
	}
	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	exist := false
	index := 0
	for i, val := range b.shapes {
		_, ok := val.(Circle)
		if ok {
			exist = true
			index++
			b.shapes = append(b.shapes[:i-index], b.shapes[i-index+1:]...)
		}
	}
	if !exist {
		return errors.New("circles are not exist in the list")
	}

	return nil
}

package rectangle

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(w, h float64) Rectangle {
	return Rectangle{
		width:  w,
		height: h,
	}
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Info(res float64) string {
	return fmt.Sprintf("Прямоугольник: ширина %v, высота %v\nПлощадь: %v\n", r.Width(), r.Height(), res)
}

func (r Rectangle) Width() float64 {
	return r.width
}

func (r Rectangle) Height() float64 {
	return r.height
}

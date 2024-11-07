package triangle

import "fmt"

type Triangle struct {
	base   float64
	height float64
}

func NewTriangle(base, height float64) Triangle {
	return Triangle{
		base:   base,
		height: height,
	}
}

func (t Triangle) Area() float64 {
	return 0.5 * t.height * t.base
}

func (t Triangle) Info(res float64) string {
	return fmt.Sprintf("Треугольник: основание %v, высота %v\nПлощадь: %v\n", t.Base(), t.Height(), res)
}

func (t Triangle) Base() float64 {
	return t.base
}

func (t Triangle) Height() float64 {
	return t.height
}

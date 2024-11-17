package circle

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func NewCircle(r float64) Circle {
	return Circle{
		radius: r,
	}
}

func (c Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (c Circle) Info(res float64) string {
	return fmt.Sprintf("Круг: радиус %v\nПлощадь: %v\n", c.Radius(), res)
}

func (c Circle) Radius() float64 {
	return c.radius
}

func (c Circle) IsValid() bool {
	return c.radius > 0
}

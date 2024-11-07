package circle

import (
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

func (c Circle) Radius() float64 {
	return c.radius
}

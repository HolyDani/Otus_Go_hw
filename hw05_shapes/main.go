package main

import (
	"fmt"

	shapes "github.com/Otus_hw/HolyDani/hw05_shapes/interface"
	"github.com/Otus_hw/HolyDani/hw05_shapes/shapes/circle"
	"github.com/Otus_hw/HolyDani/hw05_shapes/shapes/rectangle"
	"github.com/Otus_hw/HolyDani/hw05_shapes/shapes/triangle"
)

func main() {
	circle1 := circle.NewCircle(5)
	rectangle1 := rectangle.NewRectangle(10, 5)
	triangle1 := triangle.NewTriangle(8, 6)
	res, _ := calculateArea(circle1)
	fmt.Printf("Круг: радиус %v\nПлощадь: %v\n\n", circle1.Radius(), res)
	res, _ = calculateArea(rectangle1)
	fmt.Printf("Прямоугольник: ширина %v, высота %v\nПлощадь: %v\n\n", rectangle1.Width(), rectangle1.Height(), res)
	res, _ = calculateArea(triangle1)
	fmt.Printf("Треугольник: основание %v, высота %v\nПлощадь: %v\n\n", triangle1.Base(), triangle1.Height(), res)
}

func calculateArea(s shapes.Shape) (float64, string) {
	if s == nil {
		return 0, "Переданный объект не является фигурой"
	}
	return s.Area(), ""
}

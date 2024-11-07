package main

import (
	"errors"
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
	res, err := calculateArea(circle1)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	} else {
		fmt.Println(printInfo(circle1, res))
	}

	res, err = calculateArea(rectangle1)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	} else {
		fmt.Println(printInfo(rectangle1, res))
	}

	res, err = calculateArea(triangle1)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	} else {
		fmt.Println(printInfo(triangle1, res))
	}
}

func calculateArea(s shapes.Shape) (float64, error) {
	if s == nil {
		return 0, errors.New("переданный объект не является фигурой")
	}
	return s.Area(), nil
}

func printInfo(s shapes.Shape, res float64) string {
	return s.Info(res)
}

package main

import (
	"errors"
	"math"
	"testing"

	shapes "github.com/Otus_hw/HolyDani/hw06_testing/hw05/interface"
	"github.com/Otus_hw/HolyDani/hw06_testing/hw05/shapes/circle"
	"github.com/Otus_hw/HolyDani/hw06_testing/hw05/shapes/rectangle"
	"github.com/Otus_hw/HolyDani/hw06_testing/hw05/shapes/triangle"
	"github.com/stretchr/testify/assert"
)

type emptyShape struct{}

func (e emptyShape) Area() float64 {
	return 0
}

func (e emptyShape) Info(_ float64) string {
	return ""
}

func (e emptyShape) IsValid() bool {
	return false
}

func TestCalculateArea(t *testing.T) {
	tests := []struct {
		shape        shapes.Shape
		expectedArea float64
		expectedErr  error
	}{
		{
			shape:        circle.NewCircle(5),
			expectedArea: 5 * 5 * math.Pi,
			expectedErr:  nil,
		},
		{
			shape:        rectangle.NewRectangle(10, 5),
			expectedArea: 50,
			expectedErr:  nil,
		},
		{
			shape:        triangle.NewTriangle(8, 3),
			expectedArea: 8 * 3 * 0.5,
			expectedErr:  nil,
		},
		{
			shape:        emptyShape{},
			expectedArea: 0,
			expectedErr:  errors.New("переданный объект не является фигурой"),
		},
	}
	for _, test := range tests {
		got, err := calculateArea(test.shape)
		assert.Equal(t, test.expectedArea, got)
		assert.Equal(t, test.expectedErr, err)
	}
}

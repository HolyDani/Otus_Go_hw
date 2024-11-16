package main

import (
	"errors"
	"math"
	"testing"

	shapes "github.com/Otus_hw/HolyDani/hw05_shapes/interface"
	"github.com/Otus_hw/HolyDani/hw05_shapes/shapes/circle"
	"github.com/Otus_hw/HolyDani/hw05_shapes/shapes/rectangle"
	"github.com/Otus_hw/HolyDani/hw05_shapes/shapes/triangle"
	"github.com/stretchr/testify/assert"
)

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
			shape:        nil,
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

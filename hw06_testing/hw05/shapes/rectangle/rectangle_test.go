package rectangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRectangle(t *testing.T) {
	tests := []struct {
		width    float64
		height   float64
		expected Rectangle
	}{
		{
			width:    5,
			height:   8,
			expected: Rectangle{width: 5, height: 8},
		},
		{
			width:    0,
			height:   0,
			expected: Rectangle{width: 0, height: 0},
		},
		{
			width:    0,
			height:   4,
			expected: Rectangle{width: 0, height: 4},
		},
		{
			width:    5,
			height:   0,
			expected: Rectangle{width: 5, height: 0},
		},
		{
			width:    5.45,
			height:   9.70,
			expected: Rectangle{width: 5.45, height: 9.70},
		},
	}
	for _, test := range tests {
		rect := NewRectangle(test.width, test.height)
		assert.Equal(t, test.expected.width, rect.Width())
		assert.Equal(t, test.expected.height, rect.Height())
	}
}

func TestArea_Rectangle(t *testing.T) {
	tests := []struct {
		width    float64
		height   float64
		expected float64
	}{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{5, 10, 50},
		{7.45, 2.2, 7.45 * 2.2},
	}
	for _, test := range tests {
		rect := NewRectangle(test.width, test.height)
		got := rect.Area()
		assert.Equal(t, test.expected, got)
	}
}

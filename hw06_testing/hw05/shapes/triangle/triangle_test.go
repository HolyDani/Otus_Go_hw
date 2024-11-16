package triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTriangle(t *testing.T) {
	tests := []struct {
		base     float64
		height   float64
		expected Triangle
	}{
		{
			base:     8,
			height:   5,
			expected: Triangle{base: 8, height: 5},
		},
		{
			base:     0,
			height:   0,
			expected: Triangle{base: 0, height: 0},
		},
		{
			base:     0,
			height:   8.5,
			expected: Triangle{base: 0, height: 8.5},
		},
	}
	for _, test := range tests {
		triangle := NewTriangle(test.base, test.height)
		assert.Equal(t, test.expected.base, triangle.Base())
		assert.Equal(t, test.expected.height, triangle.Height())
	}
}

func TestArea_Triangle(t *testing.T) {
	tests := []struct {
		base     float64
		height   float64
		expected float64
	}{
		{8, 6, 8 * 6 * 0.5},
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{7.2, 3.5, 7.2 * 3.5 * 0.5},
	}
	for _, test := range tests {
		triangle := NewTriangle(test.base, test.height)
		got := triangle.Area()
		assert.Equal(t, test.expected, got)
	}
}

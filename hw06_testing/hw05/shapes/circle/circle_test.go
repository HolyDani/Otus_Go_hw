package circle

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCircle(t *testing.T) {
	tests := []struct {
		radius   float64
		expected Circle
	}{
		{
			radius:   5,
			expected: Circle{radius: 5},
		},
		{
			radius:   0,
			expected: Circle{radius: 0},
		},
		{
			radius:   24.7569,
			expected: Circle{radius: 24.7569},
		},
	}
	for _, test := range tests {
		got := NewCircle(test.radius)
		assert.Equal(t, test.expected.radius, got.Radius())
	}
}

func TestArea_Circle(t *testing.T) {
	tests := []struct {
		name     string
		radius   float64
		expected float64
	}{
		{
			name:     "radius with value 0",
			radius:   0,
			expected: 0,
		},
		{
			name:     "radius with value 1",
			radius:   1,
			expected: 1 * math.Pi,
		},
		{
			name:     "radius with value 5",
			radius:   5,
			expected: 25 * math.Pi,
		},
		{
			name:     "radius with value 3,75",
			radius:   3.75,
			expected: (3.75 * 3.75) * math.Pi,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			circle := Circle{test.radius}
			got := circle.Area()
			assert.Equal(t, test.expected, got)
		})
	}
}

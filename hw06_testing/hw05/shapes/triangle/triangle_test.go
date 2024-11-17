package triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

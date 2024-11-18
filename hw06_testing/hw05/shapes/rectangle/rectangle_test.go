package rectangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

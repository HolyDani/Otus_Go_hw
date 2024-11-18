package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateBoard(t *testing.T) {
	tests := []struct {
		size     int
		expected string
	}{
		{1, "#\n"},
		{2, "# \n #\n"},
		{3, "# #\n # \n# #\n"},
		{4, "# # \n # #\n# # \n # #\n"},
		{0, ""},
		{-1, ""},
	}

	for _, test := range tests {
		got := GenerateBoard(test.size)
		assert.Equal(t, test.expected, got)
	}
}

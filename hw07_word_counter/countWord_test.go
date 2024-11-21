package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWord(t *testing.T) {
	tests := []struct {
		input          string
		expectedResult map[string]uint
		expectedErr    error
	}{
		{
			input: "Hello my friend. You are my best friend!",
			expectedResult: map[string]uint{
				"hello":  1,
				"my":     2,
				"friend": 2,
				"you":    1,
				"are":    1,
				"best":   1,
			},
			expectedErr: nil,
		},
		{
			input:          "",
			expectedResult: nil,
			expectedErr:    fmt.Errorf("%w", ErrEmptyString),
		},
		{
			input:          "                ",
			expectedResult: nil,
			expectedErr:    fmt.Errorf("%w", ErrEmptyString),
		},
		{
			input: "      Hello    ,   this       is     my            swamp",
			expectedResult: map[string]uint{
				"hello": 1,
				"this":  1,
				"is":    1,
				"my":    1,
				"swamp": 1,
			},
			expectedErr: nil,
		},
		{
			input: "This is my 1st dog. I am 27!",
			expectedResult: map[string]uint{
				"this": 1,
				"is":   1,
				"my":   1,
				"1st":  1,
				"dog":  1,
				"i":    1,
				"am":   1,
				"27":   1,
			},
			expectedErr: nil,
		},
	}
	for _, test := range tests {
		text := test.input
		got, err := countWords(text)
		assert.Equal(t, test.expectedResult, got)
		assert.Equal(t, test.expectedErr, err)
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		sl             []int
		item           int
		expectedResult int
	}{
		{
			sl:             []int{1, 5, 8, 12, 25},
			item:           12,
			expectedResult: 3,
		},
		{
			sl:             []int{1, 5, 8, 12, 25},
			item:           13,
			expectedResult: -1,
		},
		{
			sl:             []int{},
			item:           5,
			expectedResult: -1,
		},
		{
			sl:             []int{5, 9, 4, 7, 1},
			item:           9,
			expectedResult: 4,
		},
	}
	for _, test := range tests {
		sortedSl := Qsort(test.sl)
		got := BinarySearch(test.item, sortedSl)
		assert.Equal(t, test.expectedResult, got)
	}
}

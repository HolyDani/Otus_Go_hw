package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareBooks(t *testing.T) {
	tests := []struct {
		name     string
		book1    *Book
		book2    *Book
		comp     Compare
		expected bool
	}{
		{
			name:     "Year compare: book2 newer than book1",
			book1:    NewBook(1, 2023, 220, 4.5, "", ""),
			book2:    NewBook(2, 2024, 240, 3.9, "", ""),
			comp:     MakeCompare(Year),
			expected: false,
		},
		{
			name:     "Year compare: book1 newer than book2",
			book1:    NewBook(1, 2024, 240, 4.5, "", ""),
			book2:    NewBook(2, 2023, 200, 3.9, "", ""),
			comp:     MakeCompare(Year),
			expected: true,
		},
		{
			name:     "Year compare: if both param equal book1 newer than book2",
			book1:    NewBook(1, 2024, 240, 4.5, "", ""),
			book2:    NewBook(2, 2024, 200, 3.9, "", ""),
			comp:     MakeCompare(Year),
			expected: true,
		},
		{
			name:     "Size compare: book2 larger than book1",
			book1:    NewBook(1, 2020, 239, 4.5, "", ""),
			book2:    NewBook(2, 2024, 240, 3.9, "", ""),
			comp:     MakeCompare(Size),
			expected: false,
		},
		{
			name:     "Size compare: book1 larger than book2",
			book1:    NewBook(1, 2020, 240, 4.5, "", ""),
			book2:    NewBook(2, 2024, 239, 3.9, "", ""),
			comp:     MakeCompare(Size),
			expected: true,
		},
		{
			name:     "Size compare: if both param equal, book1 larger than book2",
			book1:    NewBook(1, 2020, 240, 4.5, "", ""),
			book2:    NewBook(2, 2024, 240, 3.9, "", ""),
			comp:     MakeCompare(Size),
			expected: true,
		},
		{
			name:     "Rate compare: book1 popular than book2",
			book1:    NewBook(1, 2020, 240, 4.1, "", ""),
			book2:    NewBook(2, 2024, 240, 4.0, "", ""),
			comp:     MakeCompare(Rate),
			expected: true,
		},
		{
			name:     "Rate compare: book2 popular than book1",
			book1:    NewBook(1, 2020, 240, 4.0, "", ""),
			book2:    NewBook(2, 2024, 240, 4.1, "", ""),
			comp:     MakeCompare(Rate),
			expected: false,
		},
		{
			name:     "Rate compare: if both param equal, book1 popular than book2",
			book1:    NewBook(1, 2020, 240, 4.5, "", ""),
			book2:    NewBook(2, 2024, 240, 4.5, "", ""),
			comp:     MakeCompare(Rate),
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.comp.CompareBooks(test.book1, test.book2)
			assert.Equal(t, test.expected, got)
		})
	}
}

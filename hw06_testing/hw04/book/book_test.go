package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBook(t *testing.T) {
	tests := []struct {
		id           uint
		year         uint
		size         uint
		rate         float32
		author       string
		title        string
		expectedBook *Book
	}{
		{
			id:           1,
			year:         2024,
			size:         350,
			rate:         4.3,
			author:       "Author A",
			title:        "Book A",
			expectedBook: &Book{id: 1, year: 2024, size: 350, rate: 4.3, author: "Author A", title: "Book A"},
		},
		{
			id:           5,
			year:         1991,
			size:         215,
			rate:         5.0,
			author:       "Author B",
			title:        "Book B",
			expectedBook: &Book{id: 5, year: 1991, size: 215, rate: 5.0, author: "Author B", title: "Book B"},
		},
	}
	for _, test := range tests {
		got := NewBook(test.id, test.year, test.size, test.rate, test.author, test.title)
		assert.Equal(t, test.expectedBook.id, *got.ID())
		assert.Equal(t, test.expectedBook.year, *got.Year())
		assert.Equal(t, test.expectedBook.size, *got.Size())
		assert.Equal(t, test.expectedBook.rate, *got.Rate())
		assert.Equal(t, test.expectedBook.author, *got.Author())
		assert.Equal(t, test.expectedBook.title, *got.Title())
	}
}

func TestSetters(t *testing.T) {
	tests := []struct {
		initial      Book
		setID        uint
		setYear      uint
		setSize      uint
		setRate      float32
		setAuthor    string
		setTitle     string
		expectedBook Book
	}{
		{
			initial: Book{1, 2024, 350, 4.3, "Book A", "Author A"},
			setID:   4, setYear: 2000, setSize: 150, setRate: 4.9, setAuthor: "A Author", setTitle: "A Book",
			expectedBook: Book{4, 2000, 150, 4.9, "A Book", "A Author"},
		},
		{
			initial: Book{5, 2023, 30, 3.2, "", ""},
			setID:   10, setYear: 1999, setSize: 190, setRate: 2.9, setAuthor: "B Author", setTitle: "B Book",
			expectedBook: Book{10, 1999, 190, 2.9, "B Book", "B Author"},
		},
	}
	for _, test := range tests {
		book := test.initial
		book.SetID(test.setID)
		book.SetYear(test.setYear)
		book.SetSize(test.setSize)
		book.SetRate(test.setRate)
		book.SetAuthor(test.setAuthor)
		book.SetTitle(test.setTitle)
		assert.Equal(t, test.expectedBook, book)
	}
}

func TestMakeCompare(t *testing.T) {
	tests := []struct {
		operation     State
		expectedValue State
	}{
		{Year, 1},
		{Size, 2},
		{Rate, 3},
	}
	for _, test := range tests {
		assert.Equal(t, test.expectedValue, test.operation)
	}
}

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

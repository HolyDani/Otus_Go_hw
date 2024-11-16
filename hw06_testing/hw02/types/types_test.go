package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	tests := []struct {
		input    Employee
		expexted string
	}{
		{
			input:    Employee{UserID: 1, Age: 27, Name: "Alexander", DepartmentID: 100},
			expexted: "User ID: 1; Age: 27; Name: Alexander; Department ID: 100; ",
		},
		{
			input:    Employee{UserID: 12, Age: 42, Name: "Petr", DepartmentID: 111},
			expexted: "User ID: 12; Age: 42; Name: Petr; Department ID: 111; ",
		},
		{
			input:    Employee{UserID: -10, Age: 0, Name: "", DepartmentID: 999},
			expexted: "User ID: -10; Age: 0; Name: ; Department ID: 999; ",
		},
	}
	for _, test := range tests {
		got := test.input.String()
		assert.Equal(t, test.expexted, got)
	}
}

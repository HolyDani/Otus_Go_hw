package reader

import (
	"errors"
	"testing"

	"github.com/fixme_my_friend/hw02_fix_app/types"
	"github.com/stretchr/testify/assert"
)

func TestReadJSON(t *testing.T) {
	tests := []struct {
		filePath    string
		expected    []types.Employee
		expectedErr error
	}{
		{
			filePath: "test_data/valid_path.json",
			expected: []types.Employee{
				{UserID: 1, Name: "Bob", Age: 20, DepartmentID: 5},
				{UserID: 2, Name: "Greg", Age: 42, DepartmentID: 8},
			},
			expectedErr: nil,
		},
		{
			filePath:    "test_data/invalid_path.json",
			expected:    nil,
			expectedErr: errors.New("readJSON: open test_data/invalid_path.json: no such file or directory"),
		},
		{
			filePath:    "test_data/empty_path.json",
			expected:    nil,
			expectedErr: errors.New("unmarshal JSON: unexpected end of JSON input"),
		},
	}
	for _, test := range tests {
		got, err := ReadJSON(test.filePath)
		assert.Equal(t, test.expected, got)
		assert.Equal(t, test.expectedErr, err)
	}
}

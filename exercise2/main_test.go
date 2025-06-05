package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapDetail(t *testing.T) {
	var tests = map[string]struct {
		matrix                [][]int
		expectedMaxColumnSize int
		expectedIsSquare      bool
	}{
		"case 1": {
			matrix:                [][]int{{1, 2, 3}, {4, 5, 4}, {1, 1, 1}, {7, 8, 9}},
			expectedMaxColumnSize: 3,
			expectedIsSquare:      true,
		},
		"case 2": {
			matrix:                [][]int{{1}, {4, 5, 4, 5, 5, 1}, {}, {7, 8, 9}},
			expectedMaxColumnSize: 6,
			expectedIsSquare:      false,
		},
		"case 3": {
			matrix:                [][]int{{1, 2, 3, 4, 5}},
			expectedMaxColumnSize: 5,
			expectedIsSquare:      false,
		},
		"case 4": {
			matrix:                [][]int{{1, 2}, {3, 4}},
			expectedMaxColumnSize: 2,
			expectedIsSquare:      true,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			// TODO: delete the line below
			assert.Equal(t, testCase.expectedMaxColumnSize, testCase.expectedMaxColumnSize)

			// TODO: replace the code below with the actual test
			// md := NewMapDetailImpl(testCase.matrix)
			// assert.Equal(t, testCase.expectedMaxColumnSize, md.getMaxColumnSize())
			// assert.Equal(t, testCase.expectedIsSquare, md.isSquare())
		})
	}
}

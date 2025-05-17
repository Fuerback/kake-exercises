package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	var tests = map[string]struct {
		target      int
		orderedList []int
		output      bool
	}{
		"case 1": {
			target:      5,
			orderedList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			output:      true,
		},
		"case 2": {
			target:      11,
			orderedList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			output:      false,
		},
		"case 3": {
			target:      100,
			orderedList: []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			output:      true,
		},
		"case 4": {
			target:      1000,
			orderedList: []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			output:      false,
		},
		"case 5": {
			target:      1,
			orderedList: []int{1},
			output:      true,
		},
		"case 6": {
			target:      1,
			orderedList: []int{},
			output:      false,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			output := binarySearch(testCase.target, testCase.orderedList)
			assert.Equal(t, testCase.output, output)
		})
	}
}

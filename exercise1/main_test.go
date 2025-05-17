package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumColValues(t *testing.T) {
	var tests = map[string]struct {
		input  [][]int
		output []int
	}{
		"case 1": {
			input: [][]int{
				{1, 2, 3},
				{4, 5},
				{},
				{7, 8, 9},
			},
			output: []int{12, 15, 12},
		},
		"case 2": {
			input: [][]int{
				{1, 0, 3, 12, 100},
				{4, 5, 6, 7, 8},
				{},
				{1},
				{7, 8, 9, 0, -1},
			},
			output: []int{13, 13, 18, 19, 107},
		},
		"case 3": {
			input: [][]int{
				{1},
				{},
			},
			output: []int{1},
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			output := sumColValues(testCase.input)
			assert.Equal(t, testCase.output, output)
		})
	}
}

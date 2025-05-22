package main

import "fmt"

/*
Given a 2D array of integers, write a function that returns a 1D array containing the sum of each column.

In other words, you should sum all the values of the first index of each array and return it as the first index of the output array. Do the same for the second index, and so on.

Implement the function `sumColValues`.

## Example

```go
input := [][]int{
	{1, 2, 3},
	{4, 5},
	{},
	{7, 8, 9},
}

output := []int{12, 15, 12}
```
*/

func sumColValues(data [][]int) []int {
	return []int{}
}

func main() {
	runTests()
}

func runTests() {
	testCases := []struct {
		name   string
		input  [][]int
		output []int
	}{
		{
			name: "case 1",
			input: [][]int{
				{1, 2, 3},
				{4, 5},
				{},
				{7, 8, 9},
			},
			output: []int{12, 15, 12},
		},
		{
			name: "case 2",
			input: [][]int{
				{1, 0, 3, 12, 100},
				{4, 5, 6, 7, 8},
				{},
				{1},
				{7, 8, 9, 0, -1},
			},
			output: []int{13, 13, 18, 19, 107},
		},
		{
			name: "case 3",
			input: [][]int{
				{1},
				{},
			},
			output: []int{1},
		},
	}

	passed := true
	for _, tc := range testCases {
		output := sumColValues(tc.input)
		if !slicesEqual(output, tc.output) {
			fmt.Printf("❌ Test %s failed\n", tc.name)
			fmt.Printf("   Expected: %v\n", tc.output)
			fmt.Printf("   Got:      %v\n", output)
			passed = false
		} else {
			fmt.Printf("✅ Test %s passed\n", tc.name)
		}
	}

	if passed {
		fmt.Println("\nAll tests passed! 🎉")
	} else {
		fmt.Println("\nSome tests failed 😢")
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

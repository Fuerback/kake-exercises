package main

import "fmt"

/*
Given a sorted list of integers, implement binary search the target value.

The function should return true if the target value is found in the list, and false otherwise.

### Example

```go
binarySearch(5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
```

```go
true
```

```go
binarySearch(11, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
```

```go
false
```
*/

func binarySearch(target int, list []int) bool {
	return false
}

func main() {
	runTests()
}

func runTests() {
	testCases := []struct {
		name        string
		target      int
		orderedList []int
		output      bool
	}{
		{
			name:        "case 1",
			target:      5,
			orderedList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			output:      true,
		},
		{
			name:        "case 2",
			target:      11,
			orderedList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			output:      false,
		},
		{
			name:        "case 3",
			target:      100,
			orderedList: []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			output:      true,
		},
		{
			name:        "case 4",
			target:      1000,
			orderedList: []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			output:      false,
		},
		{
			name:        "case 5",
			target:      1,
			orderedList: []int{1},
			output:      true,
		},
		{
			name:        "case 6",
			target:      1,
			orderedList: []int{},
			output:      false,
		},
	}

	passed := true
	for _, tc := range testCases {
		output := binarySearch(tc.target, tc.orderedList)
		if output != tc.output {
			fmt.Printf("❌ Test %s failed\n", tc.name)
			fmt.Printf("   Target: %d, List: %v\n", tc.target, tc.orderedList)
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

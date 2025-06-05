package main

import "fmt"

func main() {
	data := [][]int{
		{1, 2, 3},
		{4, 5},
		{},
		{7, 8, 9},
	}

	fmt.Println(sumColValues(data))
}

func sumColValues(data [][]int) []int {
	return []int{}
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	// An array of integers to be sorted
	arr := []int{4, 1, 3, 2}

	// Sort the array in ascending order
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	// Create a new array with the sorted data

	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	fmt.Println(sortedArr)
	// The ne array will now be [1, 2, 3, 4]
}

package main

import "fmt"

func main() {
	fmt.Println(FindMinimumIntInSlice([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})) // Output: 1
	fmt.Println(FindMinimumIntInSlice([]int{5, 4, 3, 2, 9, 5, 8, 3, 6, 7, -1, 7, 8, 3, 4, 6, 2, 3})) // Output: -1
	fmt.Println(FindMinimumIntInSlice([]int{-1, -2, -3, -1337, -4, -5, -6, -7, -8, -9})) // Output: -1337
}

// FindMinimumIntInSlice finds the minimum of an integer slice
func FindMinimumIntInSlice(s []int) int {
	var min int
	for i, v := range s {
		if i == 0 || v < min {
			min = v
		}
	}
	return min
}

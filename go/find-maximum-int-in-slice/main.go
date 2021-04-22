package main

import "fmt"

func main() {
	fmt.Println(FindMaximumIntInSlice([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))                             // Output: 9
	fmt.Println(FindMaximumIntInSlice([]int{5, 4, 3, 2, 9, 5, 8, 3, 6, 7, -1, 7, 8, 3, 4, 6, 2, 3})) // Output: 9
	fmt.Println(FindMaximumIntInSlice([]int{-1, -2, -3, -1337, -4, -5, -6, -7, -8, -9}))             // Output: -1
}

// FindMaximumIntInSlice finds the maximum of an integer slice
func FindMaximumIntInSlice(s []int) int {
	var max int
	for i, v := range s {
		if i == 0 || v > max {
			max = v
		}
	}
	return max
}

package main

import "fmt"

func main() {
	fmt.Println(CalculateAverageInt([]int{1, 2}))                        // Output: 1.5
	fmt.Println(CalculateAverageInt([]int{0, 10}))                       // Output: 5
	fmt.Println(CalculateAverageInt([]int{3, 5, 6, 78, 3, 2, 634, 231})) // Output: 120.25
}

// CalculateAverageInt returns the average of an integer slice
func CalculateAverageInt(s []int) float64 {
	var total int
	for _, v := range s {
		total += v
	}
	return float64(total) / float64(len(s))
}

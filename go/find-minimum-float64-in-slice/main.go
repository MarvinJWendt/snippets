package main

import "fmt"

func main() {
	var s []float64
	s = []float64{0.9, 0.8, 0.7, 0.6, 0.5, 0.4, 0.3, 0.2, 0.1}
	fmt.Println(FindMinimumFloat64InSlice(s)) // Output: 0.1

	s = []float64{0.5, 0.4, 0.3, 0.2, 0.9, 0.5, 0.8, 0.3, 0.6, 0.7, -0.1, 0.7, 0.8, 0.3, 0.4, 0.6, 0.2, 0.3}
	fmt.Println(FindMinimumFloat64InSlice(s)) // Output: -0.1

	s = []float64{-0.1, -0.2, -0.3, -0.1337, -0.4, -0.5, -0.6, -0.7, -0.8, -0.9}
	fmt.Println(FindMinimumFloat64InSlice(s)) // Output: -0.9
}

// FindMinimumFloat64InSlice finds the minimum of an float64 slice
func FindMinimumFloat64InSlice(s []float64) float64 {
	var min float64
	for i, v := range s {
		if i == 0 || v < min {
			min = v
		}
	}
	return min
}

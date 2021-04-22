package main

import "fmt"

func main() {
	var s []float32
	s = []float32{0.9, 0.8, 0.7, 0.6, 0.5, 0.4, 0.3, 0.2, 0.1}
	fmt.Println(FindMinimumFloat32InSlice(s)) // Output: 0.1

	s = []float32{0.5, 0.4, 0.3, 0.2, 0.9, 0.5, 0.8, 0.3, 0.6, 0.7, -0.1, 0.7, 0.8, 0.3, 0.4, 0.6, 0.2, 0.3}
	fmt.Println(FindMinimumFloat32InSlice(s))// Output: -0.1

	s = []float32{-0.1, -0.2, -0.3, -0.1337, -0.4, -0.5, -0.6, -0.7, -0.8, -0.9}
	fmt.Println(FindMinimumFloat32InSlice(s)) // Output: -0.9
}

// FindMinimumFloat32InSlice finds the minimum of a float32 slice
func FindMinimumFloat32InSlice(s []float32) float32 {
	var min float32
	for i, v := range s {
		if i == 0 || v < min {
			min = v
		}
	}
	return min
}

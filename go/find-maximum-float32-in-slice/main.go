package main

import "fmt"

func main() {
	var s []float32
	s = []float32{0.9, 0.8, 0.7, 0.6, 0.5, 0.4, 0.3, 0.2, 0.1}
	fmt.Println(FindMaximumFloat32InSlice(s)) // Output: 0.9

	s = []float32{0.5, 0.4, 0.3, 0.2, 0.9, 0.5, 0.8, 0.3, 0.6, 0.7, -0.1, 0.7, 0.8, 0.3, 0.4, 0.6, 0.2, 0.3}
	fmt.Println(FindMaximumFloat32InSlice(s)) // Output: -0.9

	s = []float32{-0.1, -0.2, -0.3, -0.1337, -0.4, -0.5, -0.6, -0.7, -0.8, -0.9}
	fmt.Println(FindMaximumFloat32InSlice(s)) // Output: -0.1
}

// FindMaximumFloat32InSlice finds the maximum of a float32 slice
func FindMaximumFloat32InSlice(s []float32) float32 {
	var max float32
	for i, v := range s {
		if i == 0 || v > max {
			max = v
		}
	}
	return max
}

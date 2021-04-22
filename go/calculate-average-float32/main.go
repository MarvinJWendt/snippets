package main

import "fmt"

func main() {
	fmt.Println(CalculateAverageInt([]float32{1.5, 2.5}))                                   // Output: 2
	fmt.Println(CalculateAverageInt([]float32{0.4, 0.6}))                                   // Output: 0.5
	fmt.Println(CalculateAverageInt([]float32{3, 5, 6, 78.234, 3, 2.236, 634.102, 231.20})) // Output: 120.34650421142578
}

// CalculateAverageInt returns the average of a float32 slice
func CalculateAverageInt(s []float32) float64 {
	var total float32
	for _, v := range s {
		total += v
	}
	return float64(total) / float64(len(s))
}

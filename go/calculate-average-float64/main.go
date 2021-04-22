package main

import "fmt"

func main() {
	fmt.Println(CalculateAverageInt([]float64{1.5, 2.5}))                                   // Output: 2
	fmt.Println(CalculateAverageInt([]float64{0.4, 0.6}))                                   // Output: 0.5
	fmt.Println(CalculateAverageInt([]float64{3, 5, 6, 78.234, 3, 2.236, 634.102, 231.20})) // Output: 120.34650421142578
}

// CalculateAverageInt returns the average of a float64 slice
func CalculateAverageInt(s []float64) float64 {
	var total float64
	for _, v := range s {
		total += v
	}
	return total / float64(len(s))
}

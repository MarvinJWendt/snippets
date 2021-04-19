package main

import (
	"fmt"
	"time"
)

func main() {
	one, _ := TimeExecutionAverage(10000, func(i int) error {
		_ = fmt.Sprintf("Hello, %s!", "World")

		return nil
	})

	fmt.Println(one)
}

// TimeExecutionAverage times the average execution time of a function.
func TimeExecutionAverage(count int, f func(i int) error) (time.Duration, error) {
	var total time.Duration
	for i := 0; i < count; i++ {
		start := time.Now()

		err := f(i)
		if err != nil {
			return 0, fmt.Errorf("error while calculating average execution time: %w", err)
		}

		duration := time.Since(start)
		total += duration
	}

	return total / time.Duration(count), nil
}

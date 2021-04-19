package main

import (
	"fmt"
	"time"
)

func main() {
	one, _ := TimeAverageExecutionTime(10000, func(i int) error {
		_ = fmt.Sprintf("Hello, %s!", "World")

		return nil
	})

	fmt.Println(one)
}

// TimeAverageExecutionTime times the average execution time of a function.
func TimeAverageExecutionTime(count int, f func(i int) error) (time.Duration, error) {
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

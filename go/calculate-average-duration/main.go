package main

import (
	"fmt"
	"time"
)

func main() {
	one := time.Minute * 10
	two := time.Hour
	three := time.Second * 300

	fmt.Println(CalculateAverageDuration([]time.Duration{one, two, three}))
}

// CalculateAverageDuration calculates the average of a duration slice.
func CalculateAverageDuration(durations []time.Duration) time.Duration {
	var total time.Duration
	for _, duration := range durations {
		total += duration
	}
	return total / time.Duration(len(durations))
}

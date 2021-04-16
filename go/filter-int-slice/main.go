package main

import "fmt"

func main() {
	unfiltered := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	filtered := FilterIntSlice(unfiltered, func(x int) bool {
		return x >= 3 && x <= 8
	})

	fmt.Printf("Unfiltered: %v\n", unfiltered)
	fmt.Printf("Filtered:   %v\n", filtered)
}

// FilterIntSlice filters an integer slice based on a predicate
func FilterIntSlice(s []int, keep func(x int) bool) []int {
	if len(s) == 0 {
		return s
	}

	var ret []int
	for _, v := range s {
		if keep(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

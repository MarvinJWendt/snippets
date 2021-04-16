package main

import "fmt"

func main() {
	unfiltered := []string{"one", "two", "three", "four", "five", "six"}

	filtered := FilterStringSlice(unfiltered, func(x string) bool {
		return len(x) <= 3 // Filter words which have 3 or less characters.
	})

	fmt.Printf("Unfiltered: %v\n", unfiltered)
	fmt.Printf("Filtered:   %v\n", filtered)
	// Output:
	// Unfiltered: [one two three four five six]
    // Filtered:   [one two six]
}

// FilterStringSlice filters a string slice based on a predicate
func FilterStringSlice(s []string, keep func(x string) bool) []string {
	if len(s) == 0 {
		return s
	}

	var ret []string
	for _, v := range s {
		if keep(v) {
			ret = append(ret, v)
		}
	}

	return ret
}
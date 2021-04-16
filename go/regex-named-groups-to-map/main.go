package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`(?P<first>\w+) (?P<middle>.+) (?P<last>\w+)`)
	groups := RegexNamedGroupsToMap(regex, "Marvin J Wendt")
	fmt.Printf("%#v\n", groups)
}

// RegexNamedGroupsToMap returns a map with the regex group names and their values.
// Regex: (?P<first>\w+) (?P<middle>.+) (?P<last>\w+)
// String: "Marvin J Wendt"
// Returns:
// 	m[""] = "Marvin J Wendt"
//  m["first"] = "Marvin"
//  m["middle"] = "J"
//  m["last"] = "Wendt"
func RegexNamedGroupsToMap(r *regexp.Regexp, s string) map[string]string {
	names := r.SubexpNames()
	result := r.FindAllStringSubmatch(s, -1)
	m := map[string]string{}
	for i, n := range result[0] {
		m[names[i]] = n
	}
	return m
}
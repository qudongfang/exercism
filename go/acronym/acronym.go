// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "regexp"
import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	re := regexp.MustCompile("[[:alpha:]']+")

	p := []string{}
	for _, word := range re.FindAllString(s, -1) {
		p = append(p, string(word[0]))
	}

	return strings.ToUpper(strings.Join(p, ""))
}

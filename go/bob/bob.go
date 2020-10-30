// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	saidSomething := 0 != len(remark)

	meaningful := saidSomething
	if saidSomething {
		meaningful, _ = regexp.MatchString("[[:alpha:]]", remark)
	}
	yell := meaningful && strings.ToUpper(remark) == remark

	question := saidSomething && strings.HasSuffix(remark, "?")

	if !saidSomething {
		return "Fine. Be that way!"
	}

	if yell && question {
		return "Calm down, I know what I'm doing!"
	}

	if yell {
		return "Whoa, chill out!"
	}

	if question {
		return "Sure."
	}

	return "Whatever."
}

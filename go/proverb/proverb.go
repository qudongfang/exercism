// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb contains a solution Proverb
package proverb

import "fmt"

const pattern = "For want of %s %s the %s was lost."

// Proverb given a list of inputs, generate the relevant proverb
func Proverb(rhyme []string) []string {
	proverbs := make([]string, 0, len(rhyme))

	for i := 0; i <= len(rhyme)-2; i++ {
		proverbs = append(proverbs, fmt.Sprintf(pattern, article(rhyme[i]), rhyme[i], rhyme[i+1]))
	}

	if len(rhyme) > 0 {
		proverbs = append(proverbs, fmt.Sprintf("And all for the want of %s %s.", article(rhyme[0]), rhyme[0]))
	}

	return proverbs
}

// article return "a"/"an"
func article(word string) string {
	if 0 == len(word) {
		panic("not a word")
	}

	vowels := map[byte]bool{
		'A': true,
		'a': true,
		'I': true,
		'i': true,
		'O': true,
		'o': true,
		'U': true,
		'u': true,
		'E': true,
		'e': true,
	}

	if _, ok := vowels[word[0]]; ok {
		return "an"
	}

	return "a"
}

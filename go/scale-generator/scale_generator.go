package scale

import (
	"strings"
)

// Scale scale generator
func Scale(tonic string, interval string) []string {
	var chr []string
	switch tonic {
	case "C", "a", "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#":
		chr = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		chr = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	}

	chr = reOrder(chr, strings.Title(tonic))
	if 0 == len(interval) {
		return chr
	}

	steps := map[rune]int{
		'm': 1,
		'M': 2,
		'A': 3,
	}

	r := []string{}
	c := 0

	for _, s := range interval {
		r = append(r, chr[c])
		c = (steps[s] + c) % len(chr)
	}

	return r
}

func reOrder(chr []string, tonic string) []string {
	for i, t := range chr {
		if t == tonic {
			return append(chr[i:], chr[:i]...)
		}
	}

	panic("invalid input")
}

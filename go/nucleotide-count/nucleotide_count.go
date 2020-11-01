package dna

import "errors"
import "unicode"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]uint64

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, n := range d {
		n = unicode.ToUpper(n)

		if _, ok := h[n]; ok {
			h[n]++
		} else {
			return nil, errors.New("invalid DNS strand")
		}
	}

	return h, nil
}

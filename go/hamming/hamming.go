package hamming

import (
	"errors"
)

// Distance Calculate the Hamming Distance between two DNA strands.
func Distance(a, b string) (int, error) {
	aR := []rune(a)
	bR := []rune(b)

	if len(aR) != len(bR) {
		return 0, errors.New("lengths are not equal")
	}

	d := 0

	for idx := range aR {
		if aR[idx] != bR[idx] {
			d++
		}
	}

	return d, nil
}

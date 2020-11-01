package darts

import "math"

const (
	// INNER inner circle ridus
	INNER  = 1.0

	// MIDDLE middle circle ridus
	MIDDLE = 5.0

	// OUTER outer circle ridus
	OUTER  = 10.0
)

// Score get dart score
func Score(x, y float64) int {
	r := math.Hypot(x, y)

	switch {
	case r <= INNER:
		return 10
	case r <= MIDDLE:
		return 5
	case r <= OUTER:
		return 1
	default:
		return 0
	}
}

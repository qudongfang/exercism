// Package triangle Determine if a triangle is equilateral, isosceles, or scalene.
package triangle

import "math"

// Kind type of a triangle
type Kind int

const (
	// NaT not a triangle
	NaT Kind = iota

	// Equ equilateral
	Equ

	// Iso isosceles
	Iso

	// Sca scalene
	Sca
)

// KindFromSides determine if a triangle is equilateral, isosceles, or scalene.
func KindFromSides(a, b, c float64) Kind {
	if !valid(a, b, c) {
		return NaT
	}

	m := map[float64]bool{
		a: true,
		b: true,
		c: true,
	}

	switch len(m) {
	case 2:
		return Iso
	case 1:
		return Equ
	default:
		return Sca
	}
}

func valid(a, b, c float64) bool {
	if math.IsNaN(a+b+c) ||
		math.IsInf(a+b+c, 0) {
		return false
	}

	if (a+b) < c ||
		(b+c) < a ||
		(a+c) < b {
		return false
	}

	return a > 0 && b > 0 && c > 0
}

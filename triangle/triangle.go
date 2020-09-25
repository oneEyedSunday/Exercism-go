// Package triangle should have a package comment that summarizes what it's about.
package triangle

import "math"

// Kind represents the kind of triangle
type Kind int

// The Kind of triangles
const (
	// Pick values for the following identifiers used by the test program.
	NaT = 1 // not a triangle
	Equ = 2 // equilateral
	Iso = 3 // isosceles
	Sca = 4 // scalene
)

func anyInvalidItems(arr []float64) bool {
	for _, v := range arr {
		if v <= 0 || math.IsNaN(v) || math.IsInf(v, 1) || math.IsInf(v, -1) {
			return true
		}
	}
	return false
}

// KindFromSides determines the kind of triangle from its sides.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case anyInvalidItems([]float64{a, b, c}), (a+b < c) || (a+c < b) || (b+c < a):
		return NaT
	case a == b && b == c:
		return Equ
	case a != b && b != c && a != c:
		return Sca
	case (a == b) || (b == c) || (a == c):
		return Iso
	default:
		return NaT
	}
}

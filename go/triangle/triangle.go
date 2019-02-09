package triangle

import "math"

// Kind of triangle - not, equilateral, isosceles or scalene
type Kind int

// Types of Triangles
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides returns Kind of triangle based on line lengths
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if !ValidSides(a, b, c) {
		k = NaT
	} else if a == b && a == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}

	return k
}

// ValidSides For a shape to be a triangle at all, all sides have
// to be of length > 0, and the sum of the lengths of any two sides
// must be greater than or equal to the length of the third side.
func ValidSides(a, b, c float64) bool {
	result := true

	for _, num := range []float64{a, b, c} {
		if num <= 0 || math.IsNaN(num) || math.IsInf(num, 0) {
			result = false
		}
	}

	if a+b < c || a+c < b || c+b < a {
		result = false
	}

	return result
}

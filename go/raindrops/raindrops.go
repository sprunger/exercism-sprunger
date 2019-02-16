package raindrops

import "fmt"

// Convert a number to a string, the contents of which depend on the number's factors.
//   - If the number has 3 as a factor, output 'Pling'.
//   - If the number has 5 as a factor, output 'Plang'.
//   - If the number has 7 as a factor, output 'Plong'.
//   - If the number does not have 3, 5, or 7 as a factor,
//     just pass the number's digits straight through.
func Convert(drops int) string {
	var speak = ""

	if drops < 1 {
		return speak
	}

	if drops%3 == 0 {
		speak += "Pling"
	}

	if drops%5 == 0 {
		speak += "Plang"
	}

	if drops%7 == 0 {
		speak += "Plong"
	}

	if speak == "" {
		speak += fmt.Sprint(drops)
	}

	return speak
}

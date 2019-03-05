package isogram

import "unicode"

// IsIsogram returns true of all letters of string are unique
func IsIsogram(input string) bool {
	seen := make(map[rune]bool, len(input))

	for _, char := range input {
		letter := unicode.ToLower(char)

		// For this exercise, space and hyphen are not letters
		if letter == ' ' || letter == '-' {
			continue
		}

		if seen[letter] {
			return false
		}
		seen[letter] = true
	}

	return true
}

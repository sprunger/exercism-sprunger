package pangram

import (
	"strings"
)

// IsPangram determines if a sentence is a panagram -
// using every letter of the alphabet at least once
func IsPangram(input string) bool {
	sentence := strings.ToLower(input)

	for _, c := range "abcdefghijklmnopqrstuvwxyz" {
		if !strings.ContainsRune(sentence, c) {
			return false
		}
	}

	return true
}

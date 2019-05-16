package wordcount

import (
	"regexp"
	"strings"
)

// Frequency of words in a string
type Frequency map[string]int

// WordCount returns the map of frequency of words based on the input phrase.
func WordCount(input string) Frequency {
	count := Frequency{}

	input = strings.ToLower(input)

	// Very specific to test cases provided
	// remove all punctuation except apostrophe
	re := regexp.MustCompile(`[^0-9A-Za-z_']`)
	words := re.ReplaceAllLiteralString(input, " ")

	for _, word := range strings.Fields(words) {
		word = strings.Trim(word, "'") // Just word, not quotes
		count[word]++
	}
	return count
}

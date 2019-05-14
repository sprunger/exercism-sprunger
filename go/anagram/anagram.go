package anagram

import (
	"sort"
	"strings"
)

// SortedString returns are string with all of the characters sorted
func SortedString(s string) string {
	s = strings.ToLower(s)
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// Detect possible anagrams given a word and a list of candidates
func Detect(subject string, candidates []string) (anagrams []string) {
	if len(subject) == 0 {
		return
	}

	for _, candidate := range candidates {
		if !strings.EqualFold(subject, candidate) {
			if SortedString(candidate) == SortedString(subject) {
				anagrams = append(anagrams, candidate)
			}
		}
	}
	return
}

package etl

import "strings"

// Transform a map from old system to new
func Transform(input map[int][]string) map[string]int {

	result := make(map[string]int)

	for letterValue,letters := range input {
		for _,letter := range letters {
			result[strings.ToLower(letter)] = letterValue
		}
	}
	
	return result
}

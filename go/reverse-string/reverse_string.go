package reverse

// String reverses the given string
// It's less efficient than the canonical below, but far
// more legible to me
func String(input string) (output string) {
	//var runes = make([]rune, len(input))
	for _, char := range input {
		output = string(char) + output
	}
	return
}

// Canonical String reverses given string
// func String(s string) string {
// 	runes := []rune(s)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }
package accumulate

// Accumulate will run the provided converter on the input
func Accumulate(input []string, converter func(string) string) []string {

	result := []string{}

	for _, i := range input {
		result = append(result, converter(i))
	}

	return result
}

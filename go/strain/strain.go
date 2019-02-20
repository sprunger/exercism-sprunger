package strain

// Ints represents a slice of integer values
type Ints []int

// Strings represents slice of string values
type Strings []string

// Lists represents a slice of integer slices
type Lists [][]int

// Keep Ints
func (values Ints) Keep(f func(int) bool) (result Ints) {
	for _, value := range values {
		if f(value) {
			result = append(result, value)
		}
	}
	return
}

// Discard Ints
func (values Ints) Discard(f func(int) bool) (result Ints) {
	for _, value := range values {
		if !f(value) {
			result = append(result, value)
		}
	}
	return
}

// Keep Strings
func (values Strings) Keep(f func(string) bool) (result Strings) {
	for _, value := range values {
		if f(value) {
			result = append(result, value)
		}
	}
	return
}

// Discard Strings
func (values Strings) Discard(f func(string) bool) (result Strings) {
	for _, value := range values {
		if !f(value) {
			result = append(result, value)
		}
	}
	return
}

// Keep Lists
func (values Lists) Keep(f func([]int) bool) (result Lists) {
	for _, value := range values {
		if f(value) {
			result = append(result, value)
		}
	}
	return
}

// Discard Lists
func (values Lists) Discard(f func([]int) bool) (result Lists) {
	for _, value := range values {
		if !f(value) {
			result = append(result, value)
		}
	}
	return
}

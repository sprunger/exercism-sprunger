package twofer

import "fmt"

// Share one of your two things with somebody
// If name is blank, "you" is the default
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
